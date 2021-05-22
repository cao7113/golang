package server

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func (s *ServerTestSuite) TestStatusCode() {
	reason := "testing"
	err := status.Errorf(codes.FailedPrecondition, "failed precondition %s", reason)

	err1, ok := status.FromError(err)
	s.True(ok)
	s.Equal(codes.FailedPrecondition, err1.Code())
	// "msg: failed precondition testing, code: FailedPrecondition, str: FailedPrecondition"
	logrus.Infof("msg: %s, code: %v, str: %s", err1.Message(), err1.Code(), err1.Code().String())
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
}

// The SetupSuite method will be run before any tests are run.
func (s *ServerTestSuite) SetupSuite() {
	go StartRPCServer()
	time.Sleep(1 * time.Second)
}

// The TearDownSuite method will be run after all tests have been run.
func (s *ServerTestSuite) TearDownSuite() {
}
