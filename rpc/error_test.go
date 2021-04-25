package rpc

//func TestStatusCode(t *testing.T) {
//	reason := "testing"
//	err := status.Errorf(codes.FailedPrecondition, "failed precondition %s", reason)
//
//	err1, ok := status.FromError(err)
//	assert.True(t, ok)
//	assert.Equal(t, codes.FailedPrecondition, err1.Code())
//	// "msg: failed precondition testing, code: FailedPrecondition, str: FailedPrecondition"
//	logrus.Infof("msg: %s, code: %v, str: %s", err1.Message(), err1.Code(), err1.Code().String())
//}
