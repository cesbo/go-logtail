# go-logtail

[![Go Reference](https://pkg.go.dev/badge/github.com/cesbo/go-logtail.svg)](https://pkg.go.dev/github.com/cesbo/go-logtail)

[Zerolog](https://github.com/rs/zerolog) writer for [Logtail](https://betterstack.com/logtail)

## Installation

To install the library use the following command in the project directory:

```
go get github.com/cesbo/go-logtail
```

## Quick Start

```go
token := "LOGTAIL_TOKEN"
log.Logger = logtail.NewLogtail(token).NewLogger()
```
