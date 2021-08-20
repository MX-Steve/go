package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	topic := os.Args[1]
	testHost := os.Args[2]
	consumer, err := sarama.NewConsumer([]string{"34.222.46.168:30530"}, nil)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return
	}
	fmt.Println("connnect success...")
	defer consumer.Close()
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("geet partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, p, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				if strings.Contains(string(m.Value), testHost) {
					fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
