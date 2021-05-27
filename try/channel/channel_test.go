package channel

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

func (s *ChannelSuite) TestChan() {
	s.Run("declare and select", func() {
		var ch chan int
		s.Nil(ch)

		var i int
		select {
		case i = <-ch: //The receive operation might block a goroutine because of the 'nil' channel
			log.Println("read i=", i)
		default:
			i = 999
		}
		s.Equal(999, i)
	})

	// Note that it is only necessary to close a channel if the receiver is looking for a close.
	// Closing the channel is a control signal on the channel indicating that no more data follows.
	s.Run("declare and close", func() {
		var ch chan int
		//close(ch) // panic: close of nil channel [recovered]
		ch = make(chan int)
		close(ch)
		s.NotNil(ch)
		//close(ch) // panic: close of closed channel
		j := <-ch // read from closed channel, non-block and get zero value
		s.Equal(0, j)

		i, ok := <-ch // read from closed channel and check whether closed
		s.False(ok)
		s.Equal(0, i)

		for k := range ch {
			log.Fatalf("never run here k = %d ", k)
		}

		log.Println("end")
	})

	//s.Run("close and select", func() {
	//	ch := make(chan int)
	//	close(ch)
	//	ch2 := make(chan int)
	//	close(ch2)
	//	i := 0
	//	for {
	//		select {
	//		case <-ch:
	//			log.Println("read from closed channel", i) // rand run this
	//		case <-ch2:
	//			log.Println("read from closed channel2", i) // rand run this
	//		default:
	//			log.Println("default in closed channel select", i) // never run here
	//		}
	//		i++
	//	}
	//})

	s.Run("make and read", func() {
		ch := make(chan int)
		defer close(ch)

		var i int
		select {
		case i = <-ch:
			log.Println("read i=", i)
		default:
			i = 999
		}
		s.Equal(999, i)
	})

	s.Run("write and read", func() {
		ch := make(chan int)
		go func(chan int) {
			defer close(ch)
			for i := 0; i < 3; i++ {
				ch <- i
			}
		}(ch)

		// for range是阻塞式读取channel，只有channel close之后才会结束
		for i := range ch { // will block here if not close ch
			log.Println("i =", i)
		}
		log.Println("end")
	})
}

func (s *ChannelSuite) TestChanBuf() {
	s.Run("has buf", func() {
		ch := make(chan int, 1)
		ch <- 123
		i := <-ch
		s.Equal(123, i)
		//<-ch // will block
	})
}

func TestChannelSuite(t *testing.T) {
	suite.Run(t, &ChannelSuite{})
}

type ChannelSuite struct {
	suite.Suite
}
