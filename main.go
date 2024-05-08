package main

import (
	"flag"
	"fmt"
)

func main() {

	AppBanner()

	filePath := flag.String("file", "", "Path of the input file")
	extractThis := flag.String("ex", "", "What do you want extract? Email, domain or both?")
	outputPath := flag.String("o", "", "Path of the output files")
	maxBuffer := flag.Int("buffer", 200, "Buffer's size in MB")
	cleanFilePath := flag.String("clean", "", "Path of the file to remove duplicates")
	flag.Parse()
	if *filePath != "" {
		workingBanner()
		if *extractThis == "x" {
			Extractor(filePath, "email", *outputPath, *maxBuffer)
			Extractor(filePath, "domain", *outputPath, *maxBuffer)
		} else {
			Extractor(filePath, *extractThis, *outputPath, *maxBuffer)
		}
	} else if *cleanFilePath != "" {
		workingBanner()
		RemoveDuplicates(*cleanFilePath)
	} else {
		return
	}

}

func workingBanner() {
	fmt.Println("::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("		WORKING...")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::")
}
