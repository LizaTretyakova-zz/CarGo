package main

import (
	"testing"
	"log"
)

func TestGetSizeSmall(t *testing.T) {
	log.Print(getPageSize("http://www.google.com/robots.txt"))
	log.Print(getPageSize("http://lib.ru/INPROZ/FLOBER/salammbo.txt"))
}

func TestGetSizeBig(t *testing.T) {
	log.Print(getPageSize("https://lwn.net/images/pdf/LDD3/ldd3_pdf.tar.bz2"))
}

func TestGetCorrectSize(t *testing.T) {
	expected := int64(8418)
	size := getPageSize("http://resources.jetbrains.com/assets/media/open-graph/jetbrains_250x250.png")
	log.Print(size)
	if size != expected {
		t.Fatalf("getPageSize() returned %v indtead of %v", size, expected)
	}
}

func TestGetWordsCount(t *testing.T) {
	log.Print(getWordsCount("http://lib.ru/POEZIQ/GETE/tsar.txt"))
}

func TestIncorrectUrl(t *testing.T) {
	getWordsCount("fdkghkjfdh")
	getPageSize("http://fkhgkfjh")
}