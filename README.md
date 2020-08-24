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
- [x] Shifting
- [ ] Merging

### Usage
```golang
// Loads subtitle from file 
sub, err := subtitles.Load("example.srt")
if err != nil {
    fmt.Println(err)
}

// Shifts up all subtitle blocks timestamp
sub.Shift(sub, time.ParseDuration("1m30s"))

// Shifts down all subtitle blocks timestamp
sub.Shift(sub, time.ParseDuration("-1m30s"))

// Writes subtitle according to extension format
content, err := subtitles.Write(sub, "example.srt")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(content)
}
```