//go:build !no_runtime_type_checking

package almacdkproject

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateEnvironmentType_GetParameters(scope constructs.Construct, allowedEnvironments *[]*string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if allowedEnvironments == nil {
		return fmt.Errorf("parameter allowedEnvironments is required, but nil was provided")
	}

	return nil
}

func validateEnvironmentType_SetParameters(scope constructs.Construct, environmentType *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if environmentType == nil {
		return fmt.Errorf("parameter environmentType is required, but nil was provided")
	}

	return nil
}

func validateEnvironmentType_TryGetParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

