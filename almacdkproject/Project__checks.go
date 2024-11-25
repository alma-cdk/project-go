//go:build !no_runtime_type_checking

package almacdkproject

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_Project) validateAcknowledgeWarningsParameters(acknowledgements *[]*Acknowledgeable) error {
	if acknowledgements == nil {
		return fmt.Errorf("parameter acknowledgements is required, but nil was provided")
	}
	for idx_4fccc8, v := range *acknowledgements {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter acknowledgements[%#v]", idx_4fccc8) }); err != nil {
			return err
		}
	}

	return nil
}

func (p *jsiiProxy_Project) validateSynthParameters(options *awscdk.StageSynthesisOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateProject_GetAccountParameters(scope constructs.Construct, accountType *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if accountType == nil {
		return fmt.Errorf("parameter accountType is required, but nil was provided")
	}

	return nil
}

func validateProject_GetConfigurationParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateProject_IsAppParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateProject_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateProject_IsStageParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateProject_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewProjectParameters(props *ProjectProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

