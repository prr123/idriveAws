package main

import (
//	"os"
//	"fmt"
	"log"

	idrive	"api/idrive/idriveLib"


)

func main() {

    api, err := idrive.GetIdriveApi("idriveApi.yaml")
    if err != nil {log.Fatalf("getIdriveApi: %v\n", err)}
    log.Println("success idrive api")

    secret, err := idrive.GetSecret()
    if err != nil {log.Fatalf("getSecret: %v\n", err)}
    log.Printf("secret: |%s|", secret)

	api.Secret = secret

	idrive.PrintApiObj(api)

}
