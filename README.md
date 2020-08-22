# subtitles ![Go](https://github.com/evandroeisinger/subtitles/workflows/Go/badge.svg)
Golang package for Subtitles manipulation. 

```shell
go get -u github.com/evandroeisinger/go-subtitles
```

### Supported formats
- [x] SRT
- [ ] WebVTT

### Available operations:
- [x] Parsing 
- [x] Writing
- [ ] Shifting
- [ ] Merging

### Usage
```golang
// Load subtitle
sub, err := subtitle.Load("example.srt")
if err != nil {
    fmt.Println(err)
}

// Write subtitle
content, err := subtitle.Write(sub, "example.srt")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(content)
}
```