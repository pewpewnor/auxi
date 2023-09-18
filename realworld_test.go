package auxi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pewpewnor/auxi/logmsg"
	"github.com/pewpewnor/auxi/respond"
	"github.com/pewpewnor/auxi/utils"
)

var testServer *httptest.Server

var expectedPerson = Person{
	Name: "Test Name",
	Age:  "69",
}

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func TestMain(m *testing.M) {
	log.SetPrefix("[TEST] ")
	log.SetFlags(log.Llongfile)

	mux := NewServeMux()

	mux.HandleMethods("/test", MethodHandlers{
		GET: func(w http.ResponseWriter, r *http.Request) {
			var params Person

			params.Name = r.URL.Query().Get("name")
			if params.Name == "" {
				log.Println(logmsg.Err("query string missing field 'name'"))
				respond.JSON(w,
					respond.SError("query string missing field 'name'"),
					http.StatusBadRequest)
				return
			}

			params.Age = r.URL.Query().Get("age")
			if params.Age == "" {
				log.Println(logmsg.Err("query string missing field 'age'"))
				respond.JSON(w,
					respond.SError("query string missing field 'age'"),
					http.StatusBadRequest)
				return
			}

			respond.JSON(w, params, http.StatusOK)
		},
		POST: func(w http.ResponseWriter, r *http.Request) {
			var person Person

			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				log.Println(logmsg.Err(err))
				respond.JSON(w, respond.SErrorFromErr("body malformed", err),
					http.StatusBadRequest)
				return
			}

			respond.JSON(w, person, http.StatusOK)
		},
		PUT: func(w http.ResponseWriter, r *http.Request) {
			apiKey, err := utils.NewAuthorizationHeader(r.Header).GetBearerToken()
			if err != nil {
				log.Println(logmsg.Err(err))
				respond.JSON(w,
					respond.SErrorFromErr("Auth header malformed", err),
					http.StatusBadRequest)
				return
			}

			respond.JSON(w, apiKey, http.StatusOK)
		},
		DELETE: func(w http.ResponseWriter, r *http.Request) {
			var person Person

			if err := utils.BindQueryString(r, &person); err != nil {
				log.Println(logmsg.Err(err))
				respond.JSON(w,
					respond.SErrorFromErr("Query string malformed", err),
					http.StatusBadRequest)
				return
			}

			respond.JSON(w, person, http.StatusOK)
		},
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	m.Run()
}

func TestGetHandlerAndQueryString(t *testing.T) {
	params := url.Values{}
	params.Add("name", expectedPerson.Name)
	params.Add("age", expectedPerson.Age)

	resp, err := http.Get(fmt.Sprintf(
		"%v/test?%v", testServer.URL, params.Encode()))
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but instead got %v",
			expectedStatusCode, resp.StatusCode)
	}

	var person Person
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Respond body malformed: %v", err)
	}

	if person.Name != expectedPerson.Name || person.Age != expectedPerson.Age {
		t.Errorf("Expected response %v, but instead got %v",
			expectedPerson, person)
	}
}

func TestPostHandler(t *testing.T) {
	requestBody, err := json.Marshal(expectedPerson)
	if err != nil {
		t.Fatalf("Failed to encode request to JSON: %v", err)
	}

	resp, err := http.Post(
		testServer.URL+"/test", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but instead got %v",
			expectedStatusCode, resp.StatusCode)
	}

	var person Person
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Response body malformed: %v", err)
	}

	if person.Name != expectedPerson.Name || person.Age != expectedPerson.Age {
		t.Errorf("Expected response %v, but instead got %v", expectedPerson,
			person)
	}
}

func TestPutHandlerAndAuthHeaderGrabber(t *testing.T) {
	expectedApiKey := "secretapitoken"

	req, err := http.NewRequest(http.MethodPut, testServer.URL+"/test", nil)
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+expectedApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send PUT request: %v", err)
	}
	defer resp.Body.Close()

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but instead got %v",
			expectedStatusCode, resp.StatusCode)
	}

	var apiKey string
	err = json.NewDecoder(resp.Body).Decode(&apiKey)
	if err != nil {
		t.Errorf("Response body malformed: %v", err)
	}

	if apiKey != expectedApiKey {
		t.Errorf("Expected response %v, but instead got %v", expectedApiKey,
			apiKey)
	}
}

func TestDeleteHandlerAndBindQueryString(t *testing.T) {
	params := url.Values{}
	params.Add("name", expectedPerson.Name)
	params.Add("age", expectedPerson.Age)

	req, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("%v/test?%v", testServer.URL, params.Encode()), nil)
	if err != nil {
		t.Fatalf("Failed to create DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send DELETE request: %v", err)
	}
	defer resp.Body.Close()

	if expectedStatusCode := http.StatusOK; resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but instead got %v",
			expectedStatusCode, resp.StatusCode)
	}

	var person Person
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		t.Errorf("Response body malformed: %v", err)
	}

	if person.Name != expectedPerson.Name || person.Age != expectedPerson.Age {
		t.Errorf("Expected response %v, but instead got %v", expectedPerson,
			person)
	}
}
