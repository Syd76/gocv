#ifndef _OPENCV3_CUDA_PHOTO_H_
#define _OPENCV3_CUDA_PHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/photo/cuda.hpp>

extern "C" {
#endif

#include "../core.h"
#include "cuda.h"


//  cf : cv::fastNlMeansDenoisingColored (InputArray src, OutputArray dst, float h=3, float hColor=3, int templateWindowSize=7, int searchWindowSize=21)
void FastNlMeansDenoisingColored( GpuMat m, GpuMat dst, float h_filter, float hcolor , int templateWindowSize, int searchWindowSize );

#ifdef __cplusplus
}
#endif


#endif //_OPENCV3_CUDA_PHOTO_H_
