package apiOutputFormatter

type Post struct {
	PostHeader                   []PostHeader                   `json:"PostHeader"`
	PostInstagramMedia           []PostInstagramMedia           `json:"PostInstagramMedia"`
	PostHeaderWithInstagramMedia []PostHeaderWithInstagramMedia `json:"PostHeaderWithInstagramMedia"`
	PostFriend                   []PostFriend                   `json:"PostFriend"`
	MountPath                    *string                        `json:"mount_path"`
	Accepter                     []string                       `json:"Accepter"`
}

type PostHeader struct {
	Post                             int     `json:"Post"`
	PostType                         string  `json:"PostType"`
	PostOwner                        int     `json:"PostOwner"`
	PostOwnerNickName                string  `json:"PostOwnerNickName"`
	PostOwnerBusinessPartnerRole     string  `json:"PostOwnerBusinessPartnerRole"`
	PostOwnerBusinessPartnerRoleName string  `json:"PostOwnerBusinessPartnerRoleName"`
	Description                      *string `json:"Description"`
	LongText                         string  `json:"LongText"`
	Site                             *int    `json:"Site"`
	Tag1                             *string `json:"Tag1"`
	Tag2                             *string `json:"Tag2"`
	Tag3                             *string `json:"Tag3"`
	Tag4                             *string `json:"Tag4"`
	IsPublished                      bool    `json:"IsPublished"`
	CreationDate                     string  `json:"CreationDate"`
	CreationTime                     string  `json:"CreationTime"`
	LastChangeDate                   string  `json:"LastChangeDate"`
	LastChangeTime                   string  `json:"LastChangeTime"`
	Images                           Images  `json:"Images"`
}

type PostHeaderWithInstagramMedia struct {
	Post                    int     `json:"Post"`
	PostOwner               int     `json:"PostOwner"`
	Description             *string `json:"Description"`
	LongText                string  `json:"LongText"`
	Tag1                    *string `json:"Tag1"`
	Tag2                    *string `json:"Tag2"`
	Tag3                    *string `json:"Tag3"`
	Tag4                    *string `json:"Tag4"`
	IsPublished             bool    `json:"IsPublished"`
	CreationDate            string  `json:"CreationDate"`
	CreationTime            string  `json:"CreationTime"`
	LastChangeDate          string  `json:"LastChangeDate"`
	LastChangeTime          string  `json:"LastChangeTime"`
	InstagramMediaID        string  `json:"InstagramMediaID"`
	InstagramMediaType      string  `json:"InstagramMediaType"`
	InstagramMediaCaption   *string `json:"InstagramMediaCaption"`
	InstagramMediaPermaLink string  `json:"InstagramMediaPermaLink"`
	InstagramMediaURL       string  `json:"InstagramMediaURL"`
	InstagramMediaDate      string  `json:"InstagramMediaDate"`
	InstagramMediaTime      string  `json:"InstagramMediaTime"`
	InstagramUserName       string  `json:"InstagramUserName"`
	Images                  Images  `json:"Images"`
}

type PostInstagramMedia struct {
	Post                            int     `json:"Post"`
	InstagramMediaID                string  `json:"InstagramMediaID"`
	InstagramMediaType              string  `json:"InstagramMediaType"`
	InstagramMediaCaption           *string `json:"InstagramMediaCaption"`
	InstagramMediaPermaLink         string  `json:"InstagramMediaPermaLink"`
	InstagramMediaURL               string  `json:"InstagramMediaURL"`
	InstagramMediaVideoThumbnailURL *string `json:"InstagramMediaVideoThumbnailURL"`
	InstagramMediaTimeStamp         string  `json:"InstagramMediaTimeStamp"`
	InstagramMediaDate              string  `json:"InstagramMediaDate"`
	InstagramMediaTime              string  `json:"InstagramMediaTime"`
	InstagramUserName               string  `json:"InstagramUserName"`
}

type PostFriend struct {
	Post           int    `json:"Post"`
	Friend         int    `json:"Friend"`
	FriendNickName string `json:"FriendNickName"`
}
