package csp

/*
	"base-uri", "child-src", "connect-src", "default-src",
	"font-src", "form-action", "frame-ancestors", "frame-src",
	"img-src", "manifest-src", "media-src", "object-src",
	"script-src", "style-src", "worker-src", "report-uri",
	"upgrade-insecure-requests", "block-all-mixed-content"
*/

const (
	DirectiveBaseUri                 string = "base-uri"
	DirectiveChildSrc                string = "child-src"
	DirectiveConnectSrc              string = "connect-src"
	DirectiveDefaultSrc              string = "default-src"
	DirectiveFontSrc                 string = "font-src"
	DirectiveFormAction              string = "form-action"
	DirectiveFrameAncestors          string = "frame-ancestors"
	DirectiveFrameSrc                string = "frame-src"
	DirectiveImgSrc                  string = "img-src"
	DirectiveManifestSrc             string = "manifest-src"
	DirectiveMediaSrc                string = "media-src"
	DirectiveObjectSrc               string = "object-src"
	DirectiveScriptSrc               string = "script-src"
	DirectiveStyleSrc                string = "style-src"
	DirectiveWorkerSrc               string = "worker-src"
	DirectiveReportUri               string = "report-uri"
	DirectiveUpgradeInsecureRequests string = "upgrade-insecure-requests"
	DirectiveBlockAllMixedContent    string = "block-all-mixed-content"
)

var AllDirectives []string = []string{DirectiveBaseUri, DirectiveChildSrc, DirectiveConnectSrc, DirectiveDefaultSrc,
	DirectiveFontSrc, DirectiveFormAction, DirectiveFrameAncestors, DirectiveFrameSrc, DirectiveImgSrc, DirectiveManifestSrc,
	DirectiveMediaSrc, DirectiveObjectSrc, DirectiveScriptSrc, DirectiveStyleSrc, DirectiveWorkerSrc, DirectiveReportUri,
	DirectiveUpgradeInsecureRequests, DirectiveBlockAllMixedContent}
