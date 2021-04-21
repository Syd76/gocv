#ifndef _OPENCV3_XPHOTO_H_
#define _OPENCV3_XPHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xphoto/white_balance.hpp>
extern "C" {
#endif

#include "../core.h"
    
#ifdef __cplusplus
typedef cv::Ptr<cv::xphoto::GrayworldWB>* GrayworldWB;
typedef cv::Ptr<cv::xphoto::LearningBasedWB>* LearningBasedWB;
#else
typedef void* GrayworldWB;
typedef void* LearningBasedWB;
#endif

GrayworldWB GrayworldWB_Create();
void GrayworldWB_Close(GrayworldWB b);
void GrayworldWB_SetSaturationThreshold(GrayworldWB b, float saturationThreshold);
float GrayworldWB_GetSaturationThreshold(GrayworldWB b);
void GrayworldWB_BalanceWhite(GrayworldWB b, Mat src, Mat dst); 

// ----------------------- LearningBasedWB -----------------------

LearningBasedWB LearningBasedWB_Create();
LearningBasedWB LearningBasedWB_CreateWithParams( const char *pathmodel );
void LearningBasedWB_Close(LearningBasedWB b);
void LearningBasedWB_ExtractSimpleFeatures (LearningBasedWB b , Mat src, Mat dst);
int LearningBasedWB_GetHistBinNum ( LearningBasedWB b ) ;
int LearningBasedWB_GetRangeMaxVal (LearningBasedWB b ) ;
float LearningBasedWB_GetSaturationThreshold (LearningBasedWB b ) ;
void LearningBasedWB_SetHistBinNum (LearningBasedWB b , int val);
void LearningBasedWB_SetRangeMaxVal ( LearningBasedWB b , int val);
void LearningBasedWB_SetSaturationThreshold (LearningBasedWB b , float val);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XPHOTO_H
