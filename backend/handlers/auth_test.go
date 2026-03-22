package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

<<<<<<< HEAD
func setupAuthRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}
=======
func TestRegisterCreatesUser(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
>>>>>>> f85f4d5 (Add Sprint 2 backend tests and documentation)

	router := gin.New()
	router.POST("/register", handler.Register)

	rec := performRequest(t, router, http.MethodPost, "/register", RegisterRequest{
		Email:     "student@campus.edu",
		Password:  "secret123",
		FirstName: "Sam",
		LastName:  "Student",
		Phone:     "555-1000",
		Role:      "student",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	var user models.User
	if err := db.Where("email = ?", "student@campus.edu").First(&user).Error; err != nil {
		t.Fatalf("expected user to be persisted: %v", err)
	}
}

func TestRegisterRejectsDuplicateEmail(t *testing.T) {
	db := setupTestDB(t)
	createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
	})

	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/register", handler.Register)

<<<<<<< HEAD
func TestLoginSuccess(t *testing.T) {
	r := setupAuthRouter(t)

	registerBody := testutil.MustMarshal(map[string]any{
		"email":     "student3@uf.edu",
		"password":  "password123",
		"firstName": "A",
		"lastName":  "B",
		"phone":     "+1234567890",
		"role":      "student",
	})
	registerResp := testutil.PerformRequest(r, http.MethodPost, "/api/auth/register", registerBody)
	if registerResp.Code != http.StatusCreated {
		t.Fatalf("register failed: status=%d body=%s", registerResp.Code, registerResp.Body.String())
	}

	loginBody := testutil.MustMarshal(map[string]any{
		"email":    "student3@uf.edu",
		"password": "password123",
	})
	loginResp := testutil.PerformRequest(r, http.MethodPost, "/api/auth/login", loginBody)
	if loginResp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusOK, loginResp.Code, loginResp.Body.String())
=======
	rec := performRequest(t, router, http.MethodPost, "/register", RegisterRequest{
		Email:     "student@campus.edu",
		Password:  "secret123",
		FirstName: "Sam",
		LastName:  "Student",
		Phone:     "555-1000",
		Role:      "student",
	})

	if rec.Code != http.StatusConflict {
		t.Fatalf("expected status 409, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLoginAuthenticatesUser(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/login", handler.Login)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "secret123",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[AuthResponse](t, rec)
	if resp.ID != user.ID {
		t.Fatalf("expected response ID %s, got %s", user.ID, resp.ID)
	}
}

func TestLoginRejectsInvalidPassword(t *testing.T) {
	db := setupTestDB(t)
	createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/login", handler.Login)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "wrongpass",
	})

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLogoutReturnsSuccess(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)

	router := gin.New()
	router.POST("/logout", handler.Logout)

	rec := performRequest(t, router, http.MethodPost, "/logout", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestRefreshTokenReturnsPlaceholderToken(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)

	router := gin.New()
	router.POST("/refresh", handler.RefreshToken)

	rec := performRequest(t, router, http.MethodPost, "/refresh", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
>>>>>>> f85f4d5 (Add Sprint 2 backend tests and documentation)
	}
}
