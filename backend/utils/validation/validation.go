package validation

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"regexp"
)

func InputValidator(inputData interface{}) (bool, map[string]string) {
	// create new validator instance and register custom validations
	validation := validator.New()
	err := validation.RegisterValidation("custom_email", isValidEmail)
	if err != nil {
		return false, nil
	}

	// validate input data
	err = validation.Struct(inputData)
	if err != nil {
		// if there is an error, create a map to hold fieldname and error message
		valErrors := make(map[string]string)

		// iterate through errors, return appropriate error message depending on the field with the error
		for _, vErr := range err.(validator.ValidationErrors) {
			// get field name and tag
			fieldName := vErr.StructField()
			tag := vErr.Tag()

			// return appropriate error message depending on the tage
			switch tag {
			case "required":
				valErrors[fieldName] = fmt.Sprintf("%s is required", fieldName)
			case "custom_email":
				valErrors[fieldName] = fmt.Sprintf("%s is not a valid email address", fieldName)
			case "min":
				valErrors[fieldName] = fmt.Sprintf("%s must be at least %s characters", fieldName, vErr.Param())
			case "max":
				valErrors[fieldName] = fmt.Sprintf("%s cannot be more than %s characters", fieldName, vErr.Param())
			default:
				valErrors[fieldName] = fmt.Sprintf("%s is invalid", fieldName)
			}
		}
		// return errors if any
		return false, valErrors
	}
	return true, nil
}
func ValidateRequest(w http.ResponseWriter, r *http.Request, inputData interface{}) error {
	// decode request
	err := json.NewDecoder(r.Body).Decode(inputData)
	if err != nil {
		logrus.WithError(err).Logger.Error("error decoding json")
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return fmt.Errorf("invalid JSON payload")
	}

	// call func to validate fields
	isValid, validationErrors := InputValidator(inputData)
	if !isValid {
		// return error if request is invalid
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"errors":  validationErrors,
		})
		if err != nil {
			return err
		}
		return fmt.Errorf("validation failed")
	}

	return nil
}

func isValidEmail(fl validator.FieldLevel) bool {
	// validate email fields
	email := fl.Field().String()

	// regex for a valid email address
	re := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}
