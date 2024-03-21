//go:build !no_runtime_type_checking

package almacdkproject

import (
	"fmt"
)

func (e *jsiiProxy_EnvRegExp) validateTestParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNewEnvRegExpParameters(base *string) error {
	if base == nil {
		return fmt.Errorf("parameter base is required, but nil was provided")
	}

	return nil
}

