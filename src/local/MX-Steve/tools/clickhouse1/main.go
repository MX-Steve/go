package main

import (
	"database/sql"
	"fmt"
	//clickhouse driver
	_ "github.com/ClickHouse/clickhouse-go"
	log "github.com/sirupsen/logrus"
)

func newClickhouseAgent() {
	dbURI := "tcp://192.168.1.134:9000?database=dbtest"
	conn, err := sql.Open("clickhouse", dbURI)
	if err != nil {
		log.Warn(err)
	}
	log.Info(conn)
	sq1 := `SELECT
		algorithm as algorithm,
		hashrate AS hashrate,
		hashrate AS temperatures,
		cards_num AS fan_speed,
		cards_num as power
	FROM dbtest.fact_bminer_card_stats
	limit 10;`
	rows, err := conn.Query(sq1)
	if err != nil {
		log.Warn(err)
	}
	var (
		algorithm    string
		hashrate     uint64
		temperatures uint64
		fanSpeed     uint64
		power        uint64
	)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&algorithm, &hashrate, &temperatures, &fanSpeed, &power)
		fmt.Printf("%s,%d,%d,%d,%d", algorithm, hashrate, temperatures, fanSpeed, power)
	}
}

func main() {
	newClickhouseAgent()
}
