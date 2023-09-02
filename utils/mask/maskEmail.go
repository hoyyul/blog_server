package mask

import "strings"

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "Invalid email address"
	}
	localPart, domain := parts[0], parts[1]

	if len(localPart) < 2 || len(domain) < 2 {
		return "Invalid email address"
	}

	maskedLocalPart := string(localPart[0]) + strings.Repeat("*", len(localPart)-1)
	maskedDomain := string(domain[0]) + strings.Repeat("*", len(domain)-1)

	return maskedLocalPart + "@" + maskedDomain
}
