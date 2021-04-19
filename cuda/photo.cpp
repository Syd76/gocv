#include "photo.h"

void FastNlMeansDenoisingColored( GpuMat src, GpuMat dst, float h_filter, float hcolor , int templateWindowSize, int searchWindowSize ) {
  cv::cuda::fastNlMeansDenoisingColored( *src , *dst, h_filter, hcolor, templateWindowSize, searchWindowSize);
}
