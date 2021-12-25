package model

import (
	"math/rand"
	"rating/config"
	"testing"
	"time"
)

const (
	testUser = "800000000000000000000002"
)

func generateStr() string {
	rand.Seed(time.Now().Unix())
	s := ""
	for i := 0; i < 23; i++ {
		s += string(rune('0' + rand.Intn(10)))
	}

	return "9" + s
}

func generateScore() float64 {
	rand.Seed(time.Now().Unix())
	return float64(rand.Intn(6))
}

func TestMain(m *testing.M) {
	if err := config.InitConfig(config.CfgFileNested); err != nil {
		panic(err)
	}
	if err := InitModel(); err != nil {
		panic(err)
	}
	if code := m.Run(); code != 0 {
		panic(code)
	}
}
