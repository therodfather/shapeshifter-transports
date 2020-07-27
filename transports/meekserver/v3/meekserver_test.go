package meekserver

import (
	"github.com/op/go-logging"
	"testing"
)

//if disableTLS is true, it doesnt require the cert and key
//func TestMeekServerListen(t *testing.T) {
//	meekserverTransport := NewMeekTransportServer(true, "", "", "state")
//	listener := meekserverTransport.Listen("127.0.0.1:80")
//	if listener == nil {
//		t.Fail()
//	}
//}
func TestMeekServerListen(t *testing.T) {
	acmeEmail := "brandon@operatorfoundation.org"
	keyFileName := "operatorrss.com"
	meekserverTransport := NewMeekTransportServer(false, acmeEmail, keyFileName, "state")
	if meekserverTransport == nil {
		t.Fail()
		return
	}
	_, listenErr := meekserverTransport.Listen("127.0.0.1:8080")
	if listenErr != nil {
		t.Fail()
		return
	}
}

func TestMeekServerFactoryListen(t *testing.T) {
	var log = logging.MustGetLogger("meekserver")
	acmeEmail := "brandon@operatorfoundation.org"
	acmeHostNames := "operatorrss.com"
	meekserverTransport, newError := New(false, acmeHostNames, acmeEmail,"127.0.0.1:8080", "state", log)
	if newError != nil {
		t.Fail()
		return
	}
	_, listenErr := meekserverTransport.Listen()
	if listenErr != nil {
		t.Fail()
		return
	}
}