package passgen

import (
	"reflect"
	"testing"
)

func TestNewPasswordInitializer(t *testing.T) {
	inPass := &Password{
		8,
		false,
		"test",
	}

	outPass := newPassword(8, false, "test")
	if !reflect.DeepEqual(inPass, outPass) {
		t.Fatalf("Expected returned new password (%v) to match %v", outPass, inPass)
	}

}

func TestNewPassInvalidLen(t *testing.T) {
	newPass, err := GenPass(0, false, "")
	if err != errInvalidPassLen {
		t.Fatal("Expected error got nil")
	}
	if newPass != "" {
		t.Fatalf("Expected returned password to be nil, got %s", newPass)
	}
}

func TestNewPassValidLen(t *testing.T) {
	newPass, err := GenPass(8, false, "")
	if err != nil {
		t.Fatalf("Expected error to be nil, got %v", err)
	}

	if len(newPass) < 8 {
		t.Fatalf("Expected error to be 8 characters long, got %d characters", len(newPass))
	}
}

func TestNewTokenInvalidLen(t *testing.T) {
	newToken, err := GenPass(0, true, "")
	if err != errInvalidTokenLen {
		t.Fatalf("Expected error (%v) got %v", errInvalidTokenLen, err)
	}

	if newToken != "" {
		t.Fatalf("Expected new token to be empty, got %s", newToken)
	}
}
