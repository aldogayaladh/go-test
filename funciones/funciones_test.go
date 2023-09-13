package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Uso de casos en test.
func Test_Suma(t *testing.T) {

	t.Run("La suma es 5", func(t *testing.T) {
		// Arrange.
		var a, b int
		a = 2
		b = 3
		expectedResult := 5

		// Act.
		result := Suma(a, b)

		// Assert.
		assert.Equal(t, expectedResult, result)

	})

	t.Run("La suma es 6", func(t *testing.T) {
		// Arrange.
		var a, b int
		a = 3
		b = 3
		expectedResult := 6

		// Act.
		result := Suma(a, b)

		// Assert.
		assert.Equal(t, expectedResult, result)
	})

	t.Run("La suma no es la esperada", func(t *testing.T) {
		// Arrange.
		var a, b int
		a = 2
		b = 3
		expectedResult := 6

		// Act.
		result := Suma(a, b)

		// Assert.
		assert.NotEqual(t, expectedResult, result)
	})

}

func Test_Division(t *testing.T) {
	t.Run("Happy path. Division OK", func(t *testing.T) {
		// Arrange.
		var a, b int
		a = 10
		b = 2
		expectedResult := 5

		// Act.
		result, err := Division(a, b)

		// Assert.
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error path. Division Error por cero", func(t *testing.T) {
		// Arrange.
		a := 10
		b := 0
		expectedResult := 0

		// Act.
		result, err := Division(a, b)

		// Assert.
		assert.Equal(t, expectedResult, result)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrDivisionZero)
	})
}

// Uso de test table.
func Test_SumaTable(t *testing.T) {
	// Arrange.
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name           string
		args           args
		expectedResult int
	}{
		// Casos que vamos a testear
		{
			name:           "La suma es 4",
			args:           args{a: 2, b: 2},
			expectedResult: 4,
		},
		{
			name:           "La suma es 5",
			args:           args{a: 2, b: 3},
			expectedResult: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedResult, Suma(test.args.a, test.args.b))
		})
	}
}
