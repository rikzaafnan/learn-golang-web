package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Sayhello(writer http.ResponseWriter, request *http.Request) {

	name := request.URL.Query().Get("name")

	if name == "" {

		fmt.Fprint(writer, "Hello")

	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestQuery(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello?name=eko", nil)
	recorder := httptest.NewRecorder()

	Sayhello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {

	firstName := request.URL.Query().Get("first_name")
	lastNaame := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastNaame)

}
func TestMultipleQueryParameter(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello?first_name=eko&last_name=kennedy", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprintf(writer, strings.Join(names, " "))
}
func TestMultipleParameterValues(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello?name=eko&name=kennedy", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
