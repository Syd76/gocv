package cuda


func TestFastNlMeansDenoisingColored(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in GaussianFilter test")
	}
	defer src.Close()

	cimg := NewGpuMat()
	defer cimg.Close()
	cimg.Upload(src)

  dimg := NewGpuMat()
	defer dimg.Close()

	cimg.FastNlMeansDenoisingColored(dimg, 3,3,7,21)

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty TestFastNlMeansDenoisingColored test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid TestFastNlMeansDenoisingColored test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid TestFastNlMeansDenoisingColored test cols")
	}
}
