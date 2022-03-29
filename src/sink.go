package src

type Credentials struct {
	ClientId     string
	ClientSecret string
	Extra        string
}

type Sink interface {
	GetAuthentication() *Credentials
	Upload(p *Post) (bool, error)
}
