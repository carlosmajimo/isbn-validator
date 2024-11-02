package isbn

import "testing"

func TestValidateISBN(t *testing.T) {
	tests := []struct {
		name    string
		isbn    string
		want    bool
		wantErr bool
	}{
		{
			name: "checkInvalidISBN",
			isbn: "0140441927",
			want: false,
		},
		{
			name: "checkISBNLengthTen",
			isbn: "0471958697",
			want: true,
		},
		{
			name: "checkISBNLengthThirteen",
			isbn: "9780470059029",
			want: true,
		},
		{
			name: "checkISBNWrongLength",
			isbn: "97804700590293",
			want: false,
		},
		{
			name: "checkISBNNumeric",
			isbn: "helloworld",
			want: false,
		},
		{
			name: "checkContainsX",
			isbn: "080442957X",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidateISBN(tt.isbn)
			got := validator.IsValid()
			if got != tt.want {
				t.Errorf("ValidateISBN.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
