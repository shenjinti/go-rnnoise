package gornnoise

/*
#cgo CFLAGS: -D__AVX2__ -DRNNOISE_BUILD -I. -Irnnoise -I./rnnoise/src -I./rnnoise/include
#cgo LDFLAGS: -lm

#include "rnnoise/include/rnnoise.h"
#include "rnnoise/src/celt_lpc.c"
#include "rnnoise/src/denoise.c"
#include "rnnoise/src/kiss_fft.c"
#include "rnnoise/src/nnet.c"
#include "rnnoise/src/nnet_default.c"
#include "rnnoise/src/parse_lpcnet_weights.c"
#include "rnnoise/src/pitch.c"
#include "rnnoise/src/rnn.c"
#include "rnnoise/src/rnnoise_tables.c"

*/
import "C"
import (
	"runtime"
	"unsafe"
)

type RNNoise struct {
	inst *C.struct_DenoiseState
}

func GetFrameSize() int {
	return int(C.rnnoise_get_frame_size())
}

func NewRNNoise() *RNNoise {
	r := &RNNoise{
		inst: C.rnnoise_create(nil),
	}
	if r.inst == nil {
		panic("rnnoise_create failed")
	}
	runtime.SetFinalizer(r, destroy)
	return r
}

func destroy(r *RNNoise) {
	if r.inst != nil {
		C.rnnoise_destroy(r.inst)
		r.inst = nil
	}
}

func (r *RNNoise) Process(in []byte) ([]byte, float32) {
	frame := make([]float32, len(in)/2)
	for i := 0; i < len(in)/2; i++ {
		frame[i] = float32(int16(in[i*2]) | int16(in[i*2+1])<<8)
	}
	vadProb := C.rnnoise_process_frame(r.inst, (*C.float)(unsafe.Pointer(&frame[0])), (*C.float)(unsafe.Pointer(&frame[0])))
	out := make([]byte, len(in))
	for i := 0; i < len(frame); i++ {
		out[i*2] = byte(int(frame[i]) & 0xff)
		out[i*2+1] = byte(int(frame[i]) >> 8)
	}
	return out, float32(vadProb)
}
