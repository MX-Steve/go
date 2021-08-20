package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	feet int
	Animal
}

func (d *Dog) wangwang() {
	fmt.Printf("%s在汪汪汪\n", d.name)
}

// Total is a structure
type Total struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

// structure and json serialization
func main() {
	var d = Dog{
		feet: 4,
		Animal: Animal{
			name: "旺财",
		},
	}
	d.wangwang()
	d.move()
	var t1 = Total{
		ID:     1,
		Gender: "male",
		Name:   "zhang",
	}
	fmt.Println(t1)
	b, err := json.Marshal(t1)
	if err != nil {
		fmt.Println("Json serialization error")
	}
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Printf("%#v\n", string(b))
	// str := "{\"ID\":1,\"Gender\":\"male\",\"Name\":\"zhang\"}"
	str := "{\"id\":1,\"gender\":\"male\",\"name\":\"zhang\"}"
	var t2 = &Total{}
	json.Unmarshal([]byte(str), t2)
	fmt.Println(t2)
	fmt.Printf("%T\n", t2)
}

/*

docker run --rm -d --name=clickhouse-server \
--ulimit nofile=262144:262144 \
-p 18123:8123 -p 19009:9009 -p 19090:9000 \
yandex/clickhouse-server:20.3.5.21

docker cp clickhouse-server:/etc/clickhouse-server/config.xml /app/cloud/clickhouse/conf/config.xml
docker cp clickhouse-server:/etc/clickhouse-server/users.xml /app/cloud/clickhouse/conf/users.xml

PASSWORD=$(base64 < /dev/urandom | head -c8); echo "$PASSWORD"; echo -n "$PASSWORD" | sha256sum | tr -d '-'
PkMst3xF
c498a632827ea3165203d277a88f9cd4f5bd1770d593079bdf4e822a89b54514

PASSWORD=$(base64 < /dev/urandom | head -c8); echo "$PASSWORD"; echo -n "$PASSWORD" | sha256sum | tr -d '-'
tgLjBhL+
3d4d52962cbbffe226a3c91a798d031d7ac969be41821ef1b0f3e7fa239b62f3


docker run -d --name=clickhouse-server --restart always \
-p 18123:8123 -p 19009:9009 -p 19090:9000 \
--ulimit nofile=262144:262144 \
-v /app/cloud/clickhouse/data:/var/lib/clickhouse:rw \
-v /app/cloud/clickhouse/conf/config.xml:/etc/clickhouse-server/config.xml \
-v /app/cloud/clickhouse/conf/users.xml:/etc/clickhouse-server/users.xml \
-v /app/cloud/clickhouse/log:/var/log/clickhouse-server:rw \
yandex/clickhouse-server:20.3.5.21



<root>
    <password_sha256_hex>35542ded44184b1b4b6cd621e052662578025b58b4187176a3ad2b9548c8356e</password_sha256_hex>
	<networks incl="networks" replace="replace">
		<ip>::/0</ip>
	</networks>
	<profile>default</profile>
	<quota>default</quota>
</root>
*/
