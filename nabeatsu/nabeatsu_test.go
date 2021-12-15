package nabeatsu

import "testing"

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
