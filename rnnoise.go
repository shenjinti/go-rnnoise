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

func bytesToF32Frame(b []byte) []float32 {
	f := make([]float32, GetFrameSize())
	for i := 0; i < len(b)/2; i++ {
		f[i] = float32(int16(b[i*2]) | int16(b[i*2+1])<<8)
	}
	return f
}

func f32FrameToBytes(origLen int, f []float32) []byte {
	b := make([]byte, origLen)
	for i := 0; i < origLen/2; i++ {
		b[i*2] = byte(int(f[i]) & 0xff)
		b[i*2+1] = byte(int(f[i]) >> 8)
	}
	return b
}

func (r *RNNoise) Process(in []byte) []byte {
	remain := len(in)
	var out []byte
	for remain > 0 {
		bufLen := len(in)
		if len(in) > GetFrameSize()*2 {
			bufLen = GetFrameSize() * 2
		}
		frame := bytesToF32Frame(in[:bufLen])
		C.rnnoise_process_frame(r.inst, (*C.float)(unsafe.Pointer(&frame[0])), (*C.float)(unsafe.Pointer(&frame[0])))
		in = in[bufLen:]
		remain = len(in)
		out = append(out, f32FrameToBytes(bufLen, frame)...)
	}
	return out
}
