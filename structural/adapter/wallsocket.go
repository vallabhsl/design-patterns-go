package adapter

type WallSocket interface {
	getVolts() int
}

type WallSocketAdapter struct {
}

func (adapter WallSocketAdapter) getVolts() int {
	return 240
}
