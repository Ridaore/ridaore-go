package ridaore

import (
	"fmt"
	"net"
)

type Options struct{
	Port uint16;
	Host string;
}

type Client struct {
	option 		*Options
	connection 	net.Conn
}

func New(opt *Options) Client {
	return Client{option: opt};
}

func (client *Client) Dial() *RidaoreError {
	if client.option == nil {
		return &RidaoreError{
			Message: "Client Option is null",
			Fix: "Adding option via New(opt *Options) function",
		}
	}

	address := fmt.Sprintf("%s:%d", client.option.Host, client.option.Port);

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return &RidaoreError{
			Message: fmt.Sprintf("Error while dialing to %s", address),
			Fix: fmt.Sprintf("Make sure %s is accessible", address),
		}
	}

	client.connection = conn;
	return nil;
}

func (client *Client) Set(key string, value string) *RidaoreError {
	msg := fmt.Sprintf("set %s %s", key, value);
	_, err := client.connection.Write([]byte(msg));
	if err != nil {
		return &RidaoreError{
			Message: "Failed to write into server",
			Fix: err.Error(),
		}
	}

	resp := make([]byte, 1024);
	_, err = client.connection.Read(resp)
	if err != nil {
		return &RidaoreError{
			Message: "Failed to read",
			Fix: err.Error(),
		}
	}

	return nil;
}

func (client *Client) Get(key string) (string, *RidaoreError) {
	msg := fmt.Sprintf("get %s", key);
	_, err := client.connection.Write([]byte(msg));
	if err != nil {
		return "", &RidaoreError{
			Message: "Failed to write into server",
			Fix: err.Error(),
		}
	}

	resp := make([]byte, 1024);
	_, err = client.connection.Read(resp);
	if err != nil {
		return "", &RidaoreError{
			Message: "Failed to read server",
			Fix: err.Error(),
		}
	}
	
	return string(resp), nil;
}