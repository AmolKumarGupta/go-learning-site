package core

import (
	"encoding/json"
	"os"
)

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)

	return file, err
}

func Save(fileName string, data any) error {
	path := "storage/" + fileName + ".json"

	_, err := os.Stat(path)

	if err == nil {
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var prevprevData []interface{}
		err = json.Unmarshal(file, &prevprevData)
		if err != nil {
			return err
		}

		prevprevData = append(prevprevData, data)
		encoded, err := json.Marshal(prevprevData)
		if err != nil {
			return err
		}

		err = os.WriteFile(path, encoded, 0766)
		if err != nil {
			return err
		}

		return nil
	}

	file, err := createFile(path)
	if err != nil {
		return err
	}

	jsonData := []any{data}
	encoded, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	file.WriteString(string(encoded))
	return nil
}

func getFile(fileName string, prev any) error {
	path := "storage/" + fileName + ".json"

	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	json.Unmarshal(data, prev)

	return nil
}
