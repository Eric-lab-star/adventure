# function type for option arguments

```go
var defaultTmpl *template.Template

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.tmpl = t
	}
}

func NewHandler(story Story, opts ...HandlerOption) http.Handler {
	h := handler{story, defaultTmpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	story Story
	tmpl  *template.Template
}

```

NewHandler에 opt을 넣어줄려고 하는데 옵션이라서 타입을 설정하기 어렵고 코드가 지저분해질 수 있다. 이럴 때, 함수를 타입으로 만들어주면 타입을 편하게 설정할 수 있다.

먼저, handler를 인수로 가지는 함수 타입 HandlerOption을 만들어준다. 만들지 않아도 되지만 의미론적으로 만들어준다.

다음, WithTemplate함수와 같이, 옵션을 처리하는 함수를 만들어준다.
이렇게 하면 모든 옵션함수는 같은 타입을 가지면서 NewHandler를 깔끔하게 만들 수 있다.
