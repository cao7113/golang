package channel

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

func (s *ChannelSuite) TestChanBasic() {
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
			log.Fatalf("will not run here k = %d ", k)
		}

		log.Println("end")
	})

	s.Run("close and select", func() {
		ch := make(chan int)
		close(ch)
		ch2 := make(chan int)
		close(ch2)
		i := 0
		for {
			select {
			case <-ch:
				log.Println("read from closed channel", i) // always run this
			case <-ch2:
				log.Println("read from closed channel2", i) // always run this
			default:
				log.Println("default in closed channel select", i)
			}
			i++
		}
	})

	s.Run("make and read", func() {
		ch := make(chan int)
		var i int
		select {
		case i = <-ch: // will block when no data to read
			log.Println("read i=", i)
		default:
			i = 999
		}
		s.Equal(999, i)
		close(ch)
	})

	s.Run("write and read", func() {
		ch := make(chan int)
		go func() {
			for i := 0; i < 3; i++ {
				ch <- i
			}
			close(ch)
		}()
		for i := range ch {
			log.Println("i =", i)
		}
		log.Println("end")
	})
}

func (s *ChannelSuite) TestChan() {
	ch := make(chan int)
	go func(chan int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}(ch)
	for n := range ch {
		log.Println(n)
	}
}

func TestChannelSuite(t *testing.T) {
	suite.Run(t, &ChannelSuite{})
}

type ChannelSuite struct {
	suite.Suite
}
