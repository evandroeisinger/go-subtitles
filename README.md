# subtitles ![Go](https://github.com/evandroeisinger/subtitles/workflows/Go/badge.svg)
Golang package for Subtitles manipulation. 

```shell
go get -u github.com/evandroeisinger/subtitles
```

### Supported formats
- [x] SRT
- [] WebVTT

### Available operations:
- [x] Parsing 
- [] Writing
- [] Shifting
- [] Merging

### Usage
```golang
subtitle, err := subtitle.Load("example.srt")
if err != nil {
    fmt.PrintLn(err)
}
```