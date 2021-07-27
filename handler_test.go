package learngolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {

		// logic web
		fmt.Fprint(writer, "Hello Wolrd")

	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writter, "Hello Wolrd")

	})

	mux.HandleFunc("/h1", func(writter http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writter, "Hi")

	})

	mux.HandleFunc("/images/", func(writter http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writter, "images")

	})

	mux.HandleFunc("/images/thumbnails/", func(writter http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writter, "thumbnails")

	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {

		// logic web
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)

	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
