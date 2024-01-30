package room

type Room struct {
	maxOccupants int
	musicians    map[*Musician]bool
	broadcast    chan []byte
	register     chan *Musician
	unregister   chan *Musician
}

func NewRoom() *Room {
	return &Room{
		maxOccupants: 2,
		musicians:    make(map[*Musician]bool),
		broadcast:    make(chan []byte),
		register:     make(chan *Musician),
		unregister:   make(chan *Musician),
	}
}

func (r *Room) Run() {
	for {
		select {
		case musician := <-r.register:
			// TODO: check for max in room
			r.musicians[musician] = true
		case musician := <-r.unregister:
			if _, ok := r.musicians[musician]; ok {
				close(musician.send)
				delete(r.musicians, musician)
			}
		case message := <-r.broadcast:
			for musician := range r.musicians {
				select {
				case musician.send <- message:
				default:
					close(musician.send)
					delete(r.musicians, musician)
				}
			}
		}
	}
}
