//go:build !no_runtime_type_checking

package almacdkproject

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateAccountType_GetParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateAccountType_MatchFromEnvironmentParameters(scope constructs.Construct, accounts *map[string]*Account, environmentType *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if accounts == nil {
		return fmt.Errorf("parameter accounts is required, but nil was provided")
	}
	for idx_bc62a3, v := range *accounts {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter accounts[%#v]", idx_bc62a3) }); err != nil {
			return err
		}
	}

	if environmentType == nil {
		return fmt.Errorf("parameter environmentType is required, but nil was provided")
	}

	return nil
}

func validateAccountType_SetParameters(scope constructs.Construct, accountType *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if accountType == nil {
		return fmt.Errorf("parameter accountType is required, but nil was provided")
	}

	return nil
}

