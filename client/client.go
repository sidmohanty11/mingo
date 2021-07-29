package client

type client struct {
}

func New() Client {
	client := &client{}
	return client
}

type Client interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

func (c *client) Get() {

}

func (c *client) Post() {

}

func (c *client) Put() {

}

func (c *client) Patch() {

}

func (c *client) Delete() {

}
