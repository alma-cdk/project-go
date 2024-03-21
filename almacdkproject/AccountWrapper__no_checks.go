//go:build no_runtime_type_checking

package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateAccountWrapper_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAccountWrapperParameters(scope constructs.Construct) error {
	return nil
}

