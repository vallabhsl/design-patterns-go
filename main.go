package main

import (
	"fmt"

	"github.com/vallabhsl/design-patterns-go/adapter"
)

func main() {
	fmt.Println("Design Patterns")
	fmt.Println("1. Structural: Adapter Pattern")

	wallSocket := adapter.WallSocketAdapter{}
	mobileAdapter := adapter.MobileAdapterImpl{
		wallSocket: &wallSocket,
	}
	mobile := adapter.Mobile{
		mobileAdapter: mobileAdapter,
	}

}
