```go
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, " ")
	if path == "/" || path == "" {
		path = "/intro"
	}

	if chapter, ok := h.story[path[1:]]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			fmt.Printf("tmpl.execute error:\n%v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Page not found", http.StatusNotFound)

}
```

strings.Trim의 첫번째 인수는 자르는 대상의 string, 두번째 인수는 string에서 자르고 싶은 문자.

r.URL.Path에의 / 를 제외한 문자를 h.story에 입력해주면 chapter를 얻을 수 있음. 이 챕터를 tmpl의 데이터로 넣어주면 클릭할 때마다 새로운 데이터의 페이지가 로드된다.
