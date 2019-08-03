package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"wume-composer/internal/pkg/models"
)

var (
	envVarRegexp = regexp.MustCompile(`^\${([^|]*)(?:\|([^|]*))?}$`)
)

func getFilename() (string, error) {
	for _, dir := range dirs {
		filename, _ := filepath.Abs(dir + string(filepath.Separator) + "config.json")
		if _, err := os.Stat(filename); err == nil {
			return filename, nil
		}
	}
	return "", models.NotFoundError
}

func load(file string) (File, error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	var config File
	err = jsonParser.Decode(&config)
	return config, err
}

func handleStructure(val reflect.Value) {
	switch val.Kind() {
	// TODO: Add handlers for map, slices, arrays
	// case reflect.Interface:
	// 	handleStructure(val.Elem())
	// case reflect.Map:
	// 	for _, k := range val.MapKeys() {
	// 		handleStructure(val.MapIndex(k))
	// 	}
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for j := 0; j < val.Len(); j++ {
			handleStructure(val.Index(j))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			handleStructure(val.Field(i))
		}
	case reflect.String:
		str := val.String()
		if envVarRegexp.MatchString(str) {
			if val.CanSet() {
				matches := envVarRegexp.FindAllStringSubmatch(str, -1)[0]
				// 0 => str, 1 => before '|', 3 => after '|'
				newStr, hasVal := os.LookupEnv(matches[1])
				if !hasVal && len(matches) == 3 {
					newStr = matches[2]
				}
				val.SetString(newStr)
			} else {
				fmt.Println("Impossible to set environment variable!")
			}
		}
	}
}

func parse(config File) {
	handleStructure(reflect.ValueOf(&config).Elem())
	save(config)
}

func init() {
	file, err := getFilename()
	if err != nil {
		fmt.Println("Config file not found!")
		return
	}
	config, err := load(file)
	if err != nil {
		fmt.Println("Invalid format! Error: " + err.Error())
		return
	}
	parse(config)
}
