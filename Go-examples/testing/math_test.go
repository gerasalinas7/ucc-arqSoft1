package mathutils

import "testing"

// TestSuma verifica que la función Sumar devuelva la suma correcta
func TestSuma(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{10, 20, 30},
		{-1, -1, -2},
		{0, 0, 0},
		{-5, 5, 0},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			resultado := Sumar(test.a, test.b)
			if resultado != test.expected {
				t.Errorf("Suma de %d y %d esperada %d, pero obtuvo %d", test.a, test.b, test.expected, resultado)
			}
		})
	}
}

// TestMayor verifica que la función Mayor devuelva el número mayor correctamente
func TestMayor(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{10, 20, 20},
		{-1, -2, -1},
		{0, 0, 0},
		{5, 5, 5},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			resultado := Mayor(test.a, test.b)
			if resultado != test.expected {
				t.Errorf("Mayor de %d y %d esperada %d, pero obtuvo %d", test.a, test.b, test.expected, resultado)
			}
		})
	}
}
