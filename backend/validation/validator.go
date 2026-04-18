package validation

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Init registers custom validators used by request DTOs.
func Init() error {
	engine := binding.Validator.Engine()
	v, ok := engine.(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to load validator engine")
	}

	return v.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		role := strings.ToLower(fl.Field().String())
		switch role {
		case "student", "staff", "admin":
			return true
		default:
			return false
		}
	})
}
