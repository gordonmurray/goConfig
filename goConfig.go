package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"strings"
)

func main() {

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	resp, err := svc.DescribeInstances(nil)

	if err != nil {
		panic(err)
	}

	for idx := range resp.Reservations {

		for _, inst := range resp.Reservations[idx].Instances {

			for _, tag := range inst.Tags {
				if (*tag.Key == "Name") {
					fmt.Println("Host", strings.Replace(*tag.Value, " ", "", -1))
				}
			}

			fmt.Println("\t# Type", *inst.InstanceType)
			fmt.Println("\tHostname", *inst.PublicDnsName)
			fmt.Println("\tUser ubuntu")
			s := []string{"\tIdentityFile /home/gordon/.ssh/", *inst.KeyName, ".pem"}
			fmt.Println(strings.Join(s, ""))
			fmt.Println(" ")

		}
	}
}
