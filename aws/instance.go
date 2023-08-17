package aws

import "github.com/aws/aws-sdk-go-v2/service/ec2/types"

type InstanceType struct {
	Id     string
	Type   types.InstanceType
	Tags   []types.Tag
	IP     string
	Status types.InstanceState
}
