package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cameronnewman/cspcrawler/csp"
	"github.com/goware/urlx"
	"github.com/joeguo/tldextract"
	"net/http"
	"net/url"
	"os"
)

const (
	domainSuffixListFile string = "public_suffix_list.dat"
)

// Crawler struct
type Crawler struct {
	SeedRawURL string
	SeedURL    *url.URL
	tldExtract *tldextract.TLDExtract
}

// New Crawler instance
func New(u string) (*Crawler, error) {

	rawURL, err := url.Parse(u)
	if err != nil {
		return &Crawler{}, err
	}

	c := &Crawler{
		SeedRawURL: u,
		SeedURL:    rawURL,
	}

	tldExtract, err := tldextract.New(domainSuffixListFile, false)
	if err != nil {
		return c, err
	}

	c.tldExtract = tldExtract

	return c, nil
}

// Run starts crawling the URL
func (c *Crawler) Run() {

	result := Policies{}

	tld, err := c.tld(c.SeedRawURL)
	if err != nil {
		writeToStdOut([]ContentSecurityPolicy{
			{
				Exists:       false,
				IsValid:      false,
				PolicySource: "header", // header OR meta
				Domain:       c.SeedRawURL,
				Error:        err,
			},
		})
	}

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, c.SeedRawURL, nil)
	if err != nil {
		writeToStdOut([]ContentSecurityPolicy{
			{
				Exists:       false,
				IsValid:      false,
				PolicySource: "header", // header OR meta
				TLD:          tld,
				Domain:       c.SeedRawURL,
				Error:        err,
			},
		})
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		writeToStdOut([]ContentSecurityPolicy{
			{
				Exists:       false,
				IsValid:      false,
				PolicySource: "header", // header OR meta
				TLD:          tld,
				Domain:       c.SeedRawURL,
				Error:        err,
			},
		})
	}

	if resp.StatusCode != http.StatusOK {
		writeToStdOut([]ContentSecurityPolicy{
			{
				Exists:       false,
				IsValid:      false,
				PolicySource: "header", // header OR meta
				TLD:          tld,
				Domain:       c.SeedRawURL,
				Error:        errors.New("http request failed, " + resp.Status),
			},
		})
	}

	rawPolicy, ok := c.extractCSP(resp.Header)
	if !ok {
		writeToStdOut([]ContentSecurityPolicy{
			{
				Exists:       false,
				IsValid:      false,
				PolicySource: "header", // header OR meta
				TLD:          tld,
				Domain:       c.SeedRawURL,
				Error:        err,
			},
		})
	}

	policy, err := csp.ParsePolicy(rawPolicy)
	if err != nil {
		writeToStdOut([]ContentSecurityPolicy{
			{
				Exists:       false,
				IsValid:      false,
				PolicySource: "header", // header OR meta
				TLD:          tld,
				Domain:       c.SeedRawURL,
				Error:        err,
			},
		})
	}

	directives := map[string]string{}
	for _, k := range csp.AllDirectives {
		d, ok := policy.Directives[k]
		if ok {
			directives[k] = d.Get()
		}
	}

	result = append(result, ContentSecurityPolicy{
		Exists:                  ok,
		IsValid:                 true,
		PolicySource:            "header", // header OR meta
		RawPolicy:               rawPolicy,
		Directives:              directives,
		ReportUri:               policy.ReportUri,
		UpgradeInsecureRequests: policy.UpgradeInsecureRequests,
		BlockAllMixedContent:    policy.BlockAllMixedContent,
		TLD:                     tld,
		Domain:                  c.SeedRawURL,
	})

	writeToStdOut(result)
}

func (c *Crawler) extractCSP(headers http.Header) (string, bool) {
	const (
		headerContentSecurityPolicy           string = "Content-Security-Policy"
		headerContentSecurityPolicyReportOnly string = "Content-Security-Policy-Report-Only"
	)

	csp := headers.Get(headerContentSecurityPolicy)
	cspro := headers.Get(headerContentSecurityPolicyReportOnly)

	result := ""

	if len(csp) != 0 {
		result = csp
	}

	if len(cspro) != 0 {
		result = cspro
	}

	if len(result) == 0 {
		return result, false
	}

	return result, true
}

func (c *Crawler) tld(s string) (string, error) {
	u, err := urlx.Parse(s)
	if err != nil {
		return "", err
	}

	domain := c.tldExtract.Extract(u.Host)
	return domain.Tld, nil
}

func writeToStdOut(result Policies) {
	b, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	os.Exit(0)
}
