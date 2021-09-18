package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
)

func ex1() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Do("Auth", "default", "kaituozhe")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("Set", "abc", 100)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("Set", "a2", 200)
	if err != nil {
		log.Fatal(err)
	}
	r, err := redis.Ints(conn.Do("MGet", "abc", "a2"))
	if err != nil {
		log.Fatal(err)
	}
	log.Info(r)
	for _, v := range r {
		log.Info(v)
	}
}
func ex2List() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Do("Auth", "default", "kaituozhe")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("lpush", "book_list", "hamlt", "sunkong", 300)
	if err != nil {
		log.Fatal(err)
	}
	r, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		log.Fatal(err)
	}
	log.Info(r)
	r, err = redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		log.Fatal(err)
	}
	log.Info(r)

}
func ex3Hash() {
	conn, err := redis.Dial("tcp", "172.27.42.232:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Do("Auth", "default", "kaituozhe")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("HSet", "books", "abc", 120)
	if err != nil {
		log.Fatal(err)
	}
	r, err := redis.Int(conn.Do("HGet", "books", "abc"))
	if err != nil {
		log.Fatal(err)
	}
	log.Info(r)
}
func ex4Expire() {
	conn, err := redis.Dial("tcp", "172.27.42.232:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Do("Auth", "default", "kaituozhe")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("expire", "abc", 10)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 11)
	r, err := redis.Int(conn.Do("get", "abc"))
	if err != nil {
		log.Warn(err)
	}
	log.Info(r)
}

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func ex5Pool() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "a", "abc")
	if err != nil {
		log.Fatal(err)
	}
	r, err := redis.String(conn.Do("Get", "a"))
	if err != nil {
		log.Warn(err)
	}
	log.Info(r)
}

// connecting redis db
func main() {
	testIt := true
	if testIt {
		ex5Pool()
	} else {
		ex1()
		ex2List()
		ex3Hash()
		ex4Expire()
	}
}
