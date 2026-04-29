package validation

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type roleValidationFixture struct {
	Role string `binding:"role"`
}

func TestInitRegistersRoleValidator(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatalf("expected validator init to succeed, got %v", err)
	}

	engine := binding.Validator.Engine()
	validate, ok := engine.(*validator.Validate)
	if !ok {
		t.Fatal("expected validator engine to be available")
	}

	if err := validate.Struct(roleValidationFixture{Role: "student"}); err != nil {
		t.Fatalf("expected supported role to validate, got %v", err)
	}

	if err := validate.Struct(roleValidationFixture{Role: "faculty"}); err == nil {
		t.Fatal("expected unsupported role to fail validation")
	}
}
