package contrib
//:testing

import (
		"testing"
    "gocv.io/x/gocv"
)

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
	
	
	
