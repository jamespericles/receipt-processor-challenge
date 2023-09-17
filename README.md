# Receipt Processor Challenge

## Getting Started

This project is built using TypeScript and Node.js. To run this project, you will need to have Node.js installed on your machine. You can download Node.js [here](https://nodejs.org/en/download/). You will also need to have Docker installed on your machine, if you'd prefer not to run locally. You can download Docker [here](https://www.docker.com/products/docker-desktop).

## Running the Project

There are a couple ways to run this project, either locally using Node.js or using Docker.

### Running Locally

Assuming you have Node.js installed, you will also need to have either npm or yarn installed. You can download npm [here](https://www.npmjs.com/get-npm) and yarn [here](https://classic.yarnpkg.com/en/docs/install/#mac-stable). Once you have npm or yarn installed, you can run the following commands to run the project locally:

```bash
# Install dependencies
yarn

# Run the project
yarn start
```

If everything is working correctly, you should see 'Server listening at http://localhost:3000' in your terminal.

### Running with Docker

With Docker already installed, you can run the following commands to run the project:

```bash
# Build the Docker image
docker build -t receipt-processor .

# Run the Docker image
docker run -p 3000:3000 receipt-processor
```

Similar to running locally, you should see 'Server listening at http://localhost:3000' in your terminal.

## Interacting with the API

The API includes two endpoints:

- `POST /receipts/process`
- `GET /receipts/{id}/points`

### POST /receipts/process

This endpoint takes in a JSON receipt and returns a JSON object with an ID generated by the code. The ID returned is the ID that should be passed into `GET /receipts/{id}/points` to get the number of points the receipt was awarded.

Here is a curl request you can make to test this endpoint:

```bash
curl -X POST \
  http://localhost:3000/receipts \
  -H 'Content-Type: application/json' \
  -d '{
    "retailer": "M&M Corner Market",
    "purchaseDate": "2022-03-20",
    "purchaseTime": "14:33",
    "items": [
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },{
        "shortDescription": "Gatorade",
        "price": "2.25"
      },{
        "shortDescription": "Gatorade",
        "price": "2.25"
      },{
        "shortDescription": "Gatorade",
        "price": "2.25"
      }
    ],
    "total": "9.00"
  }'
```

This should return a response similar to the following:

```json
{
  "id": "3f1cc7c3-bce6-430d-98d1-eaab1b78de4c"
}
```

### GET /receipts/{id}/points

This endpoint takes in a receipt ID and returns a JSON object containing the number of points awarded.

Here is a curl request you can make to test this endpoint:

```bash
curl -X GET http://localhost:3000/receipts/3f1cc7c3-bce6-430d-98d1-eaab1b78de4c/points
```

This should return a response similar to the following:

```json
{
  "points": 109
}
```
