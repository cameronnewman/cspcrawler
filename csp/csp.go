package csp

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gobwas/glob"
	"github.com/pkg/errors"
)

// Policy represents the entire CSP policy and its directives.
type Policy struct {
	Directives              map[string]Directive
	ReportUri               string
	UpgradeInsecureRequests bool
	BlockAllMixedContent    bool
}

// ParsePolicy parses all the directives in a CSP policy.
func ParsePolicy(policy string) (Policy, error) {
	p := Policy{
		Directives: map[string]Directive{},
	}

	fmt.Println(policy)

	policy = strings.TrimSpace(policy)

	directiveDefs := strings.Split(policy, ";")
	for _, directive := range directiveDefs {

		if len(directive) == 0 {
			continue
		}

		fields := strings.Fields(directive)
		if len(fields) == 0 {
			return Policy{}, errors.Errorf("empty directive field: %q", directive)
		}

		directiveType := fields[0]
		fmt.Println(directiveType)

		switch directiveType {

		case DirectiveBaseUri, DirectiveChildSrc, DirectiveConnectSrc, DirectiveDefaultSrc, DirectiveFontSrc,
			DirectiveFormAction, DirectiveFrameAncestors, DirectiveFrameSrc, DirectiveImgSrc, DirectiveManifestSrc,
			DirectiveMediaSrc, DirectiveObjectSrc, DirectiveScriptSrc, DirectiveStyleSrc, DirectiveWorkerSrc:
			d, err := ParseSourceDirective(fields[1:])
			if err != nil {
				return Policy{}, err
			}
			p.Directives[directiveType] = d

		case DirectiveReportUri:
			if len(fields) != 2 {
				return Policy{}, errors.Errorf("report-uri expects 1 field; got %q", directive)
			}
			if _, err := url.Parse(fields[1]); err != nil {
				return Policy{}, err
			}
			p.ReportUri = fields[1]

		case DirectiveUpgradeInsecureRequests:

			if len(fields) != 1 {
				return Policy{}, errors.Errorf("upgrade-insecure-requests expects 0 field; got %q", directive)
			}
			p.UpgradeInsecureRequests = true

		case DirectiveBlockAllMixedContent:
			if len(fields) != 1 {
				return Policy{}, errors.Errorf("block-all-mixed-content expects 0 field; got %q", directive)
			}
			p.BlockAllMixedContent = true

		default:
			return Policy{}, errors.Errorf("unknown directive %q", directive)
		}
	}

	return p, nil
}

// Directive returns the first directive that exists in the order: directive
// with the provided name, default-src, and finally 'none' directive.
func (p Policy) Directive(name string) Directive {
	d, ok := p.Directives[name]
	if ok {
		return d
	}

	// frame-ancestors defaults to always allow.
	if name == "frame-ancestors" {
		return AllowDirective{}
	}

	d, ok = p.Directives["default-src"]
	if ok {
		return d
	}

	// If no directives use default policy.
	g, err := glob.Compile("*://*")
	if err != nil {
		panic(err)
	}
	return SourceDirective{
		Hosts: []glob.Glob{g},
	}
}
