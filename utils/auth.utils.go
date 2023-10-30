package utils

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string) (string, error) {
    // Generate a salt with a cost of 12 (you can adjust the cost as needed)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}