package main

import (
	"fmt"
	"path/filepath"

	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	//Load font
	path, err := filepath.Abs("/fonts/BaiJamjuree-Regular.ttf")
	if err != nil {
		panic(err)
	}
	pdfapi.LoadConfiguration()
	if err := pdfapi.InstallFonts([]string{path}); err != nil {
		panic(err)
	}
	err = AddWatermark()

	if err != nil {
		panic(err)
	}

}

func AddWatermark() (err error) {
	pageSize, err := pdfapi.PageCountFile("test-1.pdf")
	if err != nil {
		return err
	}

	for i := 1; i <= pageSize; i++ {
		wm, err := pdfapi.TextWatermark("เท่านั้น", "fontname:BaiJamjuree-Regular, position:r, rotation: 0.9, opacity: 0.4", true, true, pdfcpu.POINTS)
		if err != nil {
			panic(err)
		}

		pdfapi.AddWatermarksFile("test-1.pdf", "", []string{fmt.Sprint(i)}, wm, nil)
	}
	return nil
}
