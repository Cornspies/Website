package templates

type SiteLocation string

const (
	LocationHome       SiteLocation = "start-home"
	LocationLinks      SiteLocation = "start-links"
	LocationForum      SiteLocation = "start-forum"
	LocationMusic      SiteLocation = "start-music"
	LocationMovies     SiteLocation = "start-movies"
	LocationTournament SiteLocation = "start-tournament"
	LocationProfile    SiteLocation = "start-profile"
	LocationNone       SiteLocation = ""
)

type M map[string]interface{}

type Page struct {
	Title           string
	CurrentLocation SiteLocation
}
