from operator import ne
from kafka import KafkaConsumer, TopicPartition
import datetime
import time
from func_timeout import func_set_timeout


@func_set_timeout(1)
def get_data(bootstrap_servers, topic, group_id, key, offset):
    consumer = KafkaConsumer(bootstrap_servers=bootstrap_servers,
                             group_id=group_id,
                             auto_offset_reset='earliest')
    consumer.assign([TopicPartition(topic=topic, partition=0)])
    msgList = []
    consumer.seek_to_end(TopicPartition(topic=topic, partition=0))
    last_offset = consumer.position(TopicPartition(topic=topic, partition=0))
    consumer.seek(TopicPartition(topic=topic, partition=0), offset)
    for msg in consumer:
        if key == msg.key.decode('utf-8'):
            msgList.append(msg.value.decode('utf-8'))
            if msg.offset == last_offset - 1:
                break
    return msgList, last_offset


if __name__ == "__main__":
    result = get_data("121.43.41.139:8090", "revolution", 'te2st6', 'test', 39)
    print(result)
