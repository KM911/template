package util

func MaxLength(s []string) int {
	currentMax := 0
	for _, v := range s {
		if len(v) > currentMax {
			currentMax = len(v)
		}
	}
	return currentMax
}
