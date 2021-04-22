package contrib

/*
#include <stdlib.h>
#include "xphoto.h"
*/
import "C"
import (
	"unsafe"
	"gocv.io/x/gocv"
)

// GrayworldWB is a wrapper around the cv::xphoto::GrayworldWB.
type GrayworldWB struct {
	// C.GrayworldWB
	p unsafe.Pointer
}

// LearningBasedWB is a wrapper around the cv::xphoto::LearningBasedWB.
type LearningBasedWB struct {
	// C.GrayworldWB
	p unsafe.Pointer
}


// ----------------------- ---------------------------------------
// ----------------------- GrayworldWB -----------------------
// ----------------------- ---------------------------------------


// NewGrayworldWBWithParams returns a new Gray-world white balance algorithm.
// of type GrayworldWB with customized parameters. GrayworldWB algorithm scales the values
// of pixels based on a gray-world assumption which states that the average of all
// channels should result in a gray image.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
//
func NewGrayworldWB() GrayworldWB {
	return GrayworldWB{p: unsafe.Pointer(C.GrayworldWB_Create())}
}

// SetSaturationThreshold set a Maximum saturation for a pixel to be included in the gray-world assumption.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#ac6e17766e394adc15588b8522202cc71
//
func (b *GrayworldWB) SetSaturationThreshold(saturationThreshold float32) {
	C.GrayworldWB_SetSaturationThreshold((C.GrayworldWB)(b.p), C.float(saturationThreshold))
	return
}

// GetSaturationThreshold return the Maximum saturation for a pixel to be included in the gray-world assumption.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#ac6e17766e394adc15588b8522202cc71
//
func (b *GrayworldWB) GetSaturationThreshold() float32 {
	return float32(C.GrayworldWB_GetSaturationThreshold((C.GrayworldWB)(b.p)))
}

// Close GrayworldWB.
func (b *GrayworldWB) Close() error {
	C.GrayworldWB_Close((C.GrayworldWB)(b.p))
	b.p = nil
	return nil
}

// BalanceWhite computes a Gray-world white balance using the current GrayworldWB.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#details
//
func (b *GrayworldWB) BalanceWhite(src gocv.Mat, dst *gocv.Mat) {
	C.GrayworldWB_BalanceWhite((C.GrayworldWB)(b.p), C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
	return
}



// ----------------------- ---------------------------------------
// ----------------------- LearningBasedWB -----------------------
// ----------------------- ---------------------------------------

func NewLearningBasedWB( ) LearningBasedWB {
	return LearningBasedWB{p: unsafe.Pointer(C.LearningBasedWB_Create())}
}

func NewLearningBasedWBWithParams(pathmodel string ) LearningBasedWB {
	
	cpath := C.CString(pathmodel)
	defer C.free(unsafe.Pointer(cpath))
	
	return LearningBasedWB{p: unsafe.Pointer(C.LearningBasedWB_CreateWithParams(cpath))}
}

// Close LearningBasedWB.
func (b *LearningBasedWB) Close() error {
	C.LearningBasedWB_Close((C.LearningBasedWB)(b.p))
	b.p = nil
	return nil
}


// func (b *LearningBasedWB) extractSimpleFeatures() float32 {
// 	return float32(C.LearningBasedWB_extractSimpleFeatures((C.LearningBasedWB)(b.p)))
// }
// C.LearningBasedWB_extractSimpleFeatures (LearningBasedWB b, Mat src, Mat dst) {
//     (*b)->extractSimpleFeatures (*src, *dst);
// }

func (b *LearningBasedWB) GetHistBinNum() int {
	return int(C.LearningBasedWB_GetHistBinNum((C.LearningBasedWB)(b.p)))
}

func (b *LearningBasedWB) GetRangeMaxVal() int {
	return int(C.LearningBasedWB_GetRangeMaxVal((C.LearningBasedWB)(b.p)))
}
 
func (b *LearningBasedWB) GetSaturationThreshold() float32 {
	return float32(C.LearningBasedWB_GetSaturationThreshold((C.LearningBasedWB)(b.p)))
}

func (b *LearningBasedWB) SetHistBinNum( val int ) {
	C.LearningBasedWB_SetHistBinNum((C.LearningBasedWB)(b.p), C.int(val))
	return 
}

func (b *LearningBasedWB) SetRangeMaxVal( val int ) {
	C.LearningBasedWB_SetRangeMaxVal((C.LearningBasedWB)(b.p), C.int(val))
	return 
}

func (b *LearningBasedWB) SetSaturationThreshold( val float32 ) {
	C.LearningBasedWB_SetSaturationThreshold( (C.LearningBasedWB)(b.p), C.float(val) )
	return
}

