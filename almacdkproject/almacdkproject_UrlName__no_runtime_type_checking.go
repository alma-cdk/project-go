//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateUrlName_GloballyParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	return nil
}

func validateUrlName_ItParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	return nil
}

func validateUrlName_WithProjectParameters(scope constructs.Construct, baseName *string, props *NameProps) error {
	return nil
}

