package listener

import "net"

// Listener interface defines the Listen method for network connections.
type Listener interface {
	Listen(network, address string) (net.Listener, error)
}

type DefaultListener struct{}

func (l *DefaultListener) Listen(network, address string) (net.Listener, error) {
	return net.Listen(network, address)
}
