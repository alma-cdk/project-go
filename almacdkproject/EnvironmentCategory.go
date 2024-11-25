package almacdkproject


// Availalbe Enviroment Categories.
//
// Categories are useful grouping to make distinction between `stable`
// environments (`staging` & `production`) from `feature` or `verification`
// environments (such as `test` or `preproduction`).
type EnvironmentCategory string

const (
	EnvironmentCategory_MOCK EnvironmentCategory = "MOCK"
	EnvironmentCategory_DEVELOPMENT EnvironmentCategory = "DEVELOPMENT"
	EnvironmentCategory_FEATURE EnvironmentCategory = "FEATURE"
	EnvironmentCategory_VERIFICATION EnvironmentCategory = "VERIFICATION"
	EnvironmentCategory_STABLE EnvironmentCategory = "STABLE"
)

