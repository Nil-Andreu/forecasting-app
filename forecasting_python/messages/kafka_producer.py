from typing import List

from confluent_kafka import Producer

from .env import (
    KAFKA_API_KEY, 
    KAFKA_GROUP_ID, 
    KAFKA_PRODUCER_ID,
    KAFKA_SECRET_API_KEY, 
    BOOTSTRAP_SERVER, 
    SECURITY_PROTOCOL, 
    SASL_MECHANISM
)


p = Producer({'bootstrap.servers': 'mybroker1,mybroker2'})


def delivery_report(err, msg):
    """ 
    Called once for each message produced to indicate delivery result.
    Triggered by poll() or flush(). 
    """

    if err is not None:
        print('Message delivery failed: {}'.format(err))

    else:
        print('Message delivered to {} [{}]'.format(msg.topic(), msg.partition()))


for data in []:
    # Trigger any available delivery report callbacks from previous produce() calls
    p.poll(1)

    # Asynchronously produce a message. The delivery report callback will
    # be triggered from the call to poll() above, or flush() below, when the
    # message has been successfully delivered or failed permanently.
    p.produce('map_insight', data.encode('utf-8'), callback=delivery_report)