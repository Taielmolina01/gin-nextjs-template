package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"movie-reservation-system/initializers"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/application"
	"net/http"
	"net/http/httptest"
	"bytes"
	"testing"
	"encoding/json"
	"movie-reservation-system/configuration"
)

func setUpRouterTest() (*gin.Engine, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connecting to mock database: %w", err)
	}

	tables := application.GetAllModels()
	for _, t := range tables {
		err = db.AutoMigrate(t)
		if err != nil {
			return nil, fmt.Errorf("Error creating tables in mock database: %w", err)
		}
	}

	config := application.LoadConfigTest("3000", "host=localhost user=postgres password=taiel0101 port=5432 sslmode=disable dbname=movie-system-db", "HS256", "ASDADASD")

	engine := application.CreateTestEngine(db, config)

	gin.SetMode(gin.ReleaseMode)

	return engine, nil
}

func UseRouter(t *testing.T) *gin.Engine {
	engine, err := setUpRouterTest()

	if err != nil {
		t.Fatalf("Error setting up test engine: %v", err)
	}
	return engine
}

func PerformRequest(t *testing.T, engine *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	var req *http.Request
	var err error

	if body == "" {
		req, err = http.NewRequest(method, path, nil)
	} else {
		req, err = http.NewRequest(method, path, bytes.NewBufferString(body))
	}

	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, req)

	return recorder
}

func PerformRequestWithRequest(t *testing.T, engine *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
    recorder := httptest.NewRecorder()
    engine.ServeHTTP(recorder, req)
    return recorder
}


func GetAccessToken(recorder *httptest.ResponseRecorder) (string, error) {
	var responseBody models.TokenResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.AccessToken.AccessToken, nil
}

type UserLoginData struct {
	FirstName	string
	LastName	string
	Email	string
	Password	string
	Role	string
}

func CreateUserAndLogin(userData UserLoginData, t *testing.T, engine *gin.Engine) string {
	jsonBody := fmt.Sprintf(`{"firstname": "%s", "lastname": "%s, "email": "%s", "password": "%s", "role": "%s"}`, userData.FirstName, userData.LastName, userData.Email, userData.Password, userData.Role)

	recorder := PerformRequest(t, engine, "POST", "/users", jsonBody)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	secondJsonBody := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, userData.Email, userData.Password)

	secondRecorder := PerformRequest(t, engine, "POST", "/login", secondJsonBody)

	if secondRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, secondRecorder.Code)
	}

	accessToken, err := GetAccessToken(secondRecorder)

	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	return accessToken
}