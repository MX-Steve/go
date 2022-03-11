package datastruct

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "local/MX-Steve/projects/grpc-demo/data-struct/proto/server/djentry"
	"os"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/influxdata/toml"
	"github.com/influxdata/toml/ast"
)

type clickhouseAgent struct {
	conn *sql.DB
}

func newClickhouseAgent(table *ast.Table) (*clickhouseAgent, error) {
	var conf struct {
		Host     string `toml:"host"`
		Port     int64  `toml:"port"`
		User     string `toml:"user"`
		Database string `toml:"database"`
	}
	tb, ok := table.Fields["clickhouse"].(*ast.Table)
	if !ok {
		return nil, errors.New("didn't find configration for clickhouse")
	}
	if err := toml.UnmarshalTable(tb, &conf); err != nil {
		return nil, err
	}

	dbURI := fmt.Sprintf("tcp://%s:%d?username=%s&database=%s&password=%s",
		conf.Host, conf.Port, conf.User, conf.Database, os.Getenv("CLOVER_CLICKHOUSE_PASSWORD"))

	conn, err := sql.Open("clickhouse", dbURI)
	if err != nil {
		return nil, err
	}

	return &clickhouseAgent{
		conn: conn,
	}, nil
}

// get user overview
func (c *clickhouseAgent) GetUserOverview(
	ctx context.Context, in *pb.UserInfoRequest,
) (*pb.UserInfoResponse, error) {
	res := userInfoMock()
	return res, nil
}
