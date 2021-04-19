
package main

import (
	"flag"
	"fmt"
	"os"
	"gocv.io/x/gocv"
  "gocv.io/x/gocv/contrib"
)

// What it does:
//
// This example uses the MergeMerten class to merge LDR image.
// then save it to an HDR image file on disk.
//
// How to run:
//
// 		go run ./cmd/xphoto/main.go file1.png file2.png file3.png fileresult.png
//
// +build example

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\thdrimage [image file 1] [image file 2] [image file 3] hdr_image.png")
		return
	}
	flag.Parse()
	if flag.NArg() < 3 {
		flag.Usage()
		return
	}

	// read images
	inputs := flag.Args()
	inputImages := make([]gocv.Mat, 3)

	for i := 0; i < 3; i++ {
		img := gocv.IMRead(inputs[i], gocv.IMReadColor)
		if img.Empty() {
			fmt.Printf("cannot read image %s\n", inputs[i])
			return
		}
		defer img.Close()
		inputImages[i] = img
	}

	saveFile := inputs[3]
	ouputHDRImage := gocv.NewMat()
	ouputWBImage := gocv.NewMat()

	defer ouputHDRImage.Close()
	defer ouputWBImage.Close()

	alignwtb := gocv.NewAlignMTBWithParams(3, 20, false)
	alignwtb.Process(inputImages, inputImages)
	defer alignwtb.Close()

	mertens := gocv.NewMergeMertens()
	mertens.Process(inputImages, &ouputHDRImage)
	defer mertens.Close()

	grayworldwb := contrib.NewGrayworldWB()
	grayworldwb.SetSaturationThreshold(0.7)
	grayworldwb.BalanceWhite(ouputHDRImage, &ouputWBImage)
	defer grayworldwb.Close()

	gocv.IMWrite(saveFile, ouputWBImage)

}
