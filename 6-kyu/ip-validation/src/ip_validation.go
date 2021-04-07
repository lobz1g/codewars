package src

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	nums := strings.Split(ip, ".")
	if len(nums) != 4 {
		return false
	}
	for _, value := range nums {
		n, err := strconv.Atoi(value)
		if err != nil || n < 0 || n > 255 || len(value) != len(strconv.Itoa(n)) {
			return false
		}
	}

	return true
}
