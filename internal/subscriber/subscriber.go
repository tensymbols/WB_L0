package subscriber

import (
	"WB_L0/internal/app"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	Conn stan.Conn
	Chan string
	Sub  stan.Subscription
	App  app.App
}

func NewSubscriber(conn stan.Conn, a app.App) *Subscriber {
	return &Subscriber{
		Conn: conn,
		App:  a,
	}
}
func (s *Subscriber) Subscribe(ch string, goCh chan any) {
	s.Chan = ch
	s.Sub, _ = s.Conn.Subscribe(ch, func(msg *stan.Msg) {
		//o, err := s.App.AddOrder(msg.Data)
		//fmt.Println(string(o), err)
		_, _ = s.App.AddOrder(msg.Data)
		goCh <- true
	}, stan.DeliverAllAvailable())
}
