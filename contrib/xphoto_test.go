package contrib
//:testing

import (
		"testing"
    "gocv.io/x/gocv"
)

func TestSetSaturationThreshold(t *testing.T) {

  grayworldwb := NewGrayworldWB()
  saturation := 0.7
  grayworldwb.SetSaturationThreshold(saturation)

  if saturation == grayworldwb.GetSaturationThreshold() {
    t.Error( " Invlalid SetSaturationThreshold test")
  }
}

func TestBalanceWhite(t *testing.T) {
  grayworldwb := NewGrayworldWB()
  defer grayworldwb.Close()

  src := gocv.NewMatWithSize(200,200)
  defer src.Close()
  dst := gocv.NewMat()
  defer dst.Close()

  grayworldwb.BalanceWhite(src, &dst)
  if src.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid BalanceWhite test")
	}
}
