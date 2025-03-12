package validations

import (
	"fmt"
	"regexp"
	"standard_app/models"
)

type User struct {}

// ValidateEmail checks whether the email is in a valid format
func ValidateEmail(email string) error {
	// Simple email regex (not comprehensive, just an example)
	re := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(re).MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// ValidateName checks whether the name contains only alphabetic characters and spaces
func ValidateName(name string) error {
	re := `^[A-Za-z\s]+$`
	if !regexp.MustCompile(re).MatchString(name) {
		return fmt.Errorf("name can only contain letters and spaces")
	}
	return nil
}

// ValidateAge checks if the age is at least 18
func ValidateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age must be at least 18")
	}
	return nil
}

// ValidateUser validates all fields and returns a slice of errors
func (u * User) ValidateUser(user models.User) []string {
	var errors []string

	// Check each field and add validation errors to the slice if necessary
	if err := ValidateEmail(user.Email); err != nil {
		errors = append(errors, err.Error())
	}

	if err := ValidateName(user.Name); err != nil {
		errors = append(errors, err.Error())
	}

	if err := ValidateAge(user.Age); err != nil {
		errors = append(errors, err.Error())
	}

	return errors
}
