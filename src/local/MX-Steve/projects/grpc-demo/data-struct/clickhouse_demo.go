package datastruct

import (
	"fmt"
	pb "local/MX-Steve/projects/grpc-demo/data-struct/proto/server/djentry"
	"time"
)

func userInfoMock() *pb.UserInfoResponse {
	nowTime := time.Now().Unix()
	res := &pb.UserInfoResponse{
		Msg: fmt.Sprintf("%d-%s", nowTime, "lisi"),
	}
	return res
}
