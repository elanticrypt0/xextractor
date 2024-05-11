package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func initExtractedResults(rules *Rules) map[string][]string {
	var extractedResults = make(map[string][]string)

	for rulesKey := range rules.Rules {
		extractedResults[rulesKey] = []string{}
	}

	return extractedResults
}

func Extractor(rules *Rules, filePath *string, extractThis string, outputPath string, maxBuffer int) {

	extractThis = strings.ToLower(extractThis)

	// chequea que lo que se quiere extraer existe en las reglas

	_, ruleExists := rules.Rules[extractThis]

	if !ruleExists {
		fmt.Println("-ex must be valid rule. For default: [email|domain|x (both)]")
		rules.PrintAvailables()
		return
	}

	fmt.Printf("Rule setted > %q \n", extractThis)

	extractedResults := initExtractedResults(rules)

	// Verificar si se proporcionó la ruta del archivo
	if *filePath == "" {
		fmt.Println("Por favor, proporciona la ruta del archivo usando la bandera -file")
		return
	}

	// Abrir el archivo de entrada
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Crear scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Calcular el tamaño máximo del buffer en bytes (300 MB)
	maxBufferSize := maxBuffer * 1024 * 1024

	// Configurar el tamaño máximo del buffer del scanner
	scanner.Buffer(make([]byte, maxBufferSize), maxBufferSize)

	// select and process the file
	extractedResults = processFile(scanner, extractThis, rules.Rules, extractedResults)

	// Verificar errores de escaneo
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo línea por línea: %v\n", err)
		return
	}

	if extractThis == "x" {
		for ruleKey := range rules.Rules {
			if len(extractedResults[ruleKey]) > 0 {
				createExtractsFile(outputPath, ruleKey, extractedResults[ruleKey])
			}
		}
	} else {
		if len(extractedResults[extractThis]) > 0 {
			createExtractsFile(outputPath, extractThis, extractedResults[extractThis])
		}
	}
}

func processFile(scanner *bufio.Scanner, extractThis string, rules map[string]string, extractedResults map[string][]string) map[string][]string {

	// Leer el archivo línea por línea
	for scanner.Scan() {
		line := scanner.Text()

		if extractThis == "x" {
			for ruleKey := range rules {
				extractedResults = processLine(ruleKey, line, rules, extractedResults)
			}
		} else {
			extractedResults = processLine(extractThis, line, rules, extractedResults)
		}

	}

	return extractedResults
}

func createExtractsFile(outputPath, extractThis string, extractedSlice []string) {

	extractedFilePath := setOutputfileName(outputPath, extractThis)

	// Crear el archivo CSV de todos los emails
	allResultssCSVFile, err := os.Create(extractedFilePath)
	if err != nil {
		fmt.Printf("Error al crear el archivo de resultados: %v\n", err)
		return
	}
	defer allResultssCSVFile.Close()

	// Escribir todos los emails en el archivo CSV
	allResultsWriter := csv.NewWriter(allResultssCSVFile)
	defer allResultsWriter.Flush()

	for _, email := range extractedSlice {
		if err := allResultsWriter.Write([]string{email}); err != nil {
			fmt.Printf("Error al escribir en el archivo CSV: %v\n", err)
			return
		}
	}

	fmt.Printf("All %s saved in %s \n", extractThis, extractedFilePath)
	fmt.Printf("Total results: %02d \n", len(extractedSlice))

}

func setOutputfileName(outputPath, extractThis string) string {
	fileTimestap := time.Now()
	fileTimestap_date := fmt.Sprintf("%d%02d%02d-%02d%02d%02d", fileTimestap.Year(), fileTimestap.Month(), fileTimestap.Day(), fileTimestap.Hour(), fileTimestap.Minute(), fileTimestap.Second())

	outputFileName := fmt.Sprintf("%s_%s.csv", extractThis, fileTimestap_date)

	if outputFileName != "" {
		outputFileName = outputPath + "/" + outputFileName
	}

	return outputFileName
}

func processLine(extractThis, line string, rules map[string]string, extractedResults map[string][]string) map[string][]string {

	// fmt.Println("Entra acá process line")

	var results = []string{}
	var selectedRegex *regexp.Regexp

	// if you are looking for domain checks if the line is not an email.
	// if is it then parseit as an email and extracts the domain
	if extractThis == "domain" {
		rule2Work := extractThis
		if strings.Contains(line, "@") {
			rule2Work = "email"
		}

		selectedRegex = regexp.MustCompile(rules[rule2Work])
		results = selectedRegex.FindAllString(line, -1)

	} else {
		selectedRegex = regexp.MustCompile(rules[extractThis])
		results = selectedRegex.FindAllString(line, -1)
	}

	// fmt.Printf("%q", selectedRegex)
	// fmt.Printf("%q \n", extractThis)

	// Parse every result on the line.
	for _, item := range results {
		// check if the line isnot an email
		if extractThis == "domain" {
			// Extraer el dominio
			parts := strings.Split(item, "@")

			if len(parts) == 2 {
				isDomainRegex := regexp.MustCompile(rules["domain"])
				if isDomainRegex.MatchString(parts[1]) {
					extractedResults[extractThis] = append(extractedResults[extractThis], strings.ToLower(parts[1]))
				}
			} else {
				extractedResults[extractThis] = append(extractedResults[extractThis], strings.ToLower(item))
			}
		} else {
			extractedResults[extractThis] = append(extractedResults[extractThis], strings.ToLower(item))
		}
	}

	return extractedResults
}
