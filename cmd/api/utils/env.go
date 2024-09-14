package utils

func IsDev(env string) bool {
	return env == "DEV"
}

func IsProd(env string) bool {
	return env == "PRODUCTION"
}