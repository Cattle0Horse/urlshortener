package hasher

import "testing"

func TestHashPassword(t *testing.T) {
	tests := []string{"abc", "sdfasd", "213asdas", "csdaca", "0342342"}

	hasher := NewPasswordHash()

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			hash, err := hasher.HashPassword(tt)
			if err != nil {
				t.Error(err)
			}

			if !hasher.ComparePassword(hash, tt) {
				t.Error("must return true")
			}
		})
	}
}
