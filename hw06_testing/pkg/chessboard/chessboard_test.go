package chessboard

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		want    string
		wantErr bool
	}{
		{"Допустимый размер доски 3", 3, "# #\n # \n# #\n", false},
		{"Допустимый размер доски 4", 4, "# # \n # #\n# # \n # #\n", false},
		{"Недопустимый размер доски 0", 0, "", true},
		{"Недопустимый размер доски -1", -1, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateBoard(tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateBoard() got = %v, want %v", got, tt.want)
			}
		})
	}
}
