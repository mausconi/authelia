package authentication

// UserDetails represent the details retrieved for a given user.
type UserDetails struct {
	Emails []string
	Groups []string
}

// UserProvider is the interface for checking user password and
// gathering user details.
type UserProvider interface {
	CheckUserPassword(username string, password string) (bool, error)
	GetDetails(username string) (*UserDetails, error)
	UpdatePassword(username string, newPassword string) error
}
