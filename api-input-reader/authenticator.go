package apiInputReader

type Authenticator struct {
	AuthenticatorUser    *AuthenticatorUser
}

type AuthenticatorUser struct {
	UserID                string    `json:"UserID"`
	IsMarkedForDeletion   *bool     `json:"IsMarkedForDeletion"`
}
