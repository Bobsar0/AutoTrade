package mock

//client directly interacts with the trading site's API
type client struct{

}

//Spins out a new client
func newClient() *client{
	return &client{

	}
}