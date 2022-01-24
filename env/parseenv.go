package env

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once
)

// Get Exports
func Get(config interface{}, envFile ...string) func() interface{} {
	once.Do(func() {
		if err := godotenv.Load(envFile...); err != nil {
			log.Print("[!] No .env file found")
			// return
		}
		formatConfig(config)
	})

	return func() interface{} {
		return config
	}
}

func formatConfig(config interface{}) {
	val := reflect.ValueOf(&config).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get("env")

		valueEnv, _ := os.LookupEnv(tag)

		switch typeField.Type.Kind() {
		case reflect.Int:
			value, _ := strconv.Atoi(valueEnv)
			valueField.Set(reflect.ValueOf(value))
		case reflect.String:
			valueField.SetString(valueEnv)
		case reflect.Bool:
			value := false
			if valueEnv == "true" {
				value = true
			}
			valueField.SetBool(value)
		}
	}
}
