package publisher

import "github.com/nats-io/stan.go"

type Publisher struct {
	Conn stan.Conn
	Chan string
}

func NewPublisher(conn stan.Conn, chanName string) *Publisher {
	return &Publisher{
		Conn: conn,
		Chan: chanName,
	}
}

func (p *Publisher) Publish(ch string, data []byte) error {
	return p.Conn.Publish(ch, data)
}
