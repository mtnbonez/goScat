package goscatclient

import (
	"bufio"
	"os"

	"github.com/google/uuid"
)

type Client interface {

	// Connect checks for
	Connect() bool
	GetInput() string
}

type IPClient struct {

	// ID of the client
	ID uuid.UUID

	// Client is the underlying client struct
	Client Client

	// IPAddress of host
	IPAddress string

	// Port number to connect
	Port string

	// Secret used to connect (if applicable)
	Secret string
}

func (client IPClient) Connect() bool {

	// TODO!

	return false
}

func (client IPClient) GetInput() string {
	return ""
}

type LocalClient struct {

	// Client is the underlying client struct
	Client Client

	Scanner *bufio.Scanner
}

func NewLocalClient() *LocalClient {
	client := LocalClient{
		Scanner: bufio.NewScanner(os.Stdin),
	}
	return &client
}

func (client LocalClient) Connect() bool {

	// Legit just connect, bro.

	return true
}

func (client LocalClient) GetInput() string {
	// Here's where we interface with the UI?
	line := ""

	if client.Scanner.Scan() {
		line = client.Scanner.Text()
	}
	return line
}
