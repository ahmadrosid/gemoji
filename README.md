![](/gemoji.png)

<p align="center">
  <a href="https://github.com/ahmadrosid/gemoji/actions"
    ><img
      src="https://github.com/ahmadrosid/gemoji/workflows/build/badge.svg"
      alt="GitHub Actions workflow status"
  /></a>
  <a href="https://pkg.go.dev/badge/github.com/ahmadrosid/gemoji"
    ><img
      src="https://pkg.go.dev/github.com/ahmadrosid/gemoji"
      alt="PkgGoDev"
  /></a>
</p>

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

### Get emoji by the name
```go
println(gemoji.Get(":rocket:")) // 🚀
println(gemoji.Get("heartbeat")) // 💓
```

## LICENSE
The MIT License (MIT) - see [`LICENSE.md`](https://github.com/ahmadrosid/gemoji/blob/master/LICENSE.md) for more details
