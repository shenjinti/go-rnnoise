# rnnoise golang bindings

- [rnnoise](https://github.com/xiph/rnnoise.git)

The RNNoise project is a noise suppression library based on a recurrent neural network.

This project is a golang bindings for the RNNoise project.

## requirements

- gcc (for cgo)

## get cmd gornnoise

``` bash
$ go install github.com/shenjinti/go-rnnoise/cmd/gornnoise

#Usage: gornnoise <raw/pcm file> <output file>
```

## Example

```go

func TestProcess(t *testing.T) {
	r := NewRNNoise()

	frameSize := GetFrameSize()
	assert.Equal(t, frameSize, 480)
	o := r.Process([]byte{1, 2, 3, 4, 5, 6})
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0}, o)
}


```
