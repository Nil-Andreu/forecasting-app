# Forecasting App
In this project, I will create a forecasting app.

We will use **Kafka**, as a low-latency queue streaming processor.
For the backend we will use FastAPI, to create the APIs needed to process the data into the frontend.
For the frontend, we will use the tool of Leaflet.JS, which will help us to create the map.

The steps that I was required to do are the following.

## 1. Set up the KafKa Broker 

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
