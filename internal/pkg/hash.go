package pkg

import "golang.org/x/crypto/bcrypt"

// Decode make hash from password
func Decode(hash, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err
}

// Encode compare hash and password and then throw an error
func Encode(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err
	}
	return hashedPassword, nil
}
