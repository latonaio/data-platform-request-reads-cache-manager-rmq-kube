package apiOutputFormatter

type DistributionProfile struct {
	DistributionProfileDistributionProfile    []DistributionProfileDistributionProfile    `json:"DistributionProfileDistributionProfile"`
	DistributionProfileText                   []DistributionProfileText                   `json:"DistributionProfileText"`
	Accepter                                  []string                                    `json:"Accepter"`
}

type DistributionProfileDistributionProfile struct {
	DistributionProfile        string	`json:"DistributionProfile"`
}

type DistributionProfileText struct {
	DistributionProfile        string `json:"DistributionProfile"`
	Language                   string `json:"Language"`
	DistributionProfileName    string `json:"DistributionProfileName"`
}
