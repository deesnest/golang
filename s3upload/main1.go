package main

import (
    "bytes"
    "log"
    "net/http"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)
  const (
    S3_REGION = "ap-south-1"
    S3_BUCKET = "golangbucket"
)
func main() {

    // create an AWS session which can be
    // reused if we're uploading many files
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(S3_REGION),
        Credentials: credentials.NewStaticCredentials(
            "AKIAWXL533ZHLA7764DX",
            "4+RzgKRjc/eImEDF5DHo4GpWYZXsbtpUMID50QkJ",
            ""),
    })

    if err != nil {
        log.Fatal(err)
    }

    err = uploadFileToS3(s, "result1.csv")

    if err != nil {
        log.Fatal(err)
    }
}
func uploadFileToS3(s *session.Session, fileName string) error {

    // open the file for use
    file, err := os.Open(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    // get the file size and read
    // the file content into a buffer
    fileInfo, _ := file.Stat()
    var size = fileInfo.Size()
    buffer := make([]byte, size)
    file.Read(buffer)

    // config settings: this is where you choose the bucket,
    // filename, content-type and storage class of the file
    // you're uploading
    _,s3err := s3.New(s).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(S3_BUCKET),
        Key:                  aws.String(fileName),
        ACL:                  aws.String("private"),
        Body:                 bytes.NewReader(buffer),
        ContentLength:        aws.Int64(size),
        ContentType:          aws.String(http.DetectContentType(buffer)),
        ContentDisposition:   aws.String("attachment"),
        ServerSideEncryption: aws.String("AES256"),
        StorageClass:         aws.String("INTELLIGENT_TIERING"),
    })
    return s3err
}
