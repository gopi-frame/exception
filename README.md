# Exception
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/exception.svg)](https://pkg.go.dev/github.com/gopi-frame/exception)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/exception)](https://goreportcard.com/report/github.com/gopi-frame/exception)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package exception defines some common exceptions.

## Installation

```shell
go get -u -v github.com/gopi-frame/exception
```

## Import

```go
import "github.com/gopi-frame/exception"
```

## Exceptions

### NoSuchMethodException

`NoSuchMethodException` is used to indicate no such method.

Example:
```go
exception.NewNoSuchMethodException(new(Object), "method")
exception.NewUnexportedMethodException("method")
```

### ArgumentException

`ArgumentException` is used to indicate invalid arguments.

Example:

```go
exception.NewArgumentException("number", "-1", "number should be positive")
exception.NewEmptyArgumentException("number")
```

### RangeException

`RangeException` is used to indicate invalid range.

Example:
```go
exception.NewRangeException(1, 10)
```

### TimeException

`TimeException` is used to indicate timeout.

Example:
```go
exception.NewTimeException()
```

### TypeException

`TypeException` is used to indicate invalid type.

Example:
```go
exception.NewTypeException("Invalid type for parameter `number`")
```

### UndefinedIndexException 

`UndefinedIndexException` is used to indicate undefined index.

Example:
```go
exception.NewUndefinedIndexException(10)
```

### UnsupportedException

`UnsupportedException` is used to indicate unsupported operation.

Example:
```go
exception.NewUnsupportedException("Unsupported operation")
```

### ValueException

`ValueException` is mainly used to convert a non `error` type value to an Exception.
Typically used in `recover` when a function panics a value instead of error.

Example:
```go
exception.NewValueException("value")
```