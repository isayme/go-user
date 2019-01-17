package password

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const algorithm = "bcrypt"
const separator = "|"

func randomSalt() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

// Generate get hashed password
func Generate(password string) (string, error) {
	salt := randomSalt()

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := strings.Join([]string{algorithm, salt, string(hashBytes)}, separator)
	return hash, nil
}

// Compare compare password
func Compare(hash, password string) bool {
	parts := strings.Split(hash, separator)

	if len(parts) != 3 {
		return false
	}

	// algorithm := parts[0] // current not used
	salt := parts[1]
	hashedPassword := parts[2]
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(salt+password)) == nil
}
