package auxi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

var testServer *httptest.Server

var expectedPerson Person

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMain(m *testing.M) {
	mux := NewServeMux()

	mux.HandleMethods("/test", MethodHandlers{
		GET: func(w http.ResponseWriter, r *http.Request) {
			var params Person

			params.Name = r.URL.Query().Get("name")
			if params.Name == "" {
				http.Error(w, "query string missing field 'name'",
					http.StatusBadRequest)
				return
			}

			age, err := strconv.Atoi(r.URL.Query().Get("age"))
			if err != nil {
				http.Error(w,
					"query string missing field 'age' or field 'age' is not an integer",
					http.StatusBadRequest)
				return
			}
			params.Age = age

			Respond.JSON(w, params, http.StatusOK)
		},
		POST: func(w http.ResponseWriter, r *http.Request) {
			var person Person

			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				http.Error(w, "body malformed", http.StatusBadRequest)
				return
			}

			Respond.JSON(w, person, http.StatusOK)
		},
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	m.Run()
}

func TestGetHandler(t *testing.T) {
	params := url.Values{}
	params.Add("name", expectedPerson.Name)
	params.Add("age", strconv.Itoa(expectedPerson.Age))

	resp, err := http.Get(fmt.Sprintf(
		"%v/test?%v", testServer.URL, params.Encode()))
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode,
			resp.StatusCode)
	}

	var person Person
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Response body malformed %v", err)
	}

	if person.Name != expectedPerson.Name || person.Age != expectedPerson.Age {
		t.Errorf("Expected response %v, but got %v", expectedPerson, person)
	}
}

func TestPostHandler(t *testing.T) {
	requestBody, err := json.Marshal(expectedPerson)
	if err != nil {
		t.Fatalf("Failed to encode person to JSON: %v", err)
	}

	resp, err := http.Post(
		testServer.URL+"/test", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode,
			resp.StatusCode)
	}

	var person Person
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Response body malformed %v", err)
	}

	if person.Name != expectedPerson.Name || person.Age != expectedPerson.Age {
		t.Errorf("Expected response %v, but got %v", expectedPerson, person)
	}
}
