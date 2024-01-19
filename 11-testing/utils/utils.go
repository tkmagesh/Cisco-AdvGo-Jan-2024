package utils

var results map[string]int

func init() {
	results = make(map[string]int)
}

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			results["non-prime"]++
			return false
		}
	}
	results["prime"]++
	return true
}
