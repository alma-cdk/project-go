// Opinionated CDK Project “Framework”
package almacdkproject


// Availalbe Enviroment Categories.
//
// Categories are useful grouping to make distinction between `stable`
// environments (`staging` & `production`) from `feature` or `verification`
// environments (such as `test` or `preproduction`).
// Experimental.
type EnvironmentCategory string

const (
	// Experimental.
	EnvironmentCategory_MOCK EnvironmentCategory = "MOCK"
	// Experimental.
	EnvironmentCategory_DEVELOPMENT EnvironmentCategory = "DEVELOPMENT"
	// Experimental.
	EnvironmentCategory_FEATURE EnvironmentCategory = "FEATURE"
	// Experimental.
	EnvironmentCategory_VERIFICATION EnvironmentCategory = "VERIFICATION"
	// Experimental.
	EnvironmentCategory_STABLE EnvironmentCategory = "STABLE"
)

