package auth

func ValidateToken(tokenString string) bool {
	return tokenString == "Bearer 210401010174"
}
