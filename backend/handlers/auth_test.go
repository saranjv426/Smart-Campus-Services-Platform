package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/middleware"
	"smart-campus-services/models"
	"smart-campus-services/testutil"
)

func setupAuthTestRouter(dbHandler *AuthHandler) *gin.Engine {
	router := gin.New()
	router.POST("/register", dbHandler.Register)
	router.POST("/login", dbHandler.Login)
	router.POST("/logout", dbHandler.Logout)
	router.POST("/refresh", dbHandler.RefreshToken)
	return router
}

func TestRegisterCreatesUser(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

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

func TestDuplicateSignup(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	body := testutil.MustMarshal(RegisterRequest{
		Email:     "student@campus.edu",
		Password:  "secret123",
		FirstName: "Sam",
		LastName:  "Student",
		Phone:     "555-1000",
		Role:      "student",
	})

	first := testutil.PerformRequest(router, http.MethodPost, "/register", body)
	if first.Code != http.StatusCreated {
		t.Fatalf("expected initial signup status 201, got %d with body %s", first.Code, first.Body.String())
	}

	duplicate := testutil.PerformRequest(router, http.MethodPost, "/register", body)
	if duplicate.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", duplicate.Code, duplicate.Body.String())
	}
}

func TestRegisterRejectsInvalidRole(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/register", map[string]any{
		"email":     "student@campus.edu",
		"password":  "secret123",
		"firstName": "Sam",
		"lastName":  "Student",
		"phone":     "555-1000",
		"role":      "visitor",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestRegisterInvalidEmail(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/register", map[string]any{
		"email":     "not-an-email",
		"password":  "secret123",
		"firstName": "Sam",
		"lastName":  "Student",
		"phone":     "555-1000",
		"role":      "student",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestRegisterMissingPhone(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/register", map[string]any{
		"email":     "student@campus.edu",
		"password":  "secret123",
		"firstName": "Sam",
		"lastName":  "Student",
		"role":      "student",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLoginAuthenticatesUser(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

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
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "wrongpass",
	})

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestProtectedMissingToken(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/logout", middleware.AuthRequired(), handler.Logout)

	rec := testutil.PerformRequest(router, http.MethodPost, "/logout", nil)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLogoutSuccess(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := testutil.PerformRequest(router, http.MethodPost, "/logout", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestRefreshTokenSuccess(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := testutil.PerformRequest(router, http.MethodPost, "/refresh", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}
