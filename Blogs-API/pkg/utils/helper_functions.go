package utils

import b "golang.org/x/crypto/bcrypt"

func SaltAndHashPassword(password string) (string, error) {
	saltAndHashPassword, err := b.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}
	return string(saltAndHashPassword), nil
}

func ComparePasswordAndHashed(hashedPassword, password string) error {
	if err := b.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
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
