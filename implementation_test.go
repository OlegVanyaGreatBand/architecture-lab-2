package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixComputeCorrect(t *testing.T) {
	// 5 = 5
	res, err := ComputePrefix("5")
	if assert.Nil(t, err) {
		assert.Equal(t, 5.0, res)
	}

	// 5 + 3 * (4 - 2) = 11
	res, err = ComputePrefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, 11.0, res)
	}

	// 10.25 * 100 / (5 - 6)^3 = −1025
	res, err = ComputePrefix("* 10.25 / 100 ^ - 5 6 3")
	if assert.Nil(t, err) {
		assert.Equal(t, -1025.0, res)
	}

	// 1 × 10 + 2 × 20 + 3 × 30 + 4 × 40 + 5 × 50 = 550
	res, err = ComputePrefix("+ * 1 10 + * 2 20 + * 3 30 + * 4 40 * 5 50")
	if assert.Nil(t, err) {
		assert.Equal(t, 550.0, res)
	}
}

func TestPrefixComputeIncorrect(t *testing.T) {
	// + - syntax error
	_, err := ComputePrefix("+")
	assert.NotNil(t, err)

	// + 5 - syntax error
	_, err = ComputePrefix("+ 5")
	assert.NotNil(t, err)

	// 5 + - syntax error
	_, err = ComputePrefix("5 +")
	assert.NotNil(t, err)

	// "" - syntax error
	_, err = ComputePrefix("")
	assert.NotNil(t, err)

	// "" - syntax error
	_, err = ComputePrefix("abc")
	assert.NotNil(t, err)
}

func ExampleComputePrefix() {
	// example of prefix expression, returning res and error
	res, _ := ComputePrefix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 4
}
