package subscriber

import (
	"WB_L0/internal/orders"
	"WB_L0/internal/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
)

type Subscriber struct {
	Conn        stan.Conn
	DurableName string
	Chan        string
	Sub         stan.Subscription
	Service     service.Service
}

func NewSubscriber(conn stan.Conn, a service.Service, durName string) *Subscriber {
	return &Subscriber{
		Conn:        conn,
		Service:     a,
		DurableName: durName,
	}
}
func (s *Subscriber) ProcessMessage(msg *stan.Msg) {

	var order orders.OrderModel
	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		log.Println("error unmarshalling message")
		return
	}
	u, err := uuid.Parse(order.OrderUID)
	if err != nil {
		log.Println("error parsing uuid")
		return
	}
	orderRaw := orders.Order{
		UID:  u,
		Data: msg.Data,
	}
	_, err = s.Service.AddOrder(orderRaw)

	if err != nil {
		log.Printf("could not add order: %v\n", err)
	}
	log.Println("successfully added order message to database")

}

func (s *Subscriber) Subscribe(ch string) {
	s.Chan = ch
	s.Sub, _ = s.Conn.Subscribe(ch, s.ProcessMessage, stan.DurableName(s.DurableName))
}
