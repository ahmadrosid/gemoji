![](/gemoji.png)

<div align="center">

  [![GitHub Actions workflow status](https://github.com/ahmadrosid/gemoji/workflows/build/badge.svg)](https://github.com/ahmadrosid/gemoji/actions) [![Go Report Card](https://pkg.go.dev/badge/github.com/ahmadrosid/gemoji)](https://pkg.go.dev/badge/github.com/ahmadrosid/gemoji) [![Go Report Card](https://goreportcard.com/badge/mattn/go-colorable)](https://goreportcard.com/report/mattn/go-colorable)

</div>

## Install
```bash
go get github.com/ahmadrosid/gemoji
```

## Usage

### Format
```go
println(gemoji.Format("This is :rocket: and :smile:"))
// This is 🚀 and 😄

println(gemoji.Format("This is :rocket: and :smile: :invalid:"))
// This is 🚀 and 😄 :invalid:
```

### Get Emoji
```go
println(gemoji.Get(":rocket:")) // 🚀
println(gemoji.Get("heartbeat")) // 💓
```

## Credit
[emojicpp](/99x/emojicpp)

## LICENSE
The MIT License (MIT) - see [`LICENSE.md`](https://github.com/ahmadrosid/gemoji/blob/master/LICENSE.md) for more details
