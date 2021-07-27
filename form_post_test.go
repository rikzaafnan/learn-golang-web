package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// cara cepat ambil data dari form tanpa parsing dahulu
	// request.PostFormValue("first_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {

	requestBody := strings.NewReader("first_name=Eko&last_name=kennedy")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "Application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
