package auxi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/pewpewnor/auxi/res"
)

type testPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var expectedTestPerson = testPerson{
	Name: "TestName",
	Age:  69,
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	var params testPerson

	params.Name = r.URL.Query().Get("name")
	if params.Name == "" {
		http.Error(w, "query string missing field 'name'", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		http.Error(w, "query string missing field 'age' or field 'age' is not an integer", http.StatusBadRequest)
		return
	}
	params.Age = age

	res.RespondWithJSON(w, params, http.StatusOK)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var person testPerson

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "body malformed", http.StatusBadRequest)
		return
	}

	res.RespondWithJSON(w, person, http.StatusOK)
}

func testGetHandler(server *httptest.Server, t *testing.T) {
	resp, err := http.Get(server.URL + "/test" + fmt.Sprintf("?name=%v&age=%v", expectedTestPerson.Name, expectedTestPerson.Age))
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, resp.StatusCode)
	}

	var person testPerson
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Response body malformed %v", err)
	}

	if person.Name != expectedTestPerson.Name || person.Age != expectedTestPerson.Age {
		t.Errorf("Expected response %v, but got %v", expectedTestPerson, person)
	}
}

func testPostHandler(server *httptest.Server, t *testing.T) {
	requestBody, err := json.Marshal(expectedTestPerson)
	if err != nil {
		t.Fatalf("Failed to encode person to JSON: %v", err)
	}

	resp, err := http.Post(server.URL+"/test", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, resp.StatusCode)
	}

	var person testPerson
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Response body malformed %v", err)
	}

	if person.Name != expectedTestPerson.Name || person.Age != expectedTestPerson.Age {
		t.Errorf("Expected response %v, but got %v", expectedTestPerson, person)
	}
}

func TestHandleMethods(t *testing.T) {
	mux := NewServeMux()

	mux.HandleMethods("/test", MethodHandlers{
		GET:  getHandler,
		POST: postHandler,
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	testGetHandler(server, t)
	testPostHandler(server, t)
}
