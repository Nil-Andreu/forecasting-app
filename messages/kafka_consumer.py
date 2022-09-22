import json
from typing import List

from pykafka import KafkaClient


def get_coordinates(path: str) -> List:
    """
    This function will return the coordinates stored in a *path* of a JSON file.
    """

    file = open("../data/route.json")
    json_array = json.load(file)
    coordinates_array = json_array["features"][0]["geometry"]["coordinates"]

    return coordinates_array


# We instantiate our kafka client that has a host in port 9092 (location of our kafka broker)
# client = KafkaClient(hosts='localhost:9092')
# topic = client.topics['map_insights']  # we can access the topics of our kafka

# And now we create the producer
# producer = topic.get_sync_producer()
# producer.produce("new message")


if __name__ == '__main__':
    res = get_coordinates("../data/route.json")
    print(res)

