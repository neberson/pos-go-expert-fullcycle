package entity

import (
	"testing"
)

func TestNewCep(t *testing.T) {
	cep := NewCep("12345678")
	if cep == nil {
		t.Error("Expected cep to not be nil")
	}
	if cep.Cep != "12345678" {
		t.Errorf("Expected cep to be 12345678, got %s", cep.Cep)
	}
}

func TestCep_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cep     string
		wantErr bool
	}{
		{
			name:    "valid cep with 8 digits",
			cep:     "12345678",
			wantErr: false,
		},
		{
			name:    "valid cep with different digits",
			cep:     "01310100",
			wantErr: false,
		},
		{
			name:    "invalid cep with less than 8 digits",
			cep:     "1234567",
			wantErr: true,
		},
		{
			name:    "invalid cep with more than 8 digits",
			cep:     "123456789",
			wantErr: true,
		},
		{
			name:    "invalid cep with letters",
			cep:     "1234567a",
			wantErr: true,
		},
		{
			name:    "invalid cep with special characters",
			cep:     "12345-678",
			wantErr: true,
		},
		{
			name:    "invalid empty cep",
			cep:     "",
			wantErr: true,
		},
		{
			name:    "invalid cep with spaces",
			cep:     "123 456 78",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCep(tt.cep)
			err := c.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != ErrInvalidCep {
				t.Errorf("Validate() error = %v, want %v", err, ErrInvalidCep)
			}
		})
	}
}
