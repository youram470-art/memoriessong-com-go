// Package memoriessongcom exposes metadata for memoriessong.com.
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
	Name            = "Memories Song"
	Domain          = "memoriessong.com"
	Homepage        = "https://memoriessong.com"
	Documentation   = "https://memoriessong.com/docs"
	Package         = "memoriessong-com"
	Repository      = "https://github.com/youram470-art/memoriessong-com-go"
	LocalRepository = "/Users/mac/Documents/code/foreversong"
	ContentPath     = "content"
	AppPath         = "src/app"
	Focus           = "AI keepsake music, memory-song creation, and emotional tribute song workflows"
)

// Hello returns a small installation check string.
func Hello() string {
	return "hello from memoriessong.com"
}

// GetSiteInfo returns metadata for memoriessong.com.
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
