package main

import (
	"testing"
	"log"
)

func TestGetSizeSmall(t *testing.T) {
	log.Print(getPageSize("http://www.google.com/robots.txt"))
	//n, _ := getSizeNoScan("http://www.google.com/robots.txt")
	//log.Print(n)
	log.Print(getPageSize("http://lib.ru/INPROZ/FLOBER/salammbo.txt"))
	//n, _ = getSizeNoScan("http://lib.ru/INPROZ/FLOBER/salammbo.txt")
	//log.Print(n)
}

func TestGetSizeBig(t *testing.T) {
	log.Print(getPageSize("https://lwn.net/images/pdf/LDD3/ldd3_pdf.tar.bz2"))
	//n, _ := getSizeNoScan("https://lwn.net/images/pdf/LDD3/ldd3_pdf.tar.bz2")
	//log.Print(n)
}

func TestGetCorrectSize(t *testing.T) {
	expected := int64(8418)
	//sizeTrusty, _ := getSizeNoScan("http://resources.jetbrains.com/assets/media/open-graph/jetbrains_250x250.png")
	size := getPageSize("http://resources.jetbrains.com/assets/media/open-graph/jetbrains_250x250.png")
	log.Print(size)
	if size != expected {
		t.Fatalf("getPageSize() returned %v indtead of %v", size, expected)
	}
}