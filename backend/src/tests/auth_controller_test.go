package tests

import (
	"net/http"
	"testing"
)

func TestLoginNonExistentUser(t *testing.T) {
	t.Log("Try to login with an user that is not registered")
	router := UseRouter(t)

	jsonBody := `{"Email": "johndoe@gmail.com", "Password": "myPassword"}`

	recorder := PerformRequest(t, router, "POST", "/login", jsonBody)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestLoginExistentUserWithWrongPassword(t *testing.T) {
	t.Log("Try to login with an user with wrong password")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"Email": "johndoe@gmail.com", "Password": "anotherPassword"}`

	secondRecorder := PerformRequest(t, router, "POST", "/login", secondJsonBody)

	if secondRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, secondRecorder.Code)
	}

	expected := `{"error":"The entered password is not the user's password"}`
	if secondRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
	}
}

func TestLoginWithCorrectFields(t *testing.T) {
	t.Log("Try to login with a registered user with correct password")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"email": "johndoe@gmail.com", "password": "myPassword"}`

	secondRecorder := PerformRequest(t, router, "POST", "/login", secondJsonBody)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}

	/*
		expected := `{"message":"The entered password is not the user's password"}`
		if secondRecorder.Body.String() != expected {
			t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
		}
	*/

	// Im not so sure about how i should test the response token, so i only check the status for now
}

func TestLogoutNotExistentUser(t *testing.T) {
	t.Log("Try to logout with an non existent user")

	router := UseRouter(t)

	recorder := PerformRequest(t, router, "POST", "/logout/johndoe@gmail.com", "jsonBody")

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, recorder.Code)
	}

	expected := `{"error":"User with email johndoe@gmail.com is not registered"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, recorder.Body.String())
	}
}

func TestLogoutWithANotLoggedUser(t *testing.T) {
	t.Log("Try to logout with a not logged user")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondRecorder := PerformRequest(t, router, "POST", "/logout/johndoe@gmail.com", "")

	if secondRecorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, secondRecorder.Code)
	}

	expected := `{"error":"Missing authentication token"}` 
	if secondRecorder.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, secondRecorder.Body.String())
	}
}

func TestLogoutWithALoggedUser(t *testing.T) {
	t.Log("Try to logout with a logged user")

	router := UseRouter(t)

	jsonBody := `{"name": "John Doe", "email": "johndoe@gmail.com", "password": "myPassword"}`

	recorder := PerformRequest(t, router, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	secondJsonBody := `{"email": "johndoe@gmail.com", "password": "myPassword"}`

	secondRecorder := PerformRequest(t, router, "POST", "/login", secondJsonBody)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}

	accessToken, err := GetAccessToken(secondRecorder)

	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	req, _ := http.NewRequest("POST", "/logout/johndoe@gmail.com", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	
	thirdRecorder := PerformRequestWithRequest(t, router, req)

	if thirdRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, thirdRecorder.Code)
	}
}

