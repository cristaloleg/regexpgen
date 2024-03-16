package regexpgen

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"regexp/syntax"
	"time"
)

var defaultMax = 32

func String(re *syntax.Regexp) (string, error) {
	buf := &bytes.Buffer{}
	err := Generate(re, buf, rand.New(rand.NewSource(time.Now().Unix())))
	return buf.String(), err
}

func GenerateString(s string, w *bytes.Buffer, rnd *rand.Rand) error {
	re, err := syntax.Parse(s, syntax.Perl)
	if err != nil {
		return err
	}
	return Generate(re, w, rnd)
}

func Generate(re *syntax.Regexp, w *bytes.Buffer, rnd *rand.Rand) error {
	if rnd == nil {
		rnd = rand.New(rand.NewSource(time.Now().Unix()))
	}
	return gen(re, w, rnd)
}

func gen(re *syntax.Regexp, w *bytes.Buffer, rnd *rand.Rand) error {
	switch re.Op {
	case syntax.OpNoMatch, syntax.OpEmptyMatch:
		return nil

	case syntax.OpLiteral:
		w.WriteString(string(re.Rune))
		return nil

	case syntax.OpCharClass:
		var sum int64
		for i := 0; i < len(re.Rune); i += 2 {
			sum += 1 + int64(re.Rune[i+1]-re.Rune[i])
		}

		nth := rune(randint(rnd, sum))
		for i := 0; i < len(re.Rune); i += 2 {
			min, max := re.Rune[i], re.Rune[i+1]
			delta := max - min
			if nth <= delta {
				w.WriteRune(min + nth)
				return nil
			}
			nth -= 1 + delta
		}
		panic("unreachable")

	case syntax.OpAnyCharNotNL:
		w.WriteRune(rune(' ' + randint(rnd, 95)))
		return nil

	case syntax.OpAnyChar:
		i := int(randint(rnd, 96))
		ch := rune(' ' + i)
		if i == 95 {
			ch = '\n'
		}
		w.WriteRune(ch)
		return nil

	case syntax.OpBeginLine:
		if w.Len() != 0 {
			w.WriteByte('\n')
		}
		return nil

	case syntax.OpEndLine:
		if w.Len() == 0 {
			return io.EOF
		}
		w.WriteByte('\n')
		return nil

	case syntax.OpBeginText:
		return nil

	case syntax.OpEndText:
		return io.EOF

	case syntax.OpWordBoundary, syntax.OpNoWordBoundary:
		panic("regexpgen: word boundaries are not supported")

	case syntax.OpCapture, syntax.OpConcat:
		for _, re := range re.Sub {
			if err := gen(re, w, rnd); err != nil {
				return err
			}
		}
		return nil

	case syntax.OpStar, syntax.OpPlus:
		var min int64
		if re.Op == syntax.OpPlus {
			min = 1
		}
		max := min + int64(defaultMax)

		for sz := min + randint(rnd, max-min+1); sz > 0; sz-- {
			for _, re := range re.Sub {
				gen(re, w, rnd)
			}
		}
		return nil

	case syntax.OpQuest:
		if randint(rnd, 1<<31-1) > 1<<31-1 {
			for _, re := range re.Sub {
				if err := gen(re, w, rnd); err != nil {
					return err
				}
			}
		}
		return nil

	case syntax.OpRepeat:
		min, max := int64(re.Min), int64(re.Max)
		if max == -1 {
			max = min + int64(defaultMax)
		}
		for sz := min + randint(rnd, max-min+1); sz > 0; sz-- {
			for _, re := range re.Sub {
				if err := gen(re, w, rnd); err != nil {
					return err
				}
			}
		}
		return nil

	case syntax.OpAlternate:
		nth := randint(rnd, int64(len(re.Sub)))
		return gen(re.Sub[nth], w, rnd)

	default:
		panic(fmt.Sprintf("unknown syntax.Op: %v", re.Op))
	}
}

func randint(r *rand.Rand, max int64) int64 {
	if max <= 1 {
		return 0
	}
	return r.Int63n(max)
}
