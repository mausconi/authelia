package schema

// ACLRule represent one ACL rule
type ACLRule struct {
	Domain    string   `yaml:"domain"`
	Policy    string   `yaml:"policy"`
	Subject   string   `yaml:"subject"`
	Networks  []string `yaml:"networks"`
	Resources []string `yaml:"resources"`
}

// AccessControlConfiguration represents the configuration related to ACLs.
type AccessControlConfiguration struct {
	DefaultPolicy string    `yaml:"default_policy"`
	Rules         []ACLRule `yaml:"rules"`
}
