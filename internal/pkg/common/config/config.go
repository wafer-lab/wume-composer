package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
)

var (
	noConfigFile = errors.New("no config file")
	envVarRegexp = regexp.MustCompile(`^\${([^|]*)(?:\|([^|]*))?}$`)
)

func getFilename() (string, error) {
	for _, filename := range paths {
		filename, _ := filepath.Abs(filename)
		if _, err := os.Stat(filename); err == nil {
			return filename, nil
		}
	}
	return "", noConfigFile
}

func load(file string) (config *File, err error) {
	configFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		_ = configFile.Close()
		return nil, err
	}

	err = configFile.Close()
	if err != nil {
		return nil, err
	}

	return config, nil
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
				log.Println("Impossible to set environment variable!")
			}
		}
	}
}

func parse(config *File) {
	handleStructure(reflect.ValueOf(config).Elem())
	save(config)
}

func init() {
	file, err := getFilename()
	if err != nil {
		log.Fatalln("Config file not found!")
	}
	config, err := load(file)
	if err != nil {
		log.Fatalln("Impossible to load config! Error: " + err.Error())
	}
	parse(config)
}
