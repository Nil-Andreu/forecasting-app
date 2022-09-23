
import json
from typing import List
from confluent_kafka import Producer


p = Producer({'bootstrap.servers': 'mybroker1,mybroker2'})


def get_coordinates(path: str) -> List:
    """
    This function will return the coordinates stored in a *path* of a JSON file.
    """

    file = open("../data/route.json")
    json_array = json.load(file)
    coordinates_array = json_array["features"][0]["geometry"]["coordinates"]

    return coordinates_array



def delivery_report(err, msg):
    """ 
    Called once for each message produced to indicate delivery result.
    Triggered by poll() or flush(). 
    """

    if err is not None:
        print('Message delivery failed: {}'.format(err))

    else:
        print('Message delivered to {} [{}]'.format(msg.topic(), msg.partition()))


coordinates = get_coordinates()
for data in coordinates:
    # Trigger any available delivery report callbacks from previous produce() calls
    p.poll(1)

    # Asynchronously produce a message. The delivery report callback will
    # be triggered from the call to poll() above, or flush() below, when the
    # message has been successfully delivered or failed permanently.
    p.produce('map_insight', data.encode('utf-8'), callback=delivery_report)