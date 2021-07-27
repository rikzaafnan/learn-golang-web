package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {

	// create cookie
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "Success create cookie")

}

func GetCookie(writer http.ResponseWriter, request *http.Request) {

	// get cookie
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {

		fmt.Fprintf(writer, "No Cookie")

	} else {

		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestCookie(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
func TestSetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=hihi", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {

		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)

	}

	// body, _ := io.ReadAll(response.Body)

	// fmt.Println(string(body))
}

func TestGetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "rick"
	cookie.Path = "/"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
