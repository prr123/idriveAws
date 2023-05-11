// idriveTest.go
// program to test idrive API
// Author: prr, azulsoftware
// Date: 11 May 2023
// copyright 2023 prr, azulsoftware
//
package main

import (
	"os"
	"fmt"
	"log"

    yaml "github.com/goccy/go-yaml"
)

type IdriveObj struct {
	Url string `yaml:"url"`
	Bucket string `yaml:"bucket"`
	Key string `yaml:"key!`
}

func main() {

	secret, err := getSecret()
	if err != nil {log.Fatalf("getSecret: %v\n", err)}
	log.Printf("secret: %s", secret)

	api, err := getIdriveApi("idriveApi.yaml")
	if err != nil {log.Fatalf("getIdriveApi: %v\n", err)}

	PrintApiObj(api)

	log.Println("success idrive")

}

func getSecret()(sec string, err error) {

	secFilnam :="secret.txt"
	dat, err := os.ReadFile(secFilnam)
	if err != nil {return "", fmt.Errorf("ReadFile %s: %v", secFilnam, err)}

	secret := string(dat)
	return secret, nil
}

func getIdriveApi(yamlFilnam string) (api *IdriveObj, err error) {

	apiObj := &IdriveObj{}
	dat, err := os.ReadFile(yamlFilnam)
	if err != nil {return nil, fmt.Errorf("ReadFile %s: %v", yamlFilnam, err)}

	err = yaml.Unmarshal(dat, apiObj)
	if err != nil {return nil, fmt.Errorf("yaml.Unmarshal: %v", err)}

	return apiObj, nil
}

func PrintApiObj(api *IdriveObj) {

	fmt.Println("************ Idrive Api Object ************")
	if api == nil {return}
	fmt.Printf("Url:    %s\n", api.Url)
	fmt.Printf("Bucket: %s\n", api.Bucket)
	fmt.Printf("Key:    %s\n", api.Key)
	fmt.Println("********* End Idrive Api Object ***********")
	return
}
