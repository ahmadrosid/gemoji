![](/gemoji.png)

<p align="center">
  <a href="https://github.com/ahmadrosid/gemoji/actions"
    ><img
      src="https://github.com/ahmadrosid/gemoji/workflows/build/badge.svg"
      alt="GitHub Actions workflow status"
  /></a>
  <a href="https://pkg.go.dev/badge/github.com/ahmadrosid/gemoji"
    ><img
      src="https://camo.githubusercontent.com/a640060291033ca59a91868a40c3b60c534b7ca7f2d1d9d56ee1d2f641357090/68747470733a2f2f706b672e676f2e6465762f62616467652f6769746875622e636f6d2f66617469682f636f6c6f72"
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
