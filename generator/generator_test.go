package generator_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/sifer169966/unique-code-generator/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	quantityConfig := generator.Generator{
		Quantity: 100,
	}
	rand.Seed(time.Now().UnixNano())
	generator := generator.NewGenerator()
	codes := generator.GenerateCodes(quantityConfig.Quantity)
	assert.Equal(t, 100, len(codes))
	for _, code := range codes {
		assert.Equal(t, 10, len(code))
	}
}
