package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (f *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(f.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("could not read line in file")
	}

	file.Close()
	return lines, nil
}

func (f *FileManager) WriteJSON(data any) error {
	file, err := os.Create(f.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}

	file.Close()
	return nil
}

func New(InputFilePath, OutputFilePath string) (FileManager, error) {
	if InputFilePath == "" || OutputFilePath == "" {
		err := errors.New("missing input or output file path")
		return FileManager{}, err
	}

	return FileManager{
		InputFilePath:  InputFilePath,
		OutputFilePath: OutputFilePath,
	}, nil
}
