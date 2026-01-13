package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransSnap() snap.Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	env := os.Getenv("MIDTRANS_ENV")

	midtransEnv := midtrans.Sandbox
	if env == "production" {
		midtransEnv = midtrans.Production
	}

	s := snap.Client{}
	s.New(serverKey, midtransEnv)

	return s
}
