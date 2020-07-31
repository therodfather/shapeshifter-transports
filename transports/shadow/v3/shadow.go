/*
	MIT License

	Copyright (c) 2020 Operator Foundation

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

// Package shadow provides a PT 2.1 Go API wrapper around the connections used by Shadowsocks
package shadow

import (
	"github.com/kataras/golog"
	"log"
	"net"
	"os"

	shadowsocks "github.com/shadowsocks/go-shadowsocks2/core"
)

func MakeLog() {
	golog.SetLevel("debug")
	golog.SetOutput(os.Stderr)
}

//Config contains the necessary command like arguments to run shadow
type ClientConfig struct {
	Password   string `json:"password"`
	CipherName string `json:"cipherName"`
	Address    string `json:"address"`
}

type ServerConfig struct {
	Password   string `json:"password"`
	CipherName string `json:"cipherName"`
	Address    string `json:"address"`
}

//Transport contains the arguments to be used with Optimizer
type Transport struct {
	Password   string
	CipherName string
	Address    string
}

//NewConfig is used to create a config for testing
func NewClientConfig(password string, cipherName string, address string) ClientConfig {
	return ClientConfig{
		Password:   password,
		CipherName: cipherName,
		Address:    address,
	}
}

func NewServerConfig(password string, cipherName string) ServerConfig {
	return ServerConfig{
		Password:   password,
		CipherName: cipherName,
	}
}

//NewTransport is used for creating a transport for Optimizer
func NewTransport(password string, cipherName string, address string) Transport {
	return Transport{
		Password:   password,
		CipherName: cipherName,
		Address:    address,
	}
}

//Listen checks for a working connection
func (config ServerConfig) Listen(address string) (net.Listener, error) {
	cipher, err := shadowsocks.PickCipher(config.CipherName, nil, config.Password)
	if err != nil {
		golog.Errorf("Failed generating ciphers:", err)
		return nil, err
	}

	listener, listenerErr := shadowsocks.Listen("tcp", address, cipher)
	if listenerErr != nil {
		log.Fatal("Failed to start listener:", listenerErr)
		return nil, listenerErr
	}
	return listener, nil
}

//Dial connects to the address on the named network
func (config ClientConfig) Dial(address string) (net.Conn, error) {
	cipher, err := shadowsocks.PickCipher(config.CipherName, nil, config.Password)
	if err != nil {
		log.Fatal("Failed generating ciphers:", err)
	}

	conn, err := shadowsocks.Dial("tcp", address, cipher)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Dial creates outgoing transport connection
func (transport *Transport) Dial() (net.Conn, error) {
	cipher, err := shadowsocks.PickCipher(transport.CipherName, nil, transport.Password)
	if err != nil {
		log.Fatal("Failed generating ciphers:", err)
	}

	return shadowsocks.Dial("tcp", transport.Address, cipher)
}

func (transport *Transport) Listen() (net.Listener, error) {
	cipher, err := shadowsocks.PickCipher(transport.CipherName, nil, transport.Password)
	if err != nil {
		log.Fatal("Failed generating ciphers:", err)
		return nil, err
	}

	listener, listenerErr := shadowsocks.Listen("tcp", transport.Address, cipher)
	if listenerErr != nil {
		log.Fatal("Failed to start listener:", listenerErr)
		return nil, listenerErr
	}
	return listener, nil
}
