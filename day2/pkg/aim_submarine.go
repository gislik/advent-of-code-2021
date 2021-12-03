package pkg

type AimSubmarine struct {
	X, Y, Aim int
}

func (sm *AimSubmarine) Move(cmd Command, n int) {
	switch cmd {
	case Up:
		sm.Aim -= n
	case Down:
		sm.Aim += n
	case Forward:
		sm.X += n
		sm.Y += sm.Aim * n
	}
}
