// >>>>>>> DO NOT EDIT THIS FILE <<<<<<<<<<
// This file is autogenerated via `aws-operator-codegen process`
// If you'd like the change anything about this file make edits to the .templ
// file in the pkg/codegen/assets directory.

package dynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsV1alpha1 "github.com/christopherhein/aws-operator/pkg/apis/operator.aws/v1alpha1"
	"github.com/christopherhein/aws-operator/pkg/config"
	"github.com/christopherhein/aws-operator/pkg/helpers"
)

// New generates a new object
func New(config *config.Config, dynamodb *awsV1alpha1.DynamoDB, topicARN string) *Cloudformation {
	return &Cloudformation{
		DynamoDB: dynamodb,
		config:					config,
    topicARN:       topicARN,
	}
}

// Cloudformation defines the dynamodb cfts
type Cloudformation struct {
	config         *config.Config
	DynamoDB *awsV1alpha1.DynamoDB
  topicARN       string
}

// StackName returns the name of the stack based on the aws-operator-config
func (s *Cloudformation) StackName() string {
	return helpers.StackName(s.config.ClusterName, "dynamodb", s.DynamoDB.Name, s.DynamoDB.Namespace)
}

// GetOutputs return the stack outputs from the DescribeStacks call
func (s *Cloudformation) GetOutputs() (map[string]string, error) {
	outputs := map[string]string{}
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	stackInputs := cloudformation.DescribeStacksInput{
		StackName:   aws.String(s.StackName()),
	}

	output, err := svc.DescribeStacks(&stackInputs)
	if err != nil {
		return nil, err
	}
	// Not sure if this is even possible
	if len(output.Stacks) != 1 {
		return nil, errors.New("no stacks returned with that stack name")
	}

	for _, out := range output.Stacks[0].Outputs {
		outputs[*out.OutputKey] = *out.OutputValue
	}

	return outputs, err
}

// CreateStack will create the stack with the supplied params
func (s *Cloudformation) CreateStack() (output *cloudformation.CreateStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	cftemplate := helpers.GetCloudFormationTemplate(s.config, "dynamodb", s.DynamoDB.Spec.CloudFormationTemplateName, s.DynamoDB.Spec.CloudFormationTemplateNamespace)

	stackInputs := cloudformation.CreateStackInput{
		StackName:   aws.String(s.StackName()),
		TemplateURL: aws.String(cftemplate),
		NotificationARNs: []*string{
			aws.String(s.topicARN),
		},
	}

	resourceName := helpers.CreateParam("ResourceName", s.DynamoDB.Name)
	resourceVersion := helpers.CreateParam("ResourceVersion", s.DynamoDB.ResourceVersion)
	namespace := helpers.CreateParam("Namespace", s.DynamoDB.Namespace)
	clusterName := helpers.CreateParam("ClusterName", s.config.ClusterName)
	tableName := helpers.CreateParam("TableName", helpers.Stringify(s.DynamoDB.Name))
	rangeAttributeNameTemp := "{{.Obj.Spec.RangeAttribute.Name}}"
	rangeAttributeNameValue, err := helpers.Templatize(rangeAttributeNameTemp, helpers.Data{Obj: s.DynamoDB, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
  rangeAttributeName := helpers.CreateParam("RangeAttributeName", helpers.Stringify(rangeAttributeNameValue))
	rangeAttributeTypeTemp := "{{.Obj.Spec.RangeAttribute.Type}}"
	rangeAttributeTypeValue, err := helpers.Templatize(rangeAttributeTypeTemp, helpers.Data{Obj: s.DynamoDB, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
  rangeAttributeType := helpers.CreateParam("RangeAttributeType", helpers.Stringify(rangeAttributeTypeValue))
	readCapacityUnitsTemp :=	"{{.Obj.Spec.ReadCapacityUnits}}"
	readCapacityUnitsValue, err := helpers.Templatize(readCapacityUnitsTemp, helpers.Data{Obj: s.DynamoDB, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	readCapacityUnits := helpers.CreateParam("ReadCapacityUnits", helpers.Stringify(readCapacityUnitsValue))
	writeCapacityUnitsTemp :=	"{{.Obj.Spec.WriteCapacityUnits}}"
	writeCapacityUnitsValue, err := helpers.Templatize(writeCapacityUnitsTemp, helpers.Data{Obj: s.DynamoDB, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	writeCapacityUnits := helpers.CreateParam("WriteCapacityUnits", helpers.Stringify(writeCapacityUnitsValue))
	hashAttributeNameTemp := "{{.Obj.Spec.HashAttribute.Name}}"
	hashAttributeNameValue, err := helpers.Templatize(hashAttributeNameTemp, helpers.Data{Obj: s.DynamoDB, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
  hashAttributeName := helpers.CreateParam("HashAttributeName", helpers.Stringify(hashAttributeNameValue))
	hashAttributeTypeTemp := "{{.Obj.Spec.HashAttribute.Type}}"
	hashAttributeTypeValue, err := helpers.Templatize(hashAttributeTypeTemp, helpers.Data{Obj: s.DynamoDB, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
  hashAttributeType := helpers.CreateParam("HashAttributeType", helpers.Stringify(hashAttributeTypeValue))

	parameters := []*cloudformation.Parameter{}
	parameters = append(parameters, resourceName)
	parameters = append(parameters, resourceVersion)
	parameters = append(parameters, namespace)
	parameters = append(parameters, clusterName)
	parameters = append(parameters, tableName)
	parameters = append(parameters, rangeAttributeName)
	parameters = append(parameters, rangeAttributeType)
	parameters = append(parameters, readCapacityUnits)
	parameters = append(parameters, writeCapacityUnits)
	parameters = append(parameters, hashAttributeName)
	parameters = append(parameters, hashAttributeType)

	stackInputs.SetParameters(parameters)

	resourceNameTag := helpers.CreateTag("ResourceName", s.DynamoDB.Name)
	resourceVersionTag := helpers.CreateTag("ResourceVersion", s.DynamoDB.ResourceVersion)
	namespaceTag := helpers.CreateTag("Namespace", s.DynamoDB.Namespace)
	clusterNameTag := helpers.CreateTag("ClusterName", s.config.ClusterName)

	tags := []*cloudformation.Tag{}
	tags = append(tags, resourceNameTag)
	tags = append(tags, resourceVersionTag)
	tags = append(tags, namespaceTag)
	tags = append(tags, clusterNameTag)

	stackInputs.SetTags(tags)

  output, err = svc.CreateStack(&stackInputs)
	return
}

// UpdateStack will update the existing stack
func (s *Cloudformation) UpdateStack(updated *awsV1alpha1.DynamoDB) (output *cloudformation.UpdateStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	cftemplate := helpers.GetCloudFormationTemplate(s.config, "dynamodb", updated.Spec.CloudFormationTemplateName, updated.Spec.CloudFormationTemplateNamespace)

	stackInputs := cloudformation.UpdateStackInput{
		StackName:   aws.String(s.StackName()),
		TemplateURL: aws.String(cftemplate),
		NotificationARNs: []*string{
			aws.String(s.topicARN),
		},
	}

	resourceName := helpers.CreateParam("ResourceName", s.DynamoDB.Name)
	resourceVersion := helpers.CreateParam("ResourceVersion", s.DynamoDB.ResourceVersion)
	namespace := helpers.CreateParam("Namespace", s.DynamoDB.Namespace)
	clusterName := helpers.CreateParam("ClusterName", s.config.ClusterName)
	tableName := helpers.CreateParam("TableName", helpers.Stringify(s.DynamoDB.Name))
	rangeAttributeNameTemp := "{{.Obj.Spec.RangeAttribute.Name}}"
	rangeAttributeNameValue, err := helpers.Templatize(rangeAttributeNameTemp, helpers.Data{Obj: updated, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	rangeAttributeName := helpers.CreateParam("RangeAttributeName", helpers.Stringify(rangeAttributeNameValue))
	rangeAttributeTypeTemp := "{{.Obj.Spec.RangeAttribute.Type}}"
	rangeAttributeTypeValue, err := helpers.Templatize(rangeAttributeTypeTemp, helpers.Data{Obj: updated, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	rangeAttributeType := helpers.CreateParam("RangeAttributeType", helpers.Stringify(rangeAttributeTypeValue))
	readCapacityUnitsTemp :=	"{{.Obj.Spec.ReadCapacityUnits}}"
	readCapacityUnitsValue, err := helpers.Templatize(readCapacityUnitsTemp, helpers.Data{Obj: updated, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	readCapacityUnits := helpers.CreateParam("ReadCapacityUnits", helpers.Stringify(readCapacityUnitsValue))
	writeCapacityUnitsTemp :=	"{{.Obj.Spec.WriteCapacityUnits}}"
	writeCapacityUnitsValue, err := helpers.Templatize(writeCapacityUnitsTemp, helpers.Data{Obj: updated, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	writeCapacityUnits := helpers.CreateParam("WriteCapacityUnits", helpers.Stringify(writeCapacityUnitsValue))
	hashAttributeNameTemp := "{{.Obj.Spec.HashAttribute.Name}}"
	hashAttributeNameValue, err := helpers.Templatize(hashAttributeNameTemp, helpers.Data{Obj: updated, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	hashAttributeName := helpers.CreateParam("HashAttributeName", helpers.Stringify(hashAttributeNameValue))
	hashAttributeTypeTemp := "{{.Obj.Spec.HashAttribute.Type}}"
	hashAttributeTypeValue, err := helpers.Templatize(hashAttributeTypeTemp, helpers.Data{Obj: updated, Config: s.config, Helpers: helpers.New()})
	if err != nil {
		return output, err
	}
	hashAttributeType := helpers.CreateParam("HashAttributeType", helpers.Stringify(hashAttributeTypeValue))

	parameters := []*cloudformation.Parameter{}
	parameters = append(parameters, resourceName)
	parameters = append(parameters, resourceVersion)
	parameters = append(parameters, namespace)
	parameters = append(parameters, clusterName)
	parameters = append(parameters, tableName)
	parameters = append(parameters, rangeAttributeName)
	parameters = append(parameters, rangeAttributeType)
	parameters = append(parameters, readCapacityUnits)
	parameters = append(parameters, writeCapacityUnits)
	parameters = append(parameters, hashAttributeName)
	parameters = append(parameters, hashAttributeType)

	stackInputs.SetParameters(parameters)

	resourceNameTag := helpers.CreateTag("ResourceName", s.DynamoDB.Name)
	resourceVersionTag := helpers.CreateTag("ResourceVersion", s.DynamoDB.ResourceVersion)
	namespaceTag := helpers.CreateTag("Namespace", s.DynamoDB.Namespace)
	clusterNameTag := helpers.CreateTag("ClusterName", s.config.ClusterName)

	tags := []*cloudformation.Tag{}
	tags = append(tags, resourceNameTag)
	tags = append(tags, resourceVersionTag)
	tags = append(tags, namespaceTag)
	tags = append(tags, clusterNameTag)

	stackInputs.SetTags(tags)

  output, err = svc.UpdateStack(&stackInputs)
	return
}

// DeleteStack will delete the stack
func (s *Cloudformation) DeleteStack() (err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	stackInputs := cloudformation.DeleteStackInput{}
	stackInputs.SetStackName(s.StackName())

  _, err = svc.DeleteStack(&stackInputs)
	return
}

// WaitUntilStackDeleted will delete the stack
func (s *Cloudformation) WaitUntilStackDeleted() (err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	stackInputs := cloudformation.DescribeStacksInput{
		StackName:   aws.String(s.StackName()),
	}

  err = svc.WaitUntilStackDeleteComplete(&stackInputs)
	return
}
