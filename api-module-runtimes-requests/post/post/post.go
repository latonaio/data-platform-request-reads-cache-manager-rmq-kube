package post

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type PostReq struct {
	Header   Header   `json:"Post"`
	Headers  []Header `json:"Posts"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Post                         int     `json:"Post"`
	PostType                     *string `json:"PostType"`
	PostOwner                    *int    `json:"PostOwner"`
	PostOwnerBusinessPartnerRole *string `json:"PostOwnerBusinessPartnerRole"`
	Description                  *string `json:"Description"`
	LongText                     *string `json:"LongText"`
	Site						 *int	 `json:"Site"`
	Tag1                         *string `json:"Tag1"`
	Tag2                         *string `json:"Tag2"`
	Tag3                         *string `json:"Tag3"`
	Tag4                         *string `json:"Tag4"`
	IsPublished					 *bool	 `json:"IsPublished"`
	CreationDate                 *string `json:"CreationDate"`
	CreationTime                 *string `json:"CreationTime"`
	LastChangeDate               *string `json:"LastChangeDate"`
	LastChangeTime               *string `json:"LastChangeTime"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
	InstagramMedia               []InstagramMedia `json:"InstagramMedia"`
	Friend                       []Friend         `json:"Friend"`
}

type InstagramMedia struct {
	Post                            int     `json:"Post"`
	InstagramMediaID                *string `json:"InstagramMediaID"`
	InstagramMediaType              *string `json:"InstagramMediaType"`
	InstagramMediaCaption           *string `json:"InstagramMediaCaption"`
	InstagramMediaPermaLink         *string `json:"InstagramMediaPermaLink"`
	InstagramMediaURL               *string `json:"InstagramMediaURL"`
	InstagramMediaVideoThumbnailURL *string `json:"InstagramMediaVideoThumbnailURL"`
	InstagramMediaTimeStamp         *string `json:"InstagramMediaTimeStamp"`
	InstagramMediaDate              *string `json:"InstagramMediaDate"`
	InstagramMediaTime              *string `json:"InstagramMediaTime"`
	InstagramUserName               *string `json:"InstagramUserName"`
	CreationDate                    *string `json:"CreationDate"`
	CreationTime                    *string `json:"CreationTime"`
	LastChangeDate                  *string `json:"LastChangeDate"`
	LastChangeTime                  *string `json:"LastChangeTime"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

type Friend struct {
	Post                int     `json:"Post"`
	Friend              int     `json:"Friend"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreatePostRequestHeader(
	requestPram *apiInputReader.Request,
	postHeader *apiInputReader.PostHeader,
) PostReq {
	req := PostReq{
		Header: Header{
			Post:                postHeader.Post,
			IsPublished:		 postHeader.IsPublished,
			IsMarkedForDeletion: postHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreatePostRequestHeaders(
	requestPram *apiInputReader.Request,
	postHeaders *apiInputReader.PostHeader,
) PostReq {
	req := PostReq{
		Header: Header{
			IsPublished:		 postHeaders.IsPublished,
			IsMarkedForDeletion: postHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
	}
	return req
}

func CreatePostRequestHeadersByPosts(
	requestPram *apiInputReader.Request,
	postHeaders []Header,
) PostReq {
	req := PostReq{
		Headers: postHeaders,
		Accepter: []string{
			"HeadersByPosts",
		},
	}
	return req
}

func CreatePostRequestHeadersByPostOwner(
	requestPram *apiInputReader.Request,
	postHeaders *apiInputReader.PostHeader,
) PostReq {
	req := PostReq{
		Header: Header{
			PostOwner:           postHeaders.PostOwner,
			IsPublished:		 postHeaders.IsPublished,
			IsMarkedForDeletion: postHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByPostOwner",
		},
	}
	return req
}

func CreatePostRequestInstagramMedia(
	requestPram *apiInputReader.Request,
	postInstagramMedia *apiInputReader.PostInstagramMedia,
) PostReq {
	req := PostReq{
		Header: Header{
			Post: postInstagramMedia.Post,
			InstagramMedia: []InstagramMedia{
				{
					IsMarkedForDeletion: postInstagramMedia.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"InstagramMedia",
		},
	}
	return req
}

func CreatePostRequestInstagramMediasByPosts(
	requestPram *apiInputReader.Request,
	postHeaders []Header,
) PostReq {
	req := PostReq{
		Headers: postHeaders,
		Accepter: []string{
			"InstagramMediasByPosts",
		},
	}
	return req
}

func CreatePostRequestFriend(
	requestPram *apiInputReader.Request,
	postFriend *apiInputReader.PostFriend,
) PostReq {
	req := PostReq{
		Header: Header{
			Post: postFriend.Post,
			Friend: []Friend{
				{
					Friend:              postFriend.Friend,
					IsMarkedForDeletion: postFriend.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Friend",
		},
	}
	return req
}

func CreatePostRequestFriends(
	requestPram *apiInputReader.Request,
	postFriends *apiInputReader.PostFriend,
) PostReq {
	req := PostReq{
		Header: Header{
			Post: postFriends.Post,
			Friend: []Friend{
				{
					IsMarkedForDeletion: postFriends.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Friends",
		},
	}
	return req
}

func PostReadsHeadersByPosts(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "reads"

	request := CreatePostRequestHeadersByPosts(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func PostReadsInstagramMediasByPosts(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "reads"

	request := CreatePostRequestInstagramMediasByPosts(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func PostReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Post,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "reads"

	var request PostReq

	if accepter == "Header" {
		request = CreatePostRequestHeader(
			requestPram,
			&apiInputReader.PostHeader{
				Post:                input.PostHeader.Post,
				IsPublished:		 input.PostHeader.IsPublished,
				IsMarkedForDeletion: input.PostHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Headers" {
		request = CreatePostRequestHeaders(
			requestPram,
			&apiInputReader.PostHeader{
				IsPublished:		 input.PostHeader.IsPublished,
				IsMarkedForDeletion: input.PostHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByPostOwner" {
		request = CreatePostRequestHeadersByPostOwner(
			requestPram,
			&apiInputReader.PostHeader{
				PostOwner:           input.PostHeader.PostOwner,
				IsPublished:		 input.PostHeader.IsPublished,
				IsMarkedForDeletion: input.PostHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "InstagramMedia" {
		request = CreatePostRequestInstagramMedia(
			requestPram,
			&apiInputReader.PostInstagramMedia{
				Post:                input.PostInstagramMedia.Post,
				IsMarkedForDeletion: input.PostInstagramMedia.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Friend" {
		request = CreatePostRequestFriend(
			requestPram,
			&apiInputReader.PostFriend{
				Post:                input.PostFriend.Post,
				Friend:              input.PostFriend.Friend,
				IsMarkedForDeletion: input.PostFriend.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Friends" {
		request = CreatePostRequestFriends(
			requestPram,
			&apiInputReader.PostFriend{
				Post: input.PostFriend.Post,
			},
		)
	}

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}
