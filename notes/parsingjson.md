# parsing json

json파일을 읽는 방법

1. Unmarshal을 이용하는 방법

```go
func main() {
    r, err := os.ReadFile("gopher.json")
    if err != nil {
        fmt.Print("failed to open file")
        os.Exit(1)
    }
    story := Story{}
    err = json.Unmarshal(r, &story)
    if err != nil {
    log.Fatal(err)
    }
}
```

json.Unmarshal()의 r은 읽을 데이터 story는 읽는 데이터를 저장할 장소이다.
읽을 데이터의 타입은 Story라고 했지만 타입을 모르는 경우 빈 인터페이스로 타입의 변수를 만들어 저장을 하는 것도 가능하다.

2. newDecoder

```go
func main() {
    r, err := os.Open("gopher.json")
    if err != nil {
        fmt.Print("failed to open file")
        os.Exit(1)
    }
    story := Story{}
    decoder := json.NewDecoder(r)
    err = decoder.Decode(&story)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(story)
}
```

json.NewDecoder()는 Reader를 인수로 받는다. Open은 r는 Reader interface를 만족함으로 인수로 넣을 수 있다.
