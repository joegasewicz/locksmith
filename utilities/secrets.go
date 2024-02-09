package utilities

import "os"

func GetTokenSecret() string {
	return os.Getenv("TOKEN_SECRET")
}
