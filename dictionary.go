package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readDictionary(filepath string) map[string]string {

	customDic := make(map[string]string)

	file, err := os.Open(filepath)
	if err != nil {
		log.Panic("No se pudo abrir el archivo ", filepath)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_aux := strings.Split(line, "=")
		if len(line_aux) == 2 {
			ruleName := strings.TrimSpace(line_aux[0])
			customDic[ruleName] = strings.TrimSpace(line_aux[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return customDic
}
