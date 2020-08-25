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
- [x] Concatenating
- [x] Merging
- [x] Shifting
- [x] Cutting 
- [ ] Slicing

#### Usage
```golang
// Loads subtitle from file 
sub, err := subtitles.Load("example.srt")
if err != nil {
    fmt.Println(err)
}

// Writes subtitle to VTT format (writes according to extension format)
content, _ := subtitles.Write(sub, "example.vtt")

...
```

#### Concating
```golang
// Loads subtitle from file 
sub_a, _ := subtitles.Load("sub_a.srt")
sub_b, _ := subtitles.Load("sub_a.srt")

// Concats subtitles fixing blocks timestamps
sub_ab := subtitles.Merge(sub_a, sub_b)

...
```

#### Merging
```golang
// Loads subtitle from file 
sub_a, _ := subtitles.Load("sub_a.srt")
sub_b, _ := subtitles.Load("sub_a.srt")

// Merges subtitles preserving timestamps
sub_ab := subtitles.Merge(sub_a, sub_b)

...
```

#### Shifting
```golang
// Loads subtitle from file 
sub, _ := subtitles.Load("example.srt")

// Shifts up all subtitle blocks timestamp
sub.Shift(time.ParseDuration("1m30s"))

// Shifts down all subtitle blocks timestamp
sub.Shift(time.ParseDuration("-1m30s"))

...
```

#### Cutting
```golang
// Loads subtitle from file 
sub, _ := subtitles.Load("example.srt")

startAt := time.ParseDuration("30s")
finishAt := time.ParseDuration("2m")

// Cuts subtitle preserving blocks
sub.Cut(startAt, finishAt)

...
```
