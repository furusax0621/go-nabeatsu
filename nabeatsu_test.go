package nabeatsu

import (
	"testing"
)

func TestIsFool(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "3の倍数",
			input: "1221",
			want:  true,
		},
		{
			name:  "3の倍数ではない",
			input: "1222",
			want:  false,
		},
		{
			name:  "3の倍数ではないが3を含む",
			input: "13",
			want:  true,
		},
		{
			name:  "3の倍数で3を含む",
			input: "33",
			want:  true,
		},
		{
			name:  "数字ではない",
			input: "nabeatsu",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFool(tt.input); got != tt.want {
				t.Errorf("IsFool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFoolExpression(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "3の倍数以外",
			input: "1",
			want:  "1",
		},
		{
			name:  "阿呆になる数字(1桁)",
			input: "3",
			want:  "ｻｧﾝwww",
		},
		{
			name:  "阿呆になる数字(2桁)",
			input: "24",
			want:  "ﾆｼﾞｭｳﾖﾝwww",
		},
		{
			name:  "阿呆になる数字(3桁)",
			input: "489",
			want:  "ﾖﾝﾋｬｸﾊﾁｼﾞｭｳｷｭｳwww",
		},
		{
			name:  "阿呆になる数字(4桁)",
			input: "6570",
			want:  "ﾛｸｾﾝｺﾞﾋｬｸﾅﾅｼﾞｭｳwww",
		},
		{
			name:  "接頭辞の読み方が変わる場合(千, 百)",
			input: "3300",
			want:  "ｻｧﾝｾﾞﾝｻｧﾝﾋﾞｬｸwww",
		},
		{
			name:  "1は発音しない",
			input: "1110",
			want:  "ｾﾝﾋｬｸｼﾞｭｳwww",
		},
		{
			name:  "接頭辞、数字両方の読み方が変わる場合(六百)",
			input: "600",
			want:  "ﾛｯﾋﾟｬｸwww",
		},
		{
			name:  "接頭辞、数字両方の読み方が変わる場合(八千八百)",
			input: "8802",
			want:  "ﾊｯｾﾝﾊｯﾋﾟｬｸﾆwww",
		},
		{
			name:  "追加の接頭辞がつく数字",
			input: "865433100001", // 865,433,100,001 (八千六百五十四億 三千三百十万 一)
			want:  "ﾊｯｾﾝﾛｯﾋﾟｬｸｺﾞｼﾞｭｳﾖﾝｵｸｻｧﾝｾﾞﾝｻｧﾝﾋﾞｬｸｼﾞｭｳﾏﾝｲﾁwww",
		},
		{
			name:  "1,000のときは特別な読み方をする",
			input: "181001000003", // 181,001,000,003 (一千八百十億 百万 三)
			want:  "ｲｯｾﾝﾊｯﾋﾟｬｸｼﾞｭｳｵｸﾋｬｸﾏﾝｻｧﾝwww",
		},
		{
			name:  "ゼロ埋めの接頭辞は省略する",
			input: "1000000002", // 1,000,000,003 (十億 二)
			want:  "ｼﾞｭｳｵｸﾆwww",
		},

		{
			name:  "スペック外の数字",
			input: "1000000000000000000000000000000000000000000000000000000000000000000000002",
			want:  "ﾑｹﾞﾝﾀﾞｧｲwww",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFoolExpression(tt.input); got != tt.want {
				t.Errorf("BuildFoolMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
