from confluent_kafka import Consumer

from forecasting_python.kafka.client.utils import error_cb
from forecasting_python.kafka.client.env import (
    KAFKA_API_KEY, 
    KAFKA_GROUP_ID, 
    KAFKA_SECRET_API_KEY, 
    BOOTSTRAP_SERVER, 
    SECURITY_PROTOCOL, 
    SASL_MECHANISM
)

# Check the following Documentation: https://github.com/confluentinc/confluent-kafka-python/blob/master/examples/confluent_cloud.py
consumer = Consumer({
    'sasl.username': KAFKA_API_KEY,
    'sasl.password': KAFKA_SECRET_API_KEY,
    'bootstrap.servers': BOOTSTRAP_SERVER,
    'security.protocol': SECURITY_PROTOCOL,
    'sasl.mechanism': SASL_MECHANISM,
    'group.id': KAFKA_GROUP_ID,  # consumers are organized in groups, so they do not repeat reading same message
    'auto.offset.reset': 'earliest',
    'error_cb': error_cb,
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



