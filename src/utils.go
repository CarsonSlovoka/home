package main

import . "github.com/CarsonSlovoka/go-pkg/v2/fmt"

var (
	pOK, pErr, pInfo *ColorPrinter
)

func init() {
	pOK = NewColorPrinter(0, 0, 0, 0, 255, 0)
	pErr = NewColorPrinter(255, 255, 255, 255, 0, 0)
	pInfo = NewColorPrinter(255, 255, 255, 0, 0, 255)
}
