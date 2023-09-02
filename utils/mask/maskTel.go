package mask

func MaskTel(phoneNumber string) string {
	if len(phoneNumber) != 11 {
		return "Invalid phone number"
	}

	return phoneNumber[:3] + "****" + phoneNumber[7:]
}
