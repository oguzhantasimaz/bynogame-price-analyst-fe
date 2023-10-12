package swa

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
