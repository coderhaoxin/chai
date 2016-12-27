package chai

import "github.com/justinas/alice"
import "github.com/go-zoo/bone"
import "net/http"

const Version = "1.0.0"

type Chai struct {
	mux   *bone.Mux
	chain alice.Chain
}

func New() *Chai {
	return &Chai{
		mux:   bone.New(),
		chain: alice.New(),
	}
}

func (a *Chai) Use(mw ...alice.Constructor) {
	a.chain = a.chain.Append(mw...)
}

func (a *Chai) Head(path string, h interface{}) {
	a.mux.Head(path, handler(h))
}

func (a *Chai) Get(path string, h interface{}) {
	a.mux.Get(path, handler(h))
}

func (a *Chai) Post(path string, h interface{}) {
	a.mux.Post(path, handler(h))
}

func (a *Chai) Put(path string, h interface{}) {
	a.mux.Put(path, handler(h))
}

func (a *Chai) Delete(path string, h interface{}) {
	a.mux.Delete(path, handler(h))
}

func (a *Chai) Options(path string, h interface{}) {
	a.mux.Options(path, handler(h))
}

func (a *Chai) Listen(addr string) error {
	handler := a.chain.Then(a.mux)
	http.Handle("/", handler)
	return http.ListenAndServe(addr, nil)
}

func (a *Chai) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := a.chain.Then(a.mux)
	handler.ServeHTTP(w, r)
}

func handler(h interface{}) http.Handler {
	switch h.(type) {
	case func(w http.ResponseWriter, r *http.Request):
		return http.HandlerFunc(h.(func(w http.ResponseWriter, r *http.Request)))
	case http.Handler:
		return h.(http.Handler)
	default:
		panic("invalid handler")
	}
}
