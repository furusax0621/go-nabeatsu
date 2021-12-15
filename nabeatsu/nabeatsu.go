package nabeatsu

import (
	"regexp"
	"strings"
)

var mustNumber = regexp.MustCompile(`^\d+$`)

// IsFool 阿呆になる数字かどうかを判定します。世界のナベアツは3の倍数と3のつく数字のときだけ阿呆になります
func IsFool(s string) bool {
	// 数字以外の何かが入ってる
	if !mustNumber.MatchString(s) {
		return false
	}

	// 3 の付く数字
	if strings.Contains(s, "3") {
		return true
	}

	// 3 の倍数
	var sum uint64
	for _, r := range s {
		sum += uint64(r - '0')
	}

	return sum%3 == 0
}
