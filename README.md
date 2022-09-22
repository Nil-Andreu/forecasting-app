# Real-Time Map
In this project, I will create a real-time map.

We will use **Kafka**, as a low-latency queue streaming processor.
For the backend we will use FastAPI, to create the APIs needed to process the data into the frontend.
For the frontend, we will use the tool of Leaflet.JS, which will help us to create the map.

The steps that I was required to do are the following.

## 1. Set up the KafKa Broker 
To run KafKA locally, we will be using Docker.
We will create the kafka setup inside of the messages folder, where inside of it will create the file of *docker-compose.yml*.

In this docker-compose file, we will have the necessary services for KafKa:
- Zookeeper
- Confluentic KafKa

And once we have this, we will up this service: docker-compose up -d. 