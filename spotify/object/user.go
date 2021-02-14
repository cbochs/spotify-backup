package object

// PrivateUser represents PrivateUserObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-privateuserobject
type PrivateUser struct {
	DisplayName     string                  `json:"display_name,omitempty"`
	ID              string                  `json:"id"`
	Country         string                  `json:"country"`
	Email           string                  `json:"email"`
	ExplicitContent ExplicitContentSettings `json:"explicit_content"`
	ExternalURLs    ExternalURL             `json:"external_urls"`
	Followers       *Followers              `json:"followers,omitempty"`
	HRef            string                  `json:"href"`
	Images          []Image                 `json:"images,omitempty"`
	Product         string                  `json:"product"`
	Type            objectType              `json:"type"`
	URI             string                  `json:"uri"`
}

// PublicUser represents PublicUserObject
// Link: https://developer.spotify.com/documentation/web-api/reference/#object-publicuserobject
type PublicUser struct {
	DisplayName  string      `json:"display_name,omitempty"`
	ID           string      `json:"id"`
	ExternalURLs ExternalURL `json:"external_urls"`
	Followers    *Followers  `json:"followers,omitempty"`
	HRef         string      `json:"href"`
	Images       []Image     `json:"images,omitempty"`
	Type         objectType  `json:"type"`
	URI          string      `json:"uri"`
}
