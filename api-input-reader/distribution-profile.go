package apiInputReader

type DistributionProfileGlobal struct {
	DistributionProfile     *DistributionProfile
	DistributionProfileText *DistributionProfileText
}

type DistributionProfile struct {
	DistributionProfile    string `json:"DistributionProfile"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}

type DistributionProfileText struct {
	DistributionProfile    string `json:"DistributionProfile"`
	Language               string `json:"Language"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}
