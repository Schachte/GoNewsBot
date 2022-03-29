package src

type Credentials struct {
	ClientId     string
	ClientSecret string
	Extra        string
}

type Sink interface {
	GetAuthentication() *Credentials
	GetUri() string
	Upload(p *Post) (bool, error)
}
