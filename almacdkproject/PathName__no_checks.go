//go:build no_runtime_type_checking

package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validatePathName_GloballyParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	return nil
}

func validatePathName_ItParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	return nil
}

func validatePathName_WithProjectParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	return nil
}

