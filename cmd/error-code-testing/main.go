package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cblkwell/go-playground/internal/aws/session"

	"fmt"
	"os"
)

// This function is for establishing our session with AWS.
func makeEC2Client(region, profile string) *ec2.EC2 {
        sess := session.MustMakeSession(region, profile)
        ec2Client := ec2.New(sess)
        return ec2Client
}

func testTerminate(instanceID string) error {
	ec2Client := makeEC2Client(os.Getenv("AWS_REGION"), os.Getenv("AWS_PROFILE"))
	terminateInput := &ec2.TerminateInstancesInput{
		DryRun:      aws.Bool(true),
		InstanceIds: []*string{aws.String(instanceID)},
	}
	_, err := ec2Client.TerminateInstances(terminateInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Code())
				fmt.Println(aerr.Error())

			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}
	return err
}

func main() {
	testTerminate(os.Args[1])
}
