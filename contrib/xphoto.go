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


// Bm3dSteps is the type for the various BM3D algorithm steps
type Bm3dSteps int

const (
  Bm3dAlgoStepAll Bm3dSteps = 0
  Bm3dAlgoSte1 Bm3dSteps = 1
  Bm3dAlgoSte2 Bm3dSteps = 2
)

type TransformTypes int

const (
	Bm3dTypeHaar TransformTypes = 0 
)

// ----------------------- ---------------------------------------
// ----------------------- Bm3dDenoising -----------------------
// ----------------------- ---------------------------------------

func ApplyChannelGains(src gocv.Mat, dst *gocv.Mat, gainB float32, gainG float32, gainR float32) {
	C.Xphoto_ApplyChannelGains( C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.float(gainB), C.float(gainG), C.float(gainR))
	return
}

func Bm3dDenoisingStep( src gocv.Mat, dststep1 *gocv.Mat, dststep2 *gocv.Mat ) {
	C.Xphoto_Bm3dDenoising_Step( C.Mat(src.Ptr()), C.Mat(dststep1.Ptr()), C.Mat(dststep2.Ptr()) ) 
	return
}

func Bm3dDenoisingStepWithParams( src gocv.Mat, dststep1 *gocv.Mat, dststep2 *gocv.Mat,
							h float32, templateWindowSize int,
							searchWindowSize int, blockMatchingStep1 int,
							blockMatchingStep2 int, groupSize int,
							slidingStep int, beta float32,
							normType gocv.NormType, step Bm3dSteps,
							transformType TransformTypes) {
	C.Xphoto_Bm3dDenoising_Step_WithParams( C.Mat(src.Ptr()), C.Mat(dststep1.Ptr()), C.Mat(dststep2.Ptr()) ,
							C.float(h), C.int(templateWindowSize),
							C.int(searchWindowSize), C.int(blockMatchingStep1) ,
							C.int(blockMatchingStep2) , C.int(groupSize) ,
							C.int(slidingStep) , C.float(beta) ,
							C.int(normType) , C.int(step) ,
							C.int(transformType) ) 
	return
}

// func Bm3dDenoising( src gocv.Mat, dststep1 *gocv.Mat, dststep2 *gocv.Mat 
// 							h float32, templateWindowSize int,
// 							searchWindowSize int,blockMatchingStep1 int,
// 							blockMatchingStep2 int,groupSize int,
// 							slidingStep int, beta float32,
// 							normType, int step int,
// 							transformType int ) {
// 	
// 	C.Xphoto_Bm3dDenoising( C.Mat(src.Ptr()), C.Mat(dststep1.Ptr()), C.Mat(dststep2.Ptr()) ,
// 							C.float(h), C.int(templateWindowSize),
// 							C.int(searchWindowSize) , C.int(blockMatchingStep1) ,
// 							C.int(blockMatchingStep2) , C.int(groupSize) ,
// 							C.int(slidingStep) , C.float(beta) ,
// 							C.int(normType) , C.int(step) ,
// 							C.int(transformType) 
// 					)
// 	return
// }

func Bm3dDenoising( src gocv.Mat, dst *gocv.Mat ) {
	C.Xphoto_Bm3dDenoising (C.Mat(src.Ptr()), C.Mat(dst.Ptr()) )
}


func Bm3dDenoisingWithParams( src gocv.Mat, dst *gocv.Mat,
							h float32, templateWindowSize int,
							searchWindowSize int,blockMatchingStep1 int,
							blockMatchingStep2 int,groupSize int,
							slidingStep int, beta float32,
							normType gocv.NormType, step Bm3dSteps,
							transformType TransformTypes ) {

	C.Xphoto_Bm3dDenoising_WithParams ( C.Mat(src.Ptr()), C.Mat(dst.Ptr()),
							C.float(h), C.int(templateWindowSize),
							C.int(searchWindowSize) , C.int(blockMatchingStep1) ,
							C.int(blockMatchingStep2) , C.int(groupSize) ,
							C.int(slidingStep) , C.float(beta) ,
							C.int(normType) , C.int(step) ,
							C.int(transformType) )
	return
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

