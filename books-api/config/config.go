package config

func DataPath() string {
	return "data.json"
}

func UserPath() string {
	return "user.json"
}

func JWTSecret() []byte {
	var key = []byte("Secret")
	return key
}
