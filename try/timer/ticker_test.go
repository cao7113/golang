package timer

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

func (s *TickerSuite) TestTick() {
	s.Run("Tick", func() {
		ch := time.After(3 * time.Second)
		tk := time.Tick(1 * time.Second)
		for {
			stop := false
			select {
			case tm := <-tk:
				log.Println(tm)
			case <-ch:
				stop = true
				log.Println("stop at ", time.Now())
			}
			if stop {
				break
			}
		}
	})

	s.Run("NewTicker", func() {
		tk := time.NewTicker(1 * time.Second)
		timeout := time.After(3 * time.Second)

		go func() {
			<-time.After(2 * time.Second)
			tk.Stop()
		}()

		for {
			stop := false
			select {
			case tm := <-tk.C:
				log.Println(tm)
			case <-timeout: // Note： 从last tick 计算还需3s
				stop = true
				log.Println("stop at ", time.Now())
			}
			if stop {
				break
			}
		}
	})
}

func TestTickerSuite(t *testing.T) {
	suite.Run(t, &TickerSuite{})
}

type TickerSuite struct {
	suite.Suite
}
