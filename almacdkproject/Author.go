package almacdkproject


// Author information.
//
// I.e. who owns/develops this project/service.
type Author struct {
	// Human-readable name for the team/contact responsible for this project/service.
	//
	// Example:
	//   'Mad Scientists'
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email address for the team/contact responsible for this project/service.
	//
	// Example:
	//   'mad.scientists@acme.example.com'
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Human-readable name for the organization responsible for this project/service.
	//
	// Example:
	//   'Acme Corp'
	//
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

