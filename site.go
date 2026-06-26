// Package memoriessongcom exposes metadata for wan27.org.
package memoriessongcom

// SiteInfo describes the public website and local publishing paths.
type SiteInfo struct {
	Name            string
	Domain          string
	Homepage        string
	Documentation   string
	Package         string
	Repository      string
	LocalRepository string
	ContentPath     string
	AppPath         string
	Focus           string
}

const (
	Name            = "Wan27"
	Domain          = "wan27.org"
	Homepage        = "https://wan27.org"
	Documentation   = "https://wan27.org/docs"
	Package         = "memoriessong-com"
	Repository      = "https://github.com/youram470-art/memoriessong-com-go"
	LocalRepository = "/Users/mac/Documents/code/foreversong"
	ContentPath     = "content"
	AppPath         = "src/app"
	Focus           = "AI video generation, prompt-to-video workflows, and Wan27 site metadata"
)

// Hello returns a small installation check string.
func Hello() string {
	return "hello from wan27.org"
}

// GetSiteInfo returns metadata for wan27.org.
func GetSiteInfo() SiteInfo {
	return SiteInfo{
		Name:            Name,
		Domain:          Domain,
		Homepage:        Homepage,
		Documentation:   Documentation,
		Package:         Package,
		Repository:      Repository,
		LocalRepository: LocalRepository,
		ContentPath:     ContentPath,
		AppPath:         AppPath,
		Focus:           Focus,
	}
}
