#ifndef _OPENCV3_XPHOTO_H_
#define _OPENCV3_XPHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xphoto/white_balance.hpp>
#include <opencv2/xphoto/bm3d_image_denoising.hpp>
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


// ----------------------- bm3d_image_denoising -----------------------

void Xphoto_ApplyChannelGains( Mat src, Mat dst , float gainB, float gainG, float gainR) ;

void Xphoto_Bm3dDenoising_Step( Mat src, Mat dststep1, Mat dststep2 ) ;
void Xphoto_Bm3dDenoising_Step_WithParams( 
							Mat src, Mat dststep1, Mat dststep2,  
							float h, int templateWindowSize,
							int searchWindowSize, int blockMatchingStep1,
							int blockMatchingStep2, int groupSize,
							int slidingStep, float beta,
							int normType, int step,
							int transformType
						 ) ;

void Xphoto_Bm3dDenoising ( Mat src, Mat dst ) ;
void Xphoto_Bm3dDenoising_WithParams (
							Mat src, Mat dst , 
							float h, int templateWindowSize,
							int searchWindowSize, int blockMatchingStep1,
							int blockMatchingStep2, int groupSize,
							int slidingStep, float beta,
							int normType, int step,
							int transformType
						  ) ;


// ----------------------- GrayworldWB -----------------------

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
