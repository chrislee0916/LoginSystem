package utils

import "golang.org/x/crypto/bcrypt"

//使用Bcrypt 慢雜湊演算法加密password
func Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// 判斷登入的password與DB裡加密過的password是否相同
func Compare(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	return nil
}
