package password_hash

import "golang.org/x/crypto/bcrypt"

type HashService struct {
	cost int
}

func (hashService *HashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashService.cost)
	if err != nil {
		return "", &HashServiceError{err.Error()}
	}
	return string(bytes), err
}

func (hashService *HashService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewHashService(cost int) HashService {
	return HashService{cost: cost}
}
