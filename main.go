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
	dicFilePath := flag.String("dic", "", "Path for your custom dictionary file")
	maxBuffer := flag.Int("buffer", 200, "Buffer's size in MB")
	cleanFilePath := flag.String("clean", "", "Path of the file to remove duplicates")
	flag.Parse()
	if *filePath != "" {
		rules := NewRules()
		workingBanner()
		if *dicFilePath != "" {
			customRules := readDictionary(*dicFilePath)
			rules.AddRuleMap(customRules)
		}
		Extractor(rules, filePath, *extractThis, *outputPath, *maxBuffer)
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
