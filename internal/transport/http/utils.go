package http

import "regexp"

func isValidEmail(email string) bool {
	// A simple email validation regex (can be refined further as needed)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
