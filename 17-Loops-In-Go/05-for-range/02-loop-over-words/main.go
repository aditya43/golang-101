package main

import (
	"fmt"
	"strings"
)

func main() {
	// this will split the following string by spaces
	// and then it will put it inside words variable
	// as a slice of strings
	words := strings.Fields("🍺 AWS Serverless - Serverless Framework, AWS SAM, Lambda, Step Functions, API Gateway, RDS, DynamoDB, ElephantSQL, CI/CD, CloudFormation, S3 Notifications, SNS, SQS, Cognito, Lambda Crons, CORS, Apache VTL (Velocity Template Language), Swagger, KMS, VPCs, DLQs, CloudWatch, CodeCommit, CodeBuild, CodePipeline, CloudFront, OICD, Kinesis, MQTT, IoT and lot more..")

	// --------------------------------
	// #1st way:
	// --------------------------------
	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-5d: %q\n", j+1, words[j])
	// }

	// --------------------------------
	// #2nd way (best):
	// --------------------------------
	for i, v := range words {
		fmt.Printf("#%-5d: %q\n", i+1, v)
	}

	// --------------------------------
	// #3rd way (reuse mechanism):
	// --------------------------------
	// var (
	// 	i int
	// 	v string
	// )

	// for i, v = range words {
	// 	fmt.Printf("#%-5d: %q\n", i+1, v)
	// }

	// fmt.Printf("Last value of i and v %d %q\n", i, v)
}
