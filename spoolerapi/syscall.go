package spoolerapi

import (
	"golang.org/x/sys/windows"
)

var (
	modwinspool = windows.NewLazySystemDLL("winspool.drv")
	modspoolss  = windows.NewLazySystemDLL("spoolss.dll")
)
