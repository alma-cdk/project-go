//go:build !no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateEnvironmentWrapper_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewEnvironmentWrapperParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

