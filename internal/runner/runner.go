package runner

import (
	"strings"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/proxify"
)

// Runner contains the internal logic of the program
type Runner struct {
	options *Options
	proxy   *proxify.Proxy
}

// NewRunner instance
func NewRunner(options *Options) (*Runner, error) {
	proxy, err := proxify.NewProxy(&proxify.Options{
		Silent:                  options.Silent,
		Directory:               options.Directory,
		CertCacheSize:           options.CertCacheSize,
		Verbose:                 options.Verbose,
		ListenAddrHTTP:          options.ListenAddrHTTP,
		ListenAddrSocks5:        options.ListenAddrSocks5,
		OutputDirectory:         options.OutputDirectory,
		RequestDSL:              options.RequestDSL,
		ResponseDSL:             options.ResponseDSL,
		UpstreamHTTPProxy:       options.UpstreamHTTPProxy,
		UpstreamSock5Proxy:      options.UpstreamSocks5Proxy,
		ListenDNSAddr:           options.ListenDNSAddr,
		DNSMapping:              options.DNSMapping,
		DNSFallbackResolver:     options.DNSFallbackResolver,
		RequestMatchReplaceDSL:  options.RequestMatchReplaceDSL,
		ResponseMatchReplaceDSL: options.ResponseMatchReplaceDSL,
		DumpRequest:             options.DumpRequest,
		DumpResponse:            options.DumpResponse,
	})
	if err != nil {
		return nil, err
	}
	return &Runner{options: options, proxy: proxy}, nil
}

// Run polling and notification
func (r *Runner) Run() error {
	// configuration summary
	if r.options.ListenAddrHTTP != "" {
		gologger.Print().Msgf("HTTP Proxy Listening on %s\n", r.options.ListenAddrHTTP)
	}
	if r.options.ListenAddrSocks5 != "" {
		gologger.Print().Msgf("Socks5 Proxy Listening on %s\n", r.options.ListenAddrSocks5)
	}

	if r.options.OutputDirectory != "" {
		gologger.Print().Msgf("Saving traffic to %s\n", r.options.OutputDirectory)
	}

	if r.options.UpstreamHTTPProxy != "" {
		gologger.Print().Msgf("Using upstream HTTP proxy: %s\n", r.options.UpstreamHTTPProxy)
	} else if r.options.UpstreamSocks5Proxy != "" {
		gologger.Print().Msgf("Using upstream SOCKS5 proxy: %s\n", r.options.UpstreamSocks5Proxy)
	}

	if r.options.DNSMapping != "" {
		for _, v := range strings.Split(r.options.DNSMapping, ",") {
			gologger.Print().Msgf("Domain => IP: %s\n", v)
		}

		if r.options.DNSFallbackResolver != "" {
			gologger.Print().Msgf("Fallback Resolver: %s\n", r.options.DNSFallbackResolver)
		}

	}

	return r.proxy.Run()
}

// Close the runner instance
func (r *Runner) Close() {
	r.proxy.Stop()
}
