#include "xphoto.h"

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






