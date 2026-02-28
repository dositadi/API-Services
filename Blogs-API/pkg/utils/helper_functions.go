package utils

import (
	m "blog/pkg/models"
	"errors"
	"regexp"

	b "golang.org/x/crypto/bcrypt"
)

func SaltAndHashPassword(password string) (string, *m.ErrorMessage) {
	if len([]byte(password)) > 50 {
		return "", &m.ErrorMessage{
			Error:   LONG_PASS_ERR,
			Details: []string{LONG_PASS_ERR_DETAIL},
			Code:    "",
		}
	}

	saltAndHashPassword, err := b.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", &m.ErrorMessage{
			Error:   INVALID_PASSWORD_ERROR,
			Details: []string{INVALID_PASSWORD_ERROR_DETAIL},
			Code:    "",
		}
	}
	return string(saltAndHashPassword), nil
}

func ComparePasswordAndHashed(hashedPassword, password string) *m.ErrorMessage {
	if err := b.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if errors.Is(err, b.ErrHashTooShort) {
			return &m.ErrorMessage{
				Error:   SERVER_ERROR,
				Details: []string{SERVER_ERROR_DETAIL},
				Code:    SERVER_ERROR_CODE,
			}
		}

		if errors.Is(err, b.ErrMismatchedHashAndPassword) {
			return &m.ErrorMessage{
				Error:   PASS_MISMATCH_ERR,
				Details: []string{PASS_MISMATCH_ERR_DETAIL},
				Code:    "",
			}
		}
		return &m.ErrorMessage{
			Error:   PASS_MISMATCH_ERR,
			Details: []string{PASS_MISMATCH_ERR_DETAIL},
			Code:    "",
		}
	}
	return nil
}

func HashedCost(hashedPassword string) (int, error) {
	cost, err := b.Cost([]byte(hashedPassword))
	if err != nil {
		return 0, err
	}
	return cost, nil
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-\._]+)@`)

	return re.Match([]byte(email))
}


