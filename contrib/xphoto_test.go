package contrib
//:testing

import (
		"testing"
    "gocv.io/x/gocv"
)


func TestBm3dDenoisingStepWithParams(t *testing.T) {
// src = Input 8-bit or 16-bit 1-channel image.
  src := gocv.NewMatWithSize(200,200,gocv.MatTypeCV8UC1)
  defer src.Close()
  dst1 := gocv.NewMat()
  defer dst1.Close()
  dst2 := gocv.NewMat()
  defer dst2.Close()

  Bm3dDenoisingStepWithParams(src, &dst1, &dst2 ,
								float32(1.0), int(4),
								int(16), int(2500),
								int(400), int(8),
								int(1), float32(2.0),
								gocv.NormL2,
								Bm3dAlgoStepAll,
								Bm3dTypeHaar )



  if src.Empty() || dst1.Rows() != src.Rows() || dst1.Cols() != src.Cols() {
		t.Error("Invlalid TestBm3dDenoisingStepWithParams test")
	}
}

func TestBm3dDenoisingWithParams(t *testing.T) {

  src := gocv.NewMatWithSize(200,200,gocv.MatTypeCV8UC1)
  defer src.Close()
  dst := gocv.NewMat()
  defer dst.Close()

  Bm3dDenoisingWithParams(src, &dst,
								float32(1.0), int(4),
								int(16), int(2500),
								int(400), int(8),
								int(1), float32(2.0),
								gocv.NormL2,
								Bm3dAlgoStepAll,
								Bm3dTypeHaar )
  
  if src.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() || dst.Type() != src.Type() {
		t.Error("Invlalid BalanceWhite test")
	}
}

func TestSetSaturationThreshold(t *testing.T) {

  grayworldwb := NewGrayworldWB()
  var saturation float32 = 0.7
  grayworldwb.SetSaturationThreshold(saturation)

  if grayworldwb.GetSaturationThreshold() < 0 {
    t.Error( " Invlalid SetSaturationThreshold test")
  }
}

func TestBalanceWhite(t *testing.T) {
  grayworldwb := NewGrayworldWB()
  defer grayworldwb.Close()

  src := gocv.NewMatWithSize(200,200,gocv.MatTypeCV8UC3)
  defer src.Close()
  dst := gocv.NewMat()
  defer dst.Close()

  grayworldwb.BalanceWhite(src, &dst)
  if src.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid BalanceWhite test")
	}
}


func TestNewLearningBasedWB(t *testing.T) {

  learningbasedwb := NewLearningBasedWB()
  defer learningbasedwb.Close()

  valueset := 2
  learningbasedwb.SetHistBinNum(valueset)
  learningbasedwb.SetRangeMaxVal(valueset)
  learningbasedwb.SetSaturationThreshold(float32(valueset))

  valuehistbinNum := learningbasedwb.GetHistBinNum()
  valuerangemaxval := learningbasedwb.GetRangeMaxVal()
  valuesaturation := learningbasedwb.GetSaturationThreshold()

  if valueset != valuehistbinNum {
		t.Error("Invalid TestNewLearningBasedWB : GetHistBinNum test")
  }

  if valueset != valuerangemaxval {
		t.Error("Invalid TestNewLearningBasedWB : GetRangeMaxVal test")
  }

  if valuesaturation < 0 {
		t.Error("Invalid TestNewLearningBasedWB : GetSaturationThreshold test")
  }
}
