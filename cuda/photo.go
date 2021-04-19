package cuda

/*
#include <stdlib.h>
#include "photo.h"
*/
import "C"


// Denoise source GpuMat into destination GpuMat.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#ga03aa4189fc3e31dafd638d90de335617
//
func (m *GpuMat) FastNlMeansDenoisingColored(dst *GpuMat, h_filter float32, hcolor float32, templateWindowSize int , searchWindowSize int) {
	C.FastNlMeansDenoisingColored(m.p, dst.p, C.float(h_filter), C.float(hcolor), C.int(templateWindowSize), C.int(searchWindowSize) )
	return
}
