package keytranslation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var (
	kt KeyTranslation
)

type KeyTranslation struct {
	Keys map[string]string `json:"keys"`
}

func (kt *KeyTranslation) Get(key string) string {
	r, ok := kt.Keys[key]
	if !ok {
		return key
	}
	return r
}

func Get(key string) string {
	return kt.Get(key)
}

func LoadKeyTranslations() {
	m, err := parseJSON("./fixtures/translations.json", &KeyTranslation{})
	if err != nil {
		fmt.Printf("error loading key translation : %s", err)
	}
	kt = *(m.(*KeyTranslation))
}

func parseJSON(fixturesPath string, defaultValue interface{}) (interface{}, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(fixturesPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return defaultValue, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var copied interface{}
	if reflect.ValueOf(defaultValue).Kind() == reflect.Ptr {
		copied = reflect.New(reflect.ValueOf(defaultValue).Elem().Type()).Interface()
	} else {
		copied = reflect.New(reflect.ValueOf(defaultValue).Type()).Interface()
	}
	if err := json.Unmarshal([]byte(byteValue), copied); err != nil {
		return defaultValue, err
	}

	return copied, nil
}
