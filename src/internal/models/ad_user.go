package models

type ADUser struct {
	UID               string
	Username          string
	Email             string
	DN                string
	CN                string
	SAMAccountName    string
	Groups            []string
	UserPrincipalName string
}
