package tools

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"runtime"
	"testing"
	"time"
)

func (s *CpuToolsSuite) TestDead() {
	logrus.Infof("cpu cores: %d", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go doTask(1)
	go doTask(2)
	go doTask(3)
	go doTask(4)
	select {}
}

func doTask(i int) {
	for {
		logrus.Infof("[%d] run at: %s", i, time.Now())
		time.Sleep(1 * time.Nanosecond)
	}
}

func TestCpuToolsSuite(t *testing.T) {
	suite.Run(t, &CpuToolsSuite{})
}

type CpuToolsSuite struct {
	suite.Suite
}

func (s *CpuToolsSuite) SetupSuite() {
	//
}

func (s *CpuToolsSuite) SetupTest() {
}

// The TearDownSuite method will be run after all tests have been run.
func (s *CpuToolsSuite) TearDownSuite() {
}
