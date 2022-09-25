from confluent_kafka import Consumer

from forecasting_python.kafka.env import KAFKA_GROUP_ID
from forecasting_python.kafka.utils import error_cb, kafka_config


# Check the following Documentation: https://github.com/confluentinc/confluent-kafka-python/blob/master/examples/confluent_cloud.py
consumer = Consumer({
    **kafka_config,
    'group.id': KAFKA_GROUP_ID,  # consumers are organized in groups, so they do not repeat reading same message
    'auto.offset.reset': 'earliest'
})

consumer.subscribe(['financials'])
    

try:
    while True:
        # Receive all the messages created each 500 milliseconds
        message = consumer.poll(0.5)

        if message is None:
            continue

        if message.error():
            print('Consumer error: {}'.format(message.error()))
            continue

        value = message.value()
        print('Consumed: {}'.format(value))

except KeyboardInterrupt:
    pass

finally:
    # When we press the KeyBoard interrupt, we close the connection of KafKa
    consumer.close()



