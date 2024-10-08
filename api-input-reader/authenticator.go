package apiInputReader

type Authenticator struct {
	AuthenticatorUser						*AuthenticatorUser
	AuthenticatorInitialEmailAuth			*AuthenticatorInitialEmailAuth
	AuthenticatorInitialSMSAuth				*AuthenticatorInitialSMSAuth
	AuthenticatorSMSAuth					*AuthenticatorSMSAuth
	AuthenticatorInitialGoogleAccountAuth	*AuthenticatorInitialGoogleAccountAuth
	AuthenticatorGoogleAccountAuth			*AuthenticatorGoogleAccountAuth
	AuthenticatorInstagramAuth				*AuthenticatorInstagramAuth
}

type AuthenticatorUser struct {
	UserID              string `json:"UserID"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type AuthenticatorInitialEmailAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type AuthenticatorInitialSMSAuth struct {
	MobilePhoneNumber   string `json:"MobilePhoneNumber"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type AuthenticatorSMSAuth struct {
	UserID              string `json:"UserID"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type AuthenticatorInitialGoogleAccountAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	IsMarkedForDeletion *bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorGoogleAccountAuth struct {
	UserID              string `json:"UserID"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type AuthenticatorInstagramAuth struct {
	UserID              string `json:"UserID"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
