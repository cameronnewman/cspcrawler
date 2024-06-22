package crawler

// Policies ...
type Policies []ContentSecurityPolicy

// ContentSecurityPolicy for a domain
type ContentSecurityPolicy struct {
	Domain                  string            `json:"domain"`
	TLD                     string            `json:"tld"`
	RawPolicy               string            `json:"raw_policy"`
	PolicySource            string            `json:"policy_source"` // header OR meta
	Exists                  bool              `json:"exists"`
	IsValid                 bool              `json:"is_valid"`
	Directives              map[string]string `json:"directives"`
	ReportUri               string            `json:"report_uri"`
	UpgradeInsecureRequests bool              `json:"upgrade_insecure_requests"`
	BlockAllMixedContent    bool              `json:"block_all_mixed_content"`
	Error                   error             `json:"error"`
}
