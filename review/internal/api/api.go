package api

type API struct {
	// router
}

func NewAPI() (*API, error) {
	return &API{}, nil
}

func (api *API) Run(addr string) {
}
