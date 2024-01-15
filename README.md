## Overview

A self-hosted REST API to convert currencies and retrieve current exchange rates, leveraging Open Exchange Rates.

## Features

- Avoid the 1,000 requests/month limit of Open Exchange Rates.
- Host the API on your own server for full control and privacy.
- Built in Go for speed, the API handles thousands of requests per second with ease.

## API Endpoints

1. **Convert Currency**

   - **Endpoint:** `GET /convert?from=GBP&to=USD&amount=100`
   - **Response:** `{"amount": 127}`
   - **Description:** Converts a specified amount from one currency to another.

2. **Get Conversion Rate**

   - **Endpoint:** `GET /rate?from=GBP&to=USD`
   - **Response:** `{"rate": 1.27}`
   - **Description:** Retrieves the current exchange rate from one currency to another.

3. **Get All Rates for a Base Currency**
   - **Endpoint:** `GET /rates?base=GBP`
   - **Response:** `{"base": "GBP", "rates": {"USD": 1.27}}`
   - **Description:** Provides exchange rates for all available currencies against a base currency. The base is optional and defaults to USD.

## Installation

The API is available as a Docker image on GitHub Container Registry. To run the API, you can use the following command:

```bash
docker run -p 8080:8080 -e SERVER_PORT=8080 -e OPEN_EXCHANGE_RATES_APP_ID=<your_app_id> ghcr.io/ClyentSoftwares/currency-api:latest
```

## Configuration

| Environment Variable       | Required | Description                                                                                                             |
| -------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------- |
| SERVER_PORT                | No       | The port on which the API server will listen for requests. Defaults to 8080.                                            |
| OPEN_EXCHANGE_RATES_APP_ID | Yes      | The App ID for your Open Exchange Rates account. You can get one for free at https://openexchangerates.org/signup/free. |
