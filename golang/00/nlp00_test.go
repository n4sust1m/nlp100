package nlp00

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "reverse \"stressed\"",
			args: "stressed",
			want: "desserts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
