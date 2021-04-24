#include "xphoto.h"

void Xphoto_ApplyChannelGains( Mat src, Mat dst , float gainB, float gainG, float gainR) {
	cv::xphoto::applyChannelGains (*src, *dst, gainB, gainG, gainR);
}

void Xphoto_Bm3dDenoising_Step( Mat src, Mat dststep1, Mat dststep2 ) {
	
		cv::xphoto::bm3dDenoising (
								*src, *dststep1, *dststep2, 
								1, 4,
								16, 2500,
								400, 8,
								1, 2.0f,
								cv::NORM_L2, cv::xphoto::BM3D_STEPALL,
								cv::xphoto::HAAR
							);
}

 
void Xphoto_Bm3dDenoising_Step_WithParams( 
							Mat src, Mat dststep1, Mat dststep2,
							float h, int templateWindowSize,
							int searchWindowSize, int blockMatchingStep1,
							int blockMatchingStep2, int groupSize,
							int slidingStep, float beta,
							int normType, int step,
							int transformType
						 ) {
	
		cv::xphoto::bm3dDenoising (
								*src, *dststep1, *dststep2, 
								h, templateWindowSize,
								searchWindowSize, blockMatchingStep1,
								blockMatchingStep2, groupSize,
								slidingStep, beta,
								normType, step,
								transformType
							);
}
 
void Xphoto_Bm3dDenoising ( Mat src, Mat dst ) {
	 
	cv::xphoto::bm3dDenoising (*src, *dst,
							1, 4,
							16, 2500,
							400, 8,
							1, 2.0f,
							cv::NORM_L2, cv::xphoto::BM3D_STEPALL,
							cv::xphoto::HAAR
						);

}

void Xphoto_Bm3dDenoising_WithParams (
							Mat src, Mat dst , float h, int templateWindowSize,
							int searchWindowSize, int blockMatchingStep1,
							int blockMatchingStep2, int groupSize,
							int slidingStep, float beta,
							int normType, int step,
							int transformType
						  ) {
	
	cv::xphoto::bm3dDenoising (*src, *dst, h, templateWindowSize,
							searchWindowSize, blockMatchingStep1,
							blockMatchingStep2, groupSize,
							slidingStep, beta,
							normType, step,
							transformType
						);

}




// ----------------------- GrayworldWB -----------------------

GrayworldWB GrayworldWB_Create() {
    return new cv::Ptr<cv::xphoto::GrayworldWB>(cv::xphoto::createGrayworldWB());
}

void GrayworldWB_Close(GrayworldWB b) {
    delete b;
}

void GrayworldWB_SetSaturationThreshold(GrayworldWB b, float saturationThreshold) {    
    (*b)->setSaturationThreshold(saturationThreshold);
}

float GrayworldWB_GetSaturationThreshold(GrayworldWB b) {
    return (*b)->getSaturationThreshold();   
}

void GrayworldWB_BalanceWhite(GrayworldWB b, Mat src, Mat dst) {
    (*b)->balanceWhite(*src, *dst);
}

// ----------------------- LearningBasedWB -----------------------

LearningBasedWB LearningBasedWB_Create() {
    return new cv::Ptr<cv::xphoto::LearningBasedWB>(cv::xphoto::createLearningBasedWB());
}

LearningBasedWB LearningBasedWB_CreateWithParams( const char* pathmodel ) {

    cv::String path( pathmodel);
    return new cv::Ptr<cv::xphoto::LearningBasedWB>(cv::xphoto::createLearningBasedWB( path ));
}

void LearningBasedWB_Close(LearningBasedWB b) {
    delete b;
}

void LearningBasedWB_ExtractSimpleFeatures (LearningBasedWB b, Mat src, Mat dst) {
    (*b)->extractSimpleFeatures (*src, *dst);
}

int LearningBasedWB_GetHistBinNum ( LearningBasedWB b )  {
    return (*b)->getHistBinNum ();
}
 
int LearningBasedWB_GetRangeMaxVal ( LearningBasedWB b )  {
    return (*b)->getRangeMaxVal ();
}
 
float LearningBasedWB_GetSaturationThreshold ( LearningBasedWB b )  {
    return (*b)->getSaturationThreshold ();
}

void LearningBasedWB_SetHistBinNum (LearningBasedWB b , int val)  {
    (*b)->setHistBinNum (val);
}
 
void LearningBasedWB_SetRangeMaxVal (LearningBasedWB b , int val) {
    (*b)->setRangeMaxVal (val);
}

void LearningBasedWB_SetSaturationThreshold ( LearningBasedWB b , float val){
    (*b)->setSaturationThreshold (val);
}

// ----------------------- LearningBasedWB -----------------------






