// WWG Server
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	hostname = "localhost"
	port     = 8080
)

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/buckets", BucketSearchHandler)

	addr := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!!")
	log.Println("The user called to say hello.")
}

func BucketSearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("The user is asking for bucket contents.")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("sydneyclimbing.com"),
	})

	if err != nil {
		fmt.Fprintf(w, "error trying to listobjectsv2: %s", err)
		return
	}

	for _, object := range output.Contents {
		fmt.Fprintf(w, "key=%s size=%d\n", aws.ToString(object.Key), object.Size)
	}

	log.Printf("And we reported %d items.\n", len(output.Contents))
}
