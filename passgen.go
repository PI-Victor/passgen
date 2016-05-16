package passgen

import "errors"

var (
	alpLower           = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alpUpper           = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	specialChar        = []byte{'`', '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '-', '+', '=', '{', '}', '[', ']', '\\', '|', ':', ';', '<', '>', ',', '.', '?', '/', '\'', '"'}
	numbers            = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	errInvalidPassLen  = errors.New("Password must be at least 8 characters!")
	errInvalidTokenLen = errors.New("Token must be at least 56 characters long")
)

func newPassword(len int, isToken bool, baseKey string) *password {
	return &password{
		len,
		isToken,
		baseKey,
	}
}

type password struct {
	len     int
	isToken bool
	baseKey string
}

func (p *password) generatePassword() (newPass string, err error) {
	return "", nil
}

func (p *password) generateToken() (newToken string, err error) {
	return "", nil
}

// GenPass generates a new password or token, if isToken is set to true it will
// generate a new min 56 character token of alphabet characters and numbers
// without using special characters or the base key.
func GenPass(passwordLen int, isToken bool, baseKey string) (password string, err error) {
	if isToken && passwordLen < 56 {
		return "", errInvalidTokenLen
	}

	if !isToken && passwordLen < 8 {
		return "", errInvalidPassLen
	}

	newPass := newPassword(passwordLen, isToken, baseKey)

	if isToken {
		password, err = newPass.generateToken()
		if err != nil {
			return "", err
		}
		return password, nil
	}

	password, err = newPass.generatePassword()
	if err != nil {
		return "", err
	}

	return password, nil
}
