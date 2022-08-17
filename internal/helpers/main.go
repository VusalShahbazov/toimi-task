package helpers

func Get(val, def string) string {
	if val != "" {
		return val
	}
	return def
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
