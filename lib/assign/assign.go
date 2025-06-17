package assign

func String(payloadString string, defaultStr string) string {
	if payloadString == "" {
		return defaultStr
	}
	return payloadString
}
