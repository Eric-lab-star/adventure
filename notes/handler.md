# template을 만드는 방법

```go
tmpl := template.Must(template.ParseFiles("../static/template.gohtml"))
    err := tmpl.Execute(w, h.story["intro"])
    if err != nil {
        fmt.Println("exetue err")
        os.Exit(1)
    }
```

template.ParseFiles를 이용해서 template 파일을 불러와서 읽을 수 있다. 이후에 Execute를 해서 템플릿에 데이터와 출력할 곳을 지정할 수 있다.

# ListenAndServe를 만족시키는 방법

1. http.Handler 인터페이스를 만족시키는 스트럭트를 만들어서 ListenAndServe의 인수로 넣어주기

```go
func NewHandler(story Story) http.Handler {
    return handler{story}
}

type handler struct {
    story Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("../static/template.gohtml"))
    err := tmpl.Execute(w, h.story["intro"])
    if err != nil {
        fmt.Println("exetue err")
        os.Exit(1)
    }
}
```

2. http.Handlefunc을 만들고 ListenAndServe에는 nil을 넣어준다.

```go
func chapter(data story.Chapter) http.HandlerFunc {
    tmpl := template.Must(template.ParseFiles("../static/template.gohtml"))
    return func(w http.ResponseWriter, r *http.Request) {
        err := tmpl.Execute(w, data)
        if err != nil {
            fmt.Println("err: tmpl execute")
            os.Exit(1)
        }

    }
}
```
