package rpc

import (
    "bufio"
    "encoding/gob"
    "net"
    "io"
)

// Dial connects to an RPC server at the specified network address.
func TestDial(network, address string) (*Client, error) {
    conn, err := net.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return NewTestClient(conn), nil
}

// NewClient returns a new Client to handle requests to the
// set of services at the other end of the connection.
// It adds a buffer to the write side of the connection so
// the header and payload are sent as a unit.
func NewTestClient(conn io.ReadWriteCloser) *Client {
    encBuf := bufio.NewWriter(conn)
    client := &Client{
        codec:       &gobClientCodec{conn, gob.NewDecoder(conn), gob.NewEncoder(encBuf), encBuf},
        pending:     make(map[uint64]*Call),
        discallback: make([]func(err error), 0),
    }
    return client
}