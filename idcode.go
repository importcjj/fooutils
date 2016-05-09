package fooutils

import "strconv"

var (
	Factories    [17]int = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	ValidateBits []byte  = []byte("10X98765432")
)

func IsValidIDCode(id string) bool {
	idBytes := []byte(id)
	if length := len(idBytes); length != 18 {
		return false
	}

	sum := 0
	lastByte := idBytes[17]
	idBytes = idBytes[:17]

	for i, b := range idBytes {
		value, err := strconv.Atoi(string(b))
		if err != nil {
			return false
		}
		sum += value * Factories[i]
	}

	index := sum % 11
	if ValidateBits[index] == lastByte {
		return true
	}

	return false
}
