package orders

import "github.com/google/uuid"

type Order struct {
	UID  uuid.UUID
	Data []byte
}
