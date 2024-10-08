package apiModuleRuntimesResponsesPost

type PostRes struct {
	Message Post `json:"message,omitempty"`
}

type Post struct {
	Header             *[]Header			`json:"Header"`
	InstagramMedia	   *[]InstagramMedia	`json:"InstagramMedia"`
	Friend             *[]Friend			`json:"Friend"`
}

type Header struct {
	Post							int		`json:"Post"`
	PostType						string	`json:"PostType"`
	PostOwner						int		`json:"PostOwner"`
	PostOwnerBusinessPartnerRole	string	`json:"PostOwnerBusinessPartnerRole"`
	Description						*string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Site                         	*int    `json:"Site"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	IsPublished					 	bool	`json:"IsPublished"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}

type InstagramMedia struct {
	Post                 			int		`json:"Post"`
	InstagramMediaID   				string  `json:"InstagramMediaID"`
	InstagramMediaType				string  `json:"InstagramMediaType"`
	InstagramMediaCaption			*string `json:"InstagramMediaCaption"`
	InstagramMediaPermaLink   		string	`json:"InstagramMediaPermaLink"`
	InstagramMediaURL				string  `json:"InstagramMediaURL"`
	InstagramMediaVideoThumbnailURL	*string `json:"InstagramMediaVideoThumbnailURL"`
	InstagramMediaTimeStamp			string  `json:"InstagramMediaTimeStamp"`
	InstagramMediaDate				string  `json:"InstagramMediaDate"`
	InstagramMediaTime				string  `json:"InstagramMediaTime"`
	InstagramUserName				string  `json:"InstagramUserName"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}

type Friend struct {
	Post                 	int     `json:"Post"`
	Friend		            int		`json:"Friend"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}

