![](/gemoji.png)

<h1></h1>

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
