# Forecasting App
In this project, I will create a forecasting app.

We will use **Kafka**, as a low-latency queue streaming processor.
For the backend we will use FastAPI, to create the APIs needed to process the data into the frontend.
For the frontend, we will use the tool of Leaflet.JS, which will help us to create the map.

The steps that I was required to do are the following.

## 1. Set up the KafKa Broker 
### Locally
To run KafKa locally, we will be using Docker.
We will create the kafka setup inside of the messages folder, where inside of it will create the file of *docker-compose.yml*.

In this docker-compose file, we will have the necessary services for KafKa:
- Zookeeper
- Confluentic KafKa

And once we have this, we will up this service: *docker-compose up -d*. 
Instead of using docker compose, we could create it by parts:
```
    docker run --name=zookeeper -d -e ZOOKEEPER_CLIENT_PORT=2181 -p 2181:2181 -p 2888:2888 -p 3888:3888 confluentinc/cp-zookeeper:latest

    Zookeeper_Server_IP=$(docker inspect zookeeper --format='{{ .NetworkSettings.IPAddress }}')

    docker run --name=kafka -e KAFKA_ZOOKEEPER_CONNECT=${Zookeeper_Server_IP}:2181 -e KAFKA_LISTENERS=PLAINTEXT://localhost:9092 -e KAFKA_ADVERTISED_LISTENERS='PLAINTEXT://broker:9092' -d -p 9092:9092 confluentinc/cp-kafka:latest

```
This is a good solution in the case our computer has *trouble with initializing zookeeper*.

To create a new topic, we would have to run: 
```
    docker compose exec broker \
        kafka-topics --create \
            --topic map_insight \
            --bootstrap-server localhost:9092 \
            --replication-factor 1 \
            --partitions 1
```
Where we would have created the new topic: map_insight.

We would also create the *getting_started.ini* settings, where will be used when starting the kafka.

### Cloud
We will use a free plan of a cloud provider of the kafka services: confluent.

For using KafKa with Python we are going to use the package of *python-confluent*.
For us to connect to it, we would need the necessary credentials (in .env.example).

Then, we would create the Consumer and Producer with Python:
- Consumer:
    - Bootstrap Server
    - SASL mechanism
    - Security Protocol
    - Username
    - Password
    - Group Id
    - KafKa Consumer Id
    - Offset
    - Max. Polling interval

- Producer:
    - Bootstrap Server
    - SASL mechanism
    - Security Protocol
    - Username
    - Password
    - Group Id
    - KafKa Producer Id

So as we can see, some part of the configuration is similar. 
