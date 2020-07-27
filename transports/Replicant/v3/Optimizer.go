package replicant

import (
	pt "github.com/OperatorFoundation/shapeshifter-ipc/v2"
	"github.com/op/go-logging"
	"golang.org/x/net/proxy"
	"net"
)

// This makes Replicant compliant with Optimizer
type TransportClient struct {
	Config  ClientConfig
	Address string
	Dialer  proxy.Dialer
	Log     *logging.Logger
}

type TransportServer struct {
	Config ServerConfig
	Address string
	Dialer  proxy.Dialer
	Log     *logging.Logger
}

func NewClient(config ClientConfig, address string, dialer proxy.Dialer, log *logging.Logger) TransportClient {
	return TransportClient{
		Config: config,
		Address: address,
		Dialer: dialer,
		Log:    log,
	}
}

func NewServer(config ServerConfig, address string, dialer proxy.Dialer, log *logging.Logger) TransportServer {
	return TransportServer{
		Config: config,
		Address: address,
		Dialer: dialer,
		Log:	log,
	}
}

func (transport TransportClient) Dial() (net.Conn, error) {
	conn, dialErr := transport.Dialer.Dial("tcp", transport.Address)
	if dialErr != nil {
		return nil, dialErr
	}

	dialConn := conn
	transportConn, err := NewClientConnection(conn, transport.Config)
	if err != nil {
		_ = dialConn.Close()
		return nil, err
	}

	return transportConn, nil
	}

func (transport TransportServer) Listen() (net.Listener, error) {
	addr, resolveErr := pt.ResolveAddr(transport.Address)
	if resolveErr != nil {
		return nil, resolveErr
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return newReplicantTransportListener(ln, transport.Config), nil
}
	//replicantTransport := New(transport.Config, transport.Dialer)
	//conn := replicantTransport.Dial(transport.Address)
	//conn, err:= replicantTransport.Dial(transport.Address), errors.New("connection failed")
	//if err != nil {
	//	return nil, err
	//} else {
	//	return conn, nil
	//}
	//return conn, nil


