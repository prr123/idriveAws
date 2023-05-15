package main

import (
//	"os"
	"fmt"
	"log"
	"context"
	"time"

	idrive	"api/idrive/idriveLib"

//	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

)

func main() {

    api, err := idrive.GetIdriveApi("idriveApi.yaml")
    if err != nil {log.Fatalf("getIdriveApi: %v\n", err)}
    log.Println("success idrive api")

    secret, err := idrive.GetSecret()
    if err != nil {log.Fatalf("getSecret: %v\n", err)}
    log.Printf("secret: %s", secret)

	api.Secret = secret

	idrive.PrintApiObj(api)

//	os.Exit(1)
	accessKey := api.Key

	cred:= credentials.NewStaticCredentialsProvider(accessKey, secret, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
 		config.WithCredentialsProvider(cred),
    	config.WithRegion("us-east-1"),
	)
	if err != nil {log.Fatalf("config: %v", err)}

	url := "https://" + api.Url
	// Create a new S3 service client
	client := s3.NewFromConfig(cfg, s3.WithEndpointResolver(
      // Applying of the Zerops Object Storage API URL endpoint.
      s3.EndpointResolverFromURL(url),
    ),
    func(opts *s3.Options) {
      // Zerops supports currently only S3 path-style addressing model.
      // The virtual-hosted style model will be supported in near future.
      opts.UsePathStyle = true
    },)

	// Set the parameters based on the CLI flag inputs.
//	params := &s3.ListObjectsV2Input{Bucket: &api.Bucket,}
/*
	if len(objectPrefix) != 0 {
		params.Prefix = &objectPrefix
	}
	if len(objectDelimiter) != 0 {
		params.Delimiter = &objectDelimiter
	}
*/

	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {log.Fatalf("ListBuckets: %v", err)}

	fmt.Printf("result: %v\n", result)

	PrintListBucketOutput(*result)
}

func PrintListBucketOutput(res s3.ListBucketsOutput) {

	fmt.Println("*************************************")
	fmt.Printf("Buckets: %d\n", len(res.Buckets))
	fmt.Printf("Owner: \n")
	owner := *(res.Owner)
	fmt.Printf("  DisplayName: %s\n", (*owner.DisplayName))

	fmt.Printf("Metadata: %v\n", res.ResultMetadata)

	for i:=0; i< len(res.Buckets); i++ {
		bucket := res.Buckets[i]
		fmt.Printf("***** Bucket: %d *******\n", i+1)
		tim := (*bucket.CreationDate).Format(time.RFC1123)
		fmt.Printf("  Name: %-20s CreationDate: %s\n", (*bucket.Name), tim)
	}

}
