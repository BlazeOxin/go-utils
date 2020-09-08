package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// GenerateRandomHEX ... generates a hex between 6(a) to 6(f)
func GenerateRandomHEX(size int) string {
	a := ""
	f := ""
	for i := 0; i < size; i++ {
		a += "a"
		f += "f"
	}
	min, _ := strconv.ParseInt(a, 16, 64)
	max, _ := strconv.ParseInt(f, 16, 64)
	number := rand.Int63n(max-min) + min
	return strings.Replace(fmt.Sprintf("%X", number), "0x", "", -1)
}
