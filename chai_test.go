package chai

import . "github.com/haoxins/supertest"
import "net/http/httptest"
import "net/http"
import "testing"
import "fmt"

// test GET
func TestGet(t *testing.T) {
	app := New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	s := httptest.NewServer(app)

	Request(s.URL).
		Get("/").
		Expect(200).
		Expect("hello").
		End()
}

// test HEAD
func TestHead(t *testing.T) {
	app := New()

	app.Head("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	s := httptest.NewServer(app)

	Request(s.URL).
		Head("/").
		Expect(200).
		End()
}

// test HEAD for GET route
func TestHeadGet(t *testing.T) {
	app := New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	s := httptest.NewServer(app)

	Request(s.URL).
		Head("/").
		Expect(200).
		End()
}

// test route precedence
func TestPrecedence(t *testing.T) {
	app := New()

	app.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	app.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})

	s := httptest.NewServer(app)

	Request(s.URL).
		Get("/foo").
		Expect(200).
		Expect("hello").
		End()
}

// test many routes
func TestMany(t *testing.T) {
	app := New()

	app.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	app.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})

	s := httptest.NewServer(app)

	Request(s.URL).
		Get("/foo").
		Expect(200).
		Expect("hello").
		End()

	Request(s.URL).
		Get("/bar").
		Expect(200).
		Expect("world").
		End()
}

// test params
func TestParams(t *testing.T) {
	app := New()

	app.Get("/user/:name/pet/:pet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get(":name")
		pet := r.URL.Query().Get(":pet")
		fmt.Println(r.URL.Query(), "99")
		fmt.Fprint(w, "user %s's pet %s", name, pet)
	})

	s := httptest.NewServer(app)

	Request(s.URL).
		Get("/user/tobi/pet/loki").
		Expect(200).
		// Expect("user tobi's pet loki").
		End()
}
