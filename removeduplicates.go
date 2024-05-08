package main

import (
	"bufio"
	"fmt"
	"os"
)

// RemoveDuplicates elimina los elementos duplicados de un archivo
func RemoveDuplicates(filePath string) error {
	// Abrir el archivo de entrada
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un mapa para almacenar elementos únicos
	uniqueElements := make(map[string]bool)

	// Crear un archivo temporal para escribir los elementos únicos
	tempFile, err := os.CreateTemp("", "tempfile")
	if err != nil {
		return fmt.Errorf("error al crear el archivo temporal: %v", err)
	}
	defer tempFile.Close()

	// Crear scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Leer el archivo línea por línea y almacenar elementos únicos en el mapa
	for scanner.Scan() {
		line := scanner.Text()
		uniqueElements[line] = true
	}

	// Verificar errores de escaneo
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error al leer el archivo línea por línea: %v", err)
	}

	// Escribir elementos únicos en el archivo temporal
	for element := range uniqueElements {
		if _, err := tempFile.WriteString(element + "\n"); err != nil {
			return fmt.Errorf("error al escribir en el archivo temporal: %v", err)
		}
	}
	fmt.Printf("Total unique results: %02d \n", len(uniqueElements))

	// Cerrar el archivo original
	if err := file.Close(); err != nil {
		return fmt.Errorf("error al cerrar el archivo original: %v", err)
	}

	// Eliminar el archivo original
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("error al eliminar el archivo original: %v", err)
	}

	// Renombrar el archivo temporal como el archivo original
	if err := os.Rename(tempFile.Name(), filePath); err != nil {
		return fmt.Errorf("error al renombrar el archivo temporal: %v", err)
	}

	return nil
}
