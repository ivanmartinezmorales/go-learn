package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepe": 20,
			"Phil": 10,
			"Ivan": 1,
			"Zack": -1,
		},
	}
	server := &PlayerServer{&store}

	t.Run("Returns Zack's score", func(t *testing.T) {
		request := newGetScoreRequest("Zack")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "-1")
	})

	t.Run("Returns Ivan's score", func(t *testing.T) {
		request := newGetScoreRequest("Ivan")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1")
	})

	t.Run("Returns Pepe's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepe")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Returns Phil's score", func(t *testing.T) {
		request := newGetScoreRequest("Phil")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("Returns 404 for missing players.", func(t *testing.T) {
		request := newGetScoreRequest("PatientZero")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

// newGetScoreRequest - Helper function that returns the GET request
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Incorrect response body, got: %s, and want %s", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("The response body is wrong, got %d, want %d", got, want)
	}
}
