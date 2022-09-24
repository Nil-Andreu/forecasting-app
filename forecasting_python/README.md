# Forecasting
We will make the forecasts of the companies that we need.
For this, we will make periodically requests to obtain financial information of the public companies.

Based on this information, we will run our model to make the predictions that we want to know:
- Amount of predicted sales
- Prediction of public price
- If the price on the next period will increment/decrement

For this, we are also going to obtain metrics on how well our algorithm is working.

## KafKa
Once we have all those predictions, we will publish them on KafKa, so they can be read by the main application made in Golang.

This way, we aim to provide advanced analytics on the public companies.