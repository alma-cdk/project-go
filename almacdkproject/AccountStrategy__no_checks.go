//go:build no_runtime_type_checking

package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateAccountStrategy_OneParameters(props *AccountStrategyOneProps) error {
	return nil
}

func validateAccountStrategy_ThreeParameters(props *AccountStrategyThreeProps) error {
	return nil
}

func validateAccountStrategy_TwoParameters(props *AccountStrategyTwoProps) error {
	return nil
}

