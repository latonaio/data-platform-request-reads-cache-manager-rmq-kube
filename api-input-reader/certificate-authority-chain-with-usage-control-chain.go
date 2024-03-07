package apiInputReader

type CertificateAuthorityChainWithUsageControlChainGlobal struct {
	CertificateAuthorityChainWithUsageControlChain *CertificateAuthorityChainWithUsageControlChain
}

type CertificateAuthorityChainWithUsageControlChain struct {
	CertificateAuthorityChain string `json:"CertificateAuthorityChain"`
	CertificateObject         string `json:"CertificateObject"`
	CertificateObjectLabel    string `json:"CertificateObjectLabel"`
	UsageControlChain         string `json:"UsageControlChain"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}
