package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/joho/godotenv"
)

func GetInstanceList() ([]InstanceType, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(os.Getenv("AWS_PROFILE")))

	ec2Client := ec2.NewFromConfig(cfg)

	result, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	var instances []InstanceType

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			instances = append(instances, InstanceType{
				Id:     *i.InstanceId,
				Tags:   i.Tags,
				IP:     *i.PrivateIpAddress,
				Status: *i.State,
				Type:   i.InstanceType,
			})
		}
	}

	return instances, nil
}
