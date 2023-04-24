package shipments

import "strings"

// SS returns the Suitability Score based on the shipment destination
// and driver's name.
func SS(shipmentDestination, driversName string) float64 {
	var ss float64
	if isEven(stringLength(shipmentDestination)) {
		ss = float64(countVowels(driversName)) * 1.0
	} else {
		ss = float64(stringLength(driversName)-countVowels(driversName)) * 1.0
	}

	gcf := GCF(
		stringLength(shipmentDestination),
		stringLength(driversName),
	)
	if gcf > 1 {
		ss *= 1.5
	}
	return ss
}

// stringLength returns the length of a given string
func stringLength(name string) int {
	letters := []rune(name)
	return len(letters)
}

// isEven returns true id the given number is even, else false
func isEven(num int) bool {
	return num%2 == 0
}

func countVowels(inputString string) int {
	inputString = strings.ToLower(inputString)
	vowels := "aeiou"
	outputString := strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return -1
		}
		return r
	}, inputString)
	return stringLength(outputString)
}

// GCF is an Euclidean algorithm implementation to find the Greatest Common Factor
// of two given numbers besides 1
func GCF(a, b int) int {
	for {
		if a > 0 && b > 0 {
			if a > b {
				a = a % b
			} else {
				b = b % a
			}
		} else {
			break
		}
	}
	if a == 0 {
		return b
	}
	return a
}
