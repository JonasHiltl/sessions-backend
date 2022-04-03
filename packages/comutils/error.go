package comutils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return status.Error(st.Code(), st.Message())
}
