package passgen

// Password is the main password struct.
type Password struct {
	len int
}

var (
	alpLower    = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alpUpper    = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	specialChar = []byte{'?', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '=', '+', '_'}
	numbers     = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

func newPassword(len int) *Password {
	return &Password{
		len: len,
	}
}

func (p *Password) genPass(len int) error {
	return nil
}
