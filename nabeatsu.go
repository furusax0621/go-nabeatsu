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

// GetFoolExpression 3の倍数と3のつく数字のときだけ阿呆になった数字の読み方をします。それ以外のケースではオリジナルの文字列を返します。
func GetFoolExpression(s string) string {
	if len(s) >= 73 { // 無量大数より上の単位は存在しないので無限大ということにしておく
		return "ﾑｹﾞﾝﾀﾞｧｲwww"
	}
	if !IsFool(s) {
		return s
	}

	// 最後に草を生やしたい
	suffix := "www"

	// 10,000未満
	if len(s) <= 4 {
		return buildFoolMessageUnder10Thousand(s) + suffix
	}
	msg := buildFoolMessageUnder10Thousand(s[len(s)-4:])
	s = s[:len(s)-4]

	// 以降、適切な接頭辞を追加しながら構築
	for i := 0; i < len(prefix) && len(s) > 0; i++ {
		p := prefix[i]
		offset := len(s) - 4
		if offset < 0 {
			offset = 0
		}

		msg = buildFoolMessageWithHighPrefix(s[offset:], p) + msg
		s = s[:offset]
	}

	return msg + suffix
}

func buildFoolMessageUnder10Thousand(s string) string {
	var msg string
	for i := 1; i <= len(s); i++ {
		d := s[len(s)-i]
		n := number[d]

		// 0 は発音しない
		if d == '0' {
			continue
		}

		// 十の位以降では 1 は発音しない
		if i > 1 && d == '1' {
			n = ""
		}

		var p string
		switch i {
		case 2: // 十の位
			p = "ｼﾞｭｳ"

		case 3: // 百の位
			p = "ﾋｬｸ"

			// 3のときは接頭辞の読み方が変わる
			if d == '3' {
				p = "ﾋﾞｬｸ"
			}

			// 6のときは特別な読み方をする
			if d == '6' {
				n = "ﾛｯ"
				p = "ﾋﾟｬｸ"
			}

			// 8のときは特別な読み方をする
			if d == '8' {
				n = "ﾊｯ"
				p = "ﾋﾟｬｸ"
			}

		case 4: // 千の位
			p = "ｾﾝ"

			// 3のときは接頭辞の読み方が変わる
			if d == '3' {
				p = "ｾﾞﾝ"
			}

			// 8のときは特別な読み方をする
			if d == '8' {
				n = "ﾊｯ"
			}

		}

		msg = n + p + msg
	}

	return msg
}

func buildFoolMessageWithHighPrefix(s, prefix string) string {
	var msg string
	for i := 1; i <= len(s); i++ {
		d := s[len(s)-i]
		n := number[d]

		// 0 は発音しない
		if d == '0' {
			continue
		}

		var p string
		switch i {
		case 2: // 十の位
			p = "ｼﾞｭｳ"

			// 1 は発音しない
			if d == '1' {
				n = ""
			}

		case 3: // 百の位
			p = "ﾋｬｸ"

			// 1 は発音しない
			if d == '1' {
				n = ""
			}

			// 3のときは接頭辞の読み方が変わる
			if d == '3' {
				p = "ﾋﾞｬｸ"
			}

			// 6のときは特別な読み方をする
			if d == '6' {
				n = "ﾛｯ"
				p = "ﾋﾟｬｸ"
			}

			// 8のときは特別な読み方をする
			if d == '8' {
				n = "ﾊｯ"
				p = "ﾋﾟｬｸ"
			}

		case 4: // 千の位
			p = "ｾﾝ"

			// 1 の読み方が変わる
			if d == '1' {
				n = "ｲｯ"
			}

			// 3のときは接頭辞の読み方が変わる
			if d == '3' {
				p = "ｾﾞﾝ"
			}

			// 8のときは特別な読み方をする
			if d == '8' {
				n = "ﾊｯ"
			}

		}

		msg = n + p + msg
	}

	// 表現すべき数字が何もなかった場合 (0 埋めのケース)
	if msg == "" {
		return ""
	}

	return msg + prefix
}

var number = map[byte]string{
	'0': "",
	'1': "ｲﾁ",
	'2': "ﾆ",
	'3': "ｻｧﾝ",
	'4': "ﾖﾝ",
	'5': "ｺﾞ",
	'6': "ﾛｸ",
	'7': "ﾅﾅ",
	'8': "ﾊﾁ",
	'9': "ｷｭｳ",
}

var prefix = []string{
	"ﾏﾝ",
	"ｵｸ",
	"ﾁｮｳ",
	"ｹｲ",
	"ｶﾞｲ",
	"ｼﾞｮ",
	"ｼﾞｮｳ",
	"ｺｳ",
	"ｶﾝ",
	"ｾｲ",
	"ｻｲ",
	"ｺﾞｸ",
	"ｺﾞｳｶﾞｼｬ",
	"ｱｿｳｷﾞ",
	"ﾅﾕﾀ",
	"ﾌｶｼｷﾞ",
	"ﾑﾘｮｳﾀｲｽｳ",
}
