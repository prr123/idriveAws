// idriveLib
// library to access idrive
// author: prr azulsoftware
// date: 12 May 2023
// copyright 2023 prr azulsoftware
//
package idriveLib

import (
	"os"
	"fmt"

    yaml "github.com/goccy/go-yaml"
)

type IdriveObj struct {
    Url string `yaml:"url"`
    Bucket string `yaml:"bucket"`
	Region string `yaml:"region"`
    Key string `yaml:"key"`
	Secret string
}


func GetSecret()(sec string, err error) {

    secFilnam :="secret.txt"
    dat, err := os.ReadFile(secFilnam)
    if err != nil {return "", fmt.Errorf("ReadFile %s: %v", secFilnam, err)}
	datLast := len(dat) -1

	// remove leading white spaces
	iFirst :=0
	for i:=0; i< datLast; i++ {
		if dat[i] != ' ' {
			iFirst = i
			break
		}
	}

	// remove all trailing non-alpha characters
	iLast := datLast
	for i:= datLast; i>0; i-- {
		let := dat[i]
    	if (let >= 'a' && let <= 'z') || (let >= 'A' && let <= 'Z') {
			iLast = i
			break
		}
	}

    return string(dat[iFirst:iLast+1]), nil
}

func GetIdriveApi(yamlFilnam string) (api *IdriveObj, err error) {

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
	fmt.Printf("Region: %s\n", api.Region)
    fmt.Printf("Bucket: %s\n", api.Bucket)
    fmt.Printf("Key:    %s\n", api.Key)
    fmt.Println("********* End Idrive Api Object ***********")
    return
}
