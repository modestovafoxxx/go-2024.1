package main

import "os"

// Пример реализации функции saveToFile
func saveToFile(fileNames []string, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, fileName := range fileNames {
		_, err = file.WriteString(fileName + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Реализовать функцию получения всех файлов из директории
func getFiles(dir string) ([]string, error) {
	return nil, nil
}

// Реализовать функцию фильтрации переданных названий файлов
func filterFiles(files []string, filter string) ([]string, error) {
	return nil, nil
}

func main() {

}
