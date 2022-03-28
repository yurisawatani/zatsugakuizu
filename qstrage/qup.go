package qstrage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const BucketName = "learning-yuri"

func GetClient(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return client, nil
}

func WhiteFile(ctx context.Context, s3path string, content io.Reader) error {
	log.Printf("Write s3 file: %s%s", BucketName, s3path)
	client, err := GetClient(ctx)
	if err != nil {
		return err
	}
	obj := &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(s3path),
		Body:   content,
	}
	if _, err := client.PutObject(ctx, obj); err != nil {
		return err
	}
	return nil
}

func readJson(filepath string) (string, error) {
	log.Println("Loading json:", filepath)
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("Failed to open file: %s: %v\n", filepath, err)
	}
	return string(bs), nil
}

func main() {
	filepath := os.Args[1]
	jstr, err := readJson(filepath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx := context.Background()
	WhiteFile(ctx, filepath, strings.NewReader(jstr))
}
