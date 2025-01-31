package passive

import (
	"github.com/ZhuriLab/Starmap/pkg/subscraping"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/alienvault"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/anubis"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/archiveis"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/binaryedge"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/bufferover"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/c99"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/censys"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/certspotter"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/chaos"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/chinaz"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/commoncrawl"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/crtsh"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/dnsdb"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/dnsdumpster"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/fofa"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/fullhunt"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/github"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/hackertarget"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/intelx"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/passivetotal"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/rapiddns"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/riddler"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/robtex"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/securitytrails"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/shodan"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/sitedossier"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/sonarsearch"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/spyse"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/sublist3r"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/threatbook"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/threatcrowd"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/threatminer"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/virustotal"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/waybackarchive"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/zoomeye"
	"github.com/ZhuriLab/Starmap/pkg/subscraping/sources/zoomeyeapi"
)

// DefaultSources contains the list of fast sources used by default.
var DefaultSources = []string{
	"alienvault",
	"anubis",
	"bufferover",
	"c99",
	"certspotter",
	"censys",
	"chaos",
	"chinaz",
	"crtsh",
	"dnsdumpster",
	"hackertarget",
	"intelx",
	"passivetotal",
	"robtex",
	"riddler",
	"securitytrails",
	"shodan",
	"spyse",
	"sublist3r",
	"threatcrowd",
	"threatminer",
	"virustotal",
	"fofa",
	"fullhunt",
}

// DefaultRecursiveSources contains list of default recursive sources
var DefaultRecursiveSources = []string{
	"alienvault",
	"binaryedge",
	"bufferover",
	"certspotter",
	"crtsh",
	"dnsdumpster",
	"hackertarget",
	"passivetotal",
	"securitytrails",
	"sonarsearch",
	"sublist3r",
	"virustotal",
}

// DefaultAllSources contains list of all sources
var DefaultAllSources = []string{
	"alienvault",
	"anubis",
	"archiveis",
	"binaryedge",
	"bufferover",
	"c99",
	"censys",
	"certspotter",
	"chaos",
	"commoncrawl",
	"crtsh",
	"dnsdumpster",
	"dnsdb",
	"github",
	"hackertarget",
	"intelx",
	"passivetotal",
	"rapiddns",
	"riddler",
	"robtex",
	"securitytrails",
	"shodan",
	"sitedossier",
	"sonarsearch",
	"spyse",
	"sublist3r",
	"threatbook",
	"threatcrowd",
	"threatminer",
	"virustotal",
	"waybackarchive",
	"zoomeye",
	"zoomeyeapi",
	"fofa",
	"fullhunt",
}

// Agent is a struct for running passive subdomain enumeration
// against a given host. It wraps subscraping package and provides
// a layer to build upon.
type Agent struct {
	sources map[string]subscraping.Source
}

// New creates a new agent for passive subdomain discovery
func New(sources, exclusions []string) *Agent {
	// Create the agent, insert the sources and remove the excluded sources
	agent := &Agent{sources: make(map[string]subscraping.Source)}

	agent.addSources(sources)
	agent.removeSources(exclusions)

	return agent
}

// addSources adds the given list of sources to the source array
func (a *Agent) addSources(sources []string) {
	for _, source := range sources {
		switch source {
		case "alienvault":
			a.sources[source] = &alienvault.Source{}
		case "anubis":
			a.sources[source] = &anubis.Source{}
		case "archiveis":
			a.sources[source] = &archiveis.Source{}
		case "binaryedge":
			a.sources[source] = &binaryedge.Source{}
		case "bufferover":
			a.sources[source] = &bufferover.Source{}
		case "c99":
			a.sources[source] = &c99.Source{}
		case "censys":
			a.sources[source] = &censys.Source{}
		case "certspotter":
			a.sources[source] = &certspotter.Source{}
		case "chaos":
			a.sources[source] = &chaos.Source{}
		case "chinaz":
			a.sources[source] = &chinaz.Source{}
		case "commoncrawl":
			a.sources[source] = &commoncrawl.Source{}
		case "crtsh":
			a.sources[source] = &crtsh.Source{}
		case "dnsdumpster":
			a.sources[source] = &dnsdumpster.Source{}
		case "dnsdb":
			a.sources[source] = &dnsdb.Source{}
		case "github":
			a.sources[source] = &github.Source{}
		case "hackertarget":
			a.sources[source] = &hackertarget.Source{}
		case "intelx":
			a.sources[source] = &intelx.Source{}
		case "passivetotal":
			a.sources[source] = &passivetotal.Source{}
		case "rapiddns":
			a.sources[source] = &rapiddns.Source{}
		case "riddler":
			a.sources[source] = &riddler.Source{}
		case "robtex":
			a.sources[source] = &robtex.Source{}
		case "securitytrails":
			a.sources[source] = &securitytrails.Source{}
		case "shodan":
			a.sources[source] = &shodan.Source{}
		case "sitedossier":
			a.sources[source] = &sitedossier.Source{}
		case "sonarsearch":
			a.sources[source] = &sonarsearch.Source{}
		case "spyse":
			a.sources[source] = &spyse.Source{}
		case "sublist3r":
			a.sources[source] = &sublist3r.Source{}
		case "threatbook":
			a.sources[source] = &threatbook.Source{}
		case "threatcrowd":
			a.sources[source] = &threatcrowd.Source{}
		case "threatminer":
			a.sources[source] = &threatminer.Source{}
		case "virustotal":
			a.sources[source] = &virustotal.Source{}
		case "waybackarchive":
			a.sources[source] = &waybackarchive.Source{}
		case "zoomeye":
			a.sources[source] = &zoomeye.Source{}
		case "zoomeyeapi":
			a.sources[source] = &zoomeyeapi.Source{}
		case "fofa":
			a.sources[source] = &fofa.Source{}
		case "fullhunt":
			a.sources[source] = &fullhunt.Source{}
		}
	}
}

// removeSources deletes the given sources from the source map
func (a *Agent) removeSources(sources []string) {
	for _, source := range sources {
		delete(a.sources, source)
	}
}
