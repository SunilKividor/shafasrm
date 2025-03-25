package ws

import "log"

func (r *Room) run() {
	for msg := range r.Bcast {
		for client := range r.Clients {
			if msg.SenderID == client.UserID {
				continue
			}
			select {
			case client.Send <- msg:
				log.Printf("Sent to %v: %v", client.UserID, msg.Content)
			default:
				log.Printf("client %v slow, dropping msg", client.UserID)
			}
		}
	}
}
