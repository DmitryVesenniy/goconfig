package ini

import (
	"log"
	"reflect"
	"strconv"
	"sync"
)

var (
	once sync.Once
)

// Get Exports
func Get(config interface{}, iniFiles ...string) func() interface{} {
	once.Do(func() {
		fields, err := Load(iniFiles...)
		if err != nil {
			log.Print("[!] No .ini file found")
		}
		formatConfig(config, fields)
	})

	return func() interface{} {
		return config
	}
}

func formatConfig(config interface{}, fields map[string]string) {
	val := reflect.ValueOf(config).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("ini")

		valueIni := ""
		for fieldName, fieldValue := range fields {
			if fieldName == tag {
				valueIni = fieldValue
			}
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			value, _ := strconv.Atoi(valueIni)
			valueField.Set(reflect.ValueOf(value))
		case reflect.String:
			valueField.SetString(valueIni)
		case reflect.Bool:
			value := false
			if valueIni == "true" {
				value = true
			}
			valueField.SetBool(value)
		}
	}
}
