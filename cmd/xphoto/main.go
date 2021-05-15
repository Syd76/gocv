package main

import (
	"flag"
	"fmt"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"os"
)

func inpaint() {

	jpgImageFsrFast := gocv.IMRead("./images/space_shuttle.jpg", gocv.IMReadColor)

	if jpgImageFsrFast.Empty() {
		fmt.Printf("Invalid read of Source Mat in TestInpaint test\n")
		return
	}

	srcFsrFast := gocv.NewMat()
	defer srcFsrFast.Close()
	//         testImage.ConvertTo(&srcFsrFast,gocv.MatTypeCV8UC3)
	sizeImage := jpgImageFsrFast.Size()
	jpgImageFsrFast.ConvertTo(&srcFsrFast, gocv.MatTypeCV8UC3)

	maskFsrFast := gocv.NewMatWithSizes(sizeImage, gocv.MatTypeCV8UC1)
	defer maskFsrFast.Close()

	dstShitMap := gocv.NewMat()
	defer dstShitMap.Close()
	contrib.Inpaint(srcFsrFast, maskFsrFast, &dstShitMap, contrib.FsrFast)

	dstFsrFast := gocv.NewMat()
	defer dstFsrFast.Close()
	contrib.Inpaint(srcFsrFast, maskFsrFast, &dstFsrFast, contrib.FsrFast)

	dstFsrBest := gocv.NewMat()
	defer dstFsrBest.Close()
	contrib.Inpaint(srcFsrFast, maskFsrFast, &dstFsrBest, contrib.FsrFast)

	if dstShitMap.Empty() || dstShitMap.Rows() != srcFsrFast.Rows() || dstShitMap.Cols() != srcFsrFast.Cols() || dstShitMap.Type() != srcFsrFast.Type() {
		fmt.Printf("Invlalid TestInpaint ShitMap test\n")
		return
	}
	fmt.Printf("ShitMap : MAT %d <> %d : %d\n", dstShitMap.Rows(), srcFsrFast.Rows(), dstShitMap.Type())
	gocv.IMWrite("ShitMap_inpaint.png", dstShitMap)

	if dstFsrFast.Empty() || dstFsrFast.Rows() != srcFsrFast.Rows() || dstFsrFast.Cols() != srcFsrFast.Cols() || dstFsrFast.Type() != srcFsrFast.Type() {
		fmt.Printf("Invlalid TestInpaint FsrFast test\n")
		return
	}
	fmt.Printf("FsrFast : MAT %d <> %d : %d\n", dstFsrFast.Rows(), srcFsrFast.Rows(), dstFsrFast.Type())
	gocv.IMWrite("FsrFast_inpaint.png", dstFsrFast)

	if dstFsrBest.Empty() || dstFsrBest.Rows() != srcFsrFast.Rows() || dstFsrBest.Cols() != srcFsrFast.Cols() || dstFsrBest.Type() != srcFsrFast.Type() {
		fmt.Printf("Invlalid TestInpaint FsrBest test\n")
		return
	}
	fmt.Printf("FsrBest : MAT %d <> %d : %d\n", dstFsrBest.Rows(), srcFsrFast.Rows(), dstFsrBest.Type())
	gocv.IMWrite("FsrBest_inpaint.png", dstFsrBest)

}

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
		inpaint()
		return
	}
	flag.Parse()
	if flag.NArg() < 3 {
		flag.Usage()

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
	alignwtb.Process(inputImages, &inputImages)
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
