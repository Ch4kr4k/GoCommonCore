package file

import (
	"bufio"
	"os"
)

func ReadFileLineByLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func checkPathType(path string) (string, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "directory", nil
	}
	return "file", nil
}
