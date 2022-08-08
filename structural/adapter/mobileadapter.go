package adapter

type MobileAdapter interface {
	get30Volts() int
}

type MobileAdapterImpl struct {
	wallSocket *WallSocketAdapter
}

func (mobileAdapter MobileAdapterImpl) get30Volts() int {
	return mobileAdapter.wallSocket.getVolts() / 8
}

type Mobile struct {
	mobileAdapter MobileAdapter
}

func (mobile Mobile) getVolts() int {
	return mobile.mobileAdapter.get30Volts()
}
