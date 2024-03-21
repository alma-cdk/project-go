//go:build !no_runtime_type_checking

package almacdkproject

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateUrlName_GloballyParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if baseName == nil {
		return fmt.Errorf("parameter baseName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateUrlName_ItParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if baseName == nil {
		return fmt.Errorf("parameter baseName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateUrlName_WithProjectParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if baseName == nil {
		return fmt.Errorf("parameter baseName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

