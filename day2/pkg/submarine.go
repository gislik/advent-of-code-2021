package pkg

type Submarine struct {
	X, Y int
}

func (sm *Submarine) Move(cmd Command, n int) {
	switch cmd {
	case Up:
		sm.Y -= n
	case Down:
		sm.Y += n
	case Forward:
		sm.X += n
	}
}
