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

func TestBuildFoolMessage(t *testing.T) {
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
			want:  "ｻｧﾝ",
		},
		{
			name:  "阿呆になる数字(2桁)",
			input: "24",
			want:  "ﾆｼﾞｭｳﾖﾝ",
		},
		{
			name:  "阿呆になる数字(3桁)",
			input: "489",
			want:  "ﾖﾝﾋｬｸﾊﾁｼﾞｭｳｷｭｳ",
		},
		{
			name:  "阿呆になる数字(4桁)",
			input: "6570",
			want:  "ﾛｸｾﾝｺﾞﾋｬｸﾅﾅｼﾞｭｳ",
		},
		{
			name:  "接頭辞の読み方が変わる場合(千, 百)",
			input: "3300",
			want:  "ｻｧﾝｾﾞﾝｻｧﾝﾋﾞｬｸ",
		},
		{
			name:  "1は発音しない",
			input: "1110",
			want:  "ｾﾝﾋｬｸｼﾞｭｳ",
		},
		{
			name:  "接頭辞、数字両方の読み方が変わる場合(六百)",
			input: "600",
			want:  "ﾛｯﾋﾟｬｸ",
		},
		{
			name:  "接頭辞、数字両方の読み方が変わる場合(八千八百)",
			input: "8802",
			want:  "ﾊｯｾﾝﾊｯﾋﾟｬｸﾆ",
		},

		// TODO: 順次対応していく
		{
			name:  "スペック外の数字",
			input: "10000",
			want:  "現在ﾅﾍﾞｱﾂは10,000以上の数を計算できません",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildFoolMessage(tt.input); got != tt.want {
				t.Errorf("BuildFoolMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
