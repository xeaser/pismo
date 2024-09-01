# Pismo API

This is a RESTful API designed to manage accounts and transactions. It provides endpoints for creating and retrieving accounts, as well as creating and filtering transactions.

## Endpoints

### Accounts

* **POST /accounts**: Creates a new account.
* **GET /accounts/{accountId}**: Retrieves an account by its ID.

### Transactions

* **POST /transactions**: Creates a new transaction.
* **POST /transactionsByFilter**: Retrieves transactions based on filters such as account ID and operation type.

## Request and Response Formats

All requests and responses are in JSON format.

### Account Model

* **Id**: Unique identifier for the account.
* **Document Number**: The document number for the account.

### Transaction Model

* **TransactionId**: Unique identifier for the transaction.
* **AccountId**: The ID of the account associated with the transaction.
* **OperationType_ID**: The type of operation (e.g., deposit, withdrawal).
* **Amount**: The amount of the transaction.
* **EventDateTimestamp**: The date and time, unix timestamp of the transaction.

### Transaction Filter Model

* **AccountId**: The ID of the account to filter transactions by.
* **OperationType**: The operation type to filter transactions by.

## Error Handling

The API returns standard HTTP error codes and a JSON error response with a message in case of an error.

## Running the Server

To run the server, execute any of the following command:
```
env=local go run ./cmd/main.go
```
```
make run
```
```
docker compose up
```
This will start the server on the port specified in the configuration file.
