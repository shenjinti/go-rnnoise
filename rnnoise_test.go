package gornnoise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	r := NewRNNoise()

	frameSize := GetFrameSize()
	assert.Equal(t, frameSize, 480)
	o, vadProb := r.Process([]byte{1, 2, 3, 4, 5, 6})
	assert.Greater(t, vadProb, float32(0.0))
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0}, o)
}
