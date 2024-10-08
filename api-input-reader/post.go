package apiInputReader

type Post struct {
	PostHeader            *PostHeader
	PostInstagramMedia    *PostInstagramMedia
	PostFriend            *PostFriend
	PostDocHeaderDoc      *PostDocHeaderDoc
}

type PostHeader struct {
	Post                     int     `json:"Post"`
	PostOwner                *int    `json:"PostOwner"`
	IsPublished				 *bool	 `json:"IsPublished"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type PostInstagramMedia struct {
	Post                     int     `json:"Post"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type PostFriend struct {
	Post				    int		`json:"Post"`
	Friend                  int		`json:"Friend"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}

type PostDocHeaderDoc struct {
	Post            		 int     `json:"Post"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
