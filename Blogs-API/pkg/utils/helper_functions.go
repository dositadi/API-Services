package utils

import (
	m "blog/pkg/models"

	b "golang.org/x/crypto/bcrypt"
)

const (
	LONG_PASS_ERR        = `Password is too long.`
	LONG_PASS_ERR_DETAIL = `The password you entered is too long, consider shortening it but not making it too simple.`

	SHORT_PASS_ERR        = `Password is too short.`
	SHORT_PASS_ERR_DETAIL = `The password you entered is too short. Password should be 8 characters long.`

	PASS_MISMATCH_ERR        = `Password mismatch.`
	PASS_MISMATCH_ERR_DETAIL = `The password you entered is incorrect.`
)

func SaltAndHashPassword(password string) (string, *m.ErrorMessage) {
	saltAndHashPassword, err := b.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		if err == b.ErrPasswordTooLong {
			return "", &m.ErrorMessage{
				Error:   LONG_PASS_ERR,
				Details: []string{LONG_PASS_ERR_DETAIL},
				Code:    "",
			}
		}
	}
	return string(saltAndHashPassword), nil
}

func ComparePasswordAndHashed(hashedPassword, password string) *m.ErrorMessage {
	if err := b.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if err.Error() == b.ErrHashTooShort.Error() {
			return &m.ErrorMessage{
				Error:   SHORT_PASS_ERR,
				Details: []string{SHORT_PASS_ERR_DETAIL},
				Code:    "",
			}
		}

		if err == b.ErrMismatchedHashAndPassword {
			return &m.ErrorMessage{
				Error:   PASS_MISMATCH_ERR,
				Details: []string{PASS_MISMATCH_ERR_DETAIL},
				Code:    "",
			}
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
