package utils

import "regexp"

func IsStrongPassword(password string) bool {
    if len(password) < 8 {
        return false
    }
    if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
        return false
    }
    if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
        return false
    }
    if ok, _ := regexp.MatchString(`[0-9]`, password); !ok {
        return false
    }
    return true
}