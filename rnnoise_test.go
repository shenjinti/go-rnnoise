package gornnoise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	r := NewRNNoise()

	o, vadProb := r.Process([]byte{1, 2, 3, 4, 5, 6})
	assert.Greater(t, vadProb, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0}, o)
}
