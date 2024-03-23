package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	result, err := HashPassword("SuperSecret")

	if err != nil {
		t.Errorf("Error thrown: %v, result: %v", err, result)
	}

	if result == "" {
		t.Error("No hashed password returned")
	}
}

func TestAuthenticatePassword(t *testing.T) {
	hashResult, err := HashPassword("SuperSecret")

	if err != nil {
		t.Errorf("Error hashing password thrown: %v, result: %v", err, hashResult)
	}

	if hashResult == "" {
		t.Error("No hashed password returned")
	}

	result := AuthenticatePassword("SuperSecret", hashResult)

	if !result {
		t.Error("Error matching password and hash do not match")
	}
}

func TestAuthenticatePasswordTableDriven(t *testing.T) {
	// Define the columns of the table
	var tests = []struct {
		name          string
		password      string
		checkPassword string
		expected      bool
	}{
		// The table
		{"Correct Password", "SuperSecret123!", "SuperSecret123!", true},
		{"Incorrect Password", "SuperSecret123!", "!123SuperSecret", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashResult, err := HashPassword(tt.password)

			if err != nil {
				t.Errorf("Error hashing password thrown: %v, result: %v", err, hashResult)
			}

			if hashResult == "" {
				t.Error("No hashed password returned")
			}

			result := AuthenticatePassword(tt.checkPassword, hashResult)

			if result != tt.expected {
				t.Errorf("Error result '%v' didn't match expected result '%v'", result, tt.expected)
			}
		})
	}
}

func TestGenerateToken(t *testing.T) {
	result, err := GenerateToken("test@email.com", "9f4f1408-d665-4f25-a252-fe9b154f337b")

	if err != nil {
		t.Errorf("Error generating token thrown: %v, result: %v", err, result)
	}

	if result == "" {
		t.Error("No token returned")
	}
}

func TestVerifyToken(t *testing.T) {
	tokenResult, err := GenerateToken("test@email.com", "9f4f1408-d665-4f25-a252-fe9b154f337b")

	if err != nil {
		t.Errorf("Error generating token thrown: %v, result: %v", err, tokenResult)
	}

	if tokenResult == "" {
		t.Error("No token returned")
	}

	result, err := VerifyToken(tokenResult)

	if err != nil {
		t.Errorf("Error verifying token thrown: %v, result: %v", err, result)
	}

	if result == "" {
		t.Error("No token returned")
	}
}

func TestGenerateUUID(t *testing.T) {
	result, err := GenerateUUID()

	if err != nil {
		t.Errorf("Error thrown: %v, result: %v", err, result)
	}

	if result == "" {
		t.Error("No ID was returned")
	}
}
