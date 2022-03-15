package datastruct

import (
	"context"
	pb "local/MX-Steve/projects/grpc-demo/data-struct/proto/server/djentry"
)

type StatisticsServer struct {
	clickhouse *clickhouseAgent
	pb.UnimplementedDjentryStatisticsServer
}

// NewStatisticsServer create an instance for database query.
func NewStatisticsServer(config string) (*StatisticsServer, error) {
	table, err := parseConfig(config)
	if err != nil {
		return nil, err
	}

	chAgent, err := newClickhouseAgent(table)
	if err != nil {
		return nil, err
	}
	return &StatisticsServer{
		clickhouse: chAgent,
	}, nil
}

func (s *StatisticsServer) GetUserInfo(
	ctx context.Context, in *pb.UserInfoRequest,
) (*pb.UserInfoResponse, error) {
	return s.clickhouse.GetUserOverview(ctx, in)
}
