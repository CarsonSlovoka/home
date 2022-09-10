package utils

import . "github.com/CarsonSlovoka/go-pkg/v2/fmt"

var (
	POk   *ColorPrinter
	PErr  *ColorPrinter
	PInfo *ColorPrinter
)

func init() {
	POk = NewColorPrinter(0, 0, 0, 0, 255, 0)
	PErr = NewColorPrinter(255, 255, 255, 255, 0, 0)
	PInfo = NewColorPrinter(255, 255, 255, 0, 0, 255)
}
