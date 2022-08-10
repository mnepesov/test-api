package nasa

type INasa interface {
}

type Nasa struct {
	ApiKey string
}

func NewNasa(apiKey string) INasa {
	return &Nasa{
		ApiKey: apiKey,
	}
}
