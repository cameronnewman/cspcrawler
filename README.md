# cspcrawler

[![Build Status](https://travis-ci.org/cameronnewman/cli.crawler.svg?branch=master)](https://travis-ci.org/cameronnewman/cli.crawler) [![GoDoc](https://godoc.org/github.com/cameronnewman/cspcrawler?status.svg)](http://godoc.org/github.com/cameronnewman/cspcrawler) [![Report card](https://goreportcard.com/badge/github.com/cameronnewman/cspcrawler)](https://goreportcard.com/report/github.com/cameronnewman/cspcrawler)

a simple Content Security Policy crawler


Usage

```
lappy:~ root$ ./cspcrawler --url https://google.com
[
	{
		"domain": "https://google.com",
		"tld": "com",
		"raw_policy": "object-src 'none';base-uri 'self';script-src 'nonce-KZgUAjcFC3GXCzOBMu5RLg' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp",
		"policy_source": "header",
		"exists": true,
		"is_valid": true,
		"directives": {
			"base-uri": "'self'",
			"object-src": "'none'",
			"script-src": "'nonce-KZgUAjcFC3GXCzOBMu5RLg' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:"
		},
		"report_uri": "https://csp.withgoogle.com/csp/gws/other-hp",
		"upgrade_insecure_requests": false,
		"block_all_mixed_content": false,
		"error": null
	}
]

```
