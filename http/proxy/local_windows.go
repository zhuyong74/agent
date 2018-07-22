// +build windows

package proxy

import (
	"net"
	"net/http"

	"bitbucket.org/portainer/agent"
	"github.com/Microsoft/go-winio"
)

// NewLocalProxy returns a pointer to a LocalProxy.
func NewLocalProxy(clusterService agent.ClusterService) *LocalProxy {
	proxy := &LocalProxy{
		transport: newNamedPipeTransport("//./pipe/docker_engine"),
	}
	return proxy
}

func newNamedPipeTransport(namedPipePath string) *http.Transport {
	return &http.Transport{
		Dial: func(proto, addr string) (conn net.Conn, err error) {
			return winio.DialPipe(namedPipePath, nil)
		},
	}
}
