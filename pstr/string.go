package pstr

import (
	"math/rand"
	"strconv"
	"time"
)

func In_array(find string, s []string) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == string(find) {
			return true
		}
	}
	return false
}

// Convert the string to a number
func Intval(s string) int {

	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// To disrupt the array
func Shuffle(sz []string) []string {
	if len(sz) <= 0 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(sz) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		tmp := sz[i]
		sz[i] = sz[j]
		sz[j] = tmp
	}
	return sz
}
