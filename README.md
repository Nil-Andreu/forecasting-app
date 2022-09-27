# Forecasting App
In this project, I will create a forecasting app.

We will use **Kafka**, as a low-latency queue streaming processor.

For the backend we will use **Golang**, to create the APIs needed to process the data into the frontend (made with **React.ts**).

And we will have microservices to injest/clean new data, as well as making applying machine learning to make forecasts. Those microservices will be in **FastAPI**, so we can take advantage of async.

![Functionality Diagram](forecasting-app-diagram.png)

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


## 2. Data Injestion

There will be no product if we do not have the data resources.

For this, we will go inside of the **injestion** folder, where we will develop our application for consuming data, and inside of the notebooks we will start injesting and processing this data.

# Bibliography

I have read different books to be able to develop this application.

I have divided the references by the thematic, so if anyone is interested could consult those books:

- **Forecasting/Statistics/Finance**:
    - Jiri Pik (2021) - Hands-On Financial Trading with Python. A Practical guide to using Zipline and other Python libraries for Backtesting Trading Strategies
    - Stefan Jansen (2020) - Machine Learning for Algorithmic Trading: Predictive models to extract signals from market and alternative data for systematic trading strategies with Python
    - Yves Hilpisch (2020) - Python for Algorithmic Trading: From Idea to Cloud Deployment

