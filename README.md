# regexpgen

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Regexp generator

## Install

Go version 1.17+

```
go get github.com/cristaloleg/regexpgen
```

## Example

```go
s := `foo(-(bar|baz)){2,4}`
var buf bytes.Buffer
if err := regexpgen.GenerateString(s, &buf, nil); err != nil {
    t.Fatal(err)
}
println(buf.String())
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristaloleg/regexpgen/workflows/build/badge.svg
[build-url]: https://github.com/cristaloleg/regexpgen/actions
[pkg-img]: https://pkg.go.dev/badge/cristaloleg/regexpgen
[pkg-url]: https://pkg.go.dev/github.com/cristaloleg/regexpgen
[reportcard-img]: https://goreportcard.com/badge/cristaloleg/regexpgen
[reportcard-url]: https://goreportcard.com/report/cristaloleg/regexpgen
[coverage-img]: https://codecov.io/gh/cristaloleg/regexpgen/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristaloleg/regexpgen
[version-img]: https://img.shields.io/github/v/release/cristaloleg/regexpgen
[version-url]: https://github.com/cristaloleg/regexpgen/releases
