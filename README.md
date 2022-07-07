# error
Go error Package

## 示例

examples/main.go
```go
import "github.com/tinystack/errors"

err := errors.New("error test")
// fmt: main.exampleError(main.go:28): example err

err := errors.Newf("example err: %s", "error message")
// fmt: main.exampleErrorf(main.go:32): example err: error message

err := errors.New("simple err")
wrapErr :=  errors.Wrap(err, "wrap error message")
// fmt: main.exampleWrap(main.go:37): wrap error message | Caused: main.exampleWrap(main.go:36): simple err

err := errors.New("simple err")
return errors.Wrapf(err, "wrap: %s", "error message")
// fmt: main.exampleWrapf(main.go:42): wrap: error message | Caused: main.exampleWrapf(main.go:41): simple err
```

## API


```go
// 简单错误返回
errors.New(text string)

// format方式错误返回
errors.errors.Newf(format string, args ...interface{})

// 嵌套方式错误返回
errors.Wrap(err error, text string) *Error

// 嵌套+format方式错误返回
errors. Wrapf(err error, format string, args ...interface{}) *Error
```