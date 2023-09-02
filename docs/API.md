# PayWalletEngine API Documentation


This is a basic MVP version of a bank account system implemented in Go using the hexagonal architecture. It includes features such as JWT authentication, account management, transaction processing, and user management.

### Base URL

The base URL for all API requests is:

```
http://<your-server-address>:8080
```

## Authentication

All API requests, except for the health and ready checks, require authentication. The API uses JSON Web Tokens (JWT) for authentication. You must include a valid JWT in the `Authorization` header of your request.

## API Endpoints

### Health and Ready Checks

- `GET /alive`
    - Description: Checks if the server is alive.
    - Authentication: No
    - Response: `{"message": "I am Alive!"}`

- `GET /ready`
    - Description: Checks if the server is ready to accept requests.
    - Authentication: No
    - Response: `{"message": "I am Ready!"}`

### Users

- `POST /api/v1/user`
    - Description: Creates a new user.
    - Authentication: Yes
    - Request Body: `User`

- `GET /api/v1/users/{id}`
    - Description: Retrieves a user by their ID.
    - Authentication: Yes
    - Parameters:
        - `id` (integer): The ID of the user.
    - Response: `User`

- `GET /api/v1/users/email/{email}`
    - Description: Retrieves a user by their email.
    - Authentication: Yes
    - Parameters:
        - `email` (string): The email of the user.
    - Response: `User`

- `GET /api/v1/users/username/{username}`
    - Description: Retrieves a user by their username.
    - Authentication: Yes
    - Parameters:
        - `username` (string): The username of the user.
    - Response: `User`

- `PUT /api/v1/users/{id}`
    - Description: Updates a user by their ID.
    - Authentication: Yes
    - Parameters:
        - `id` (integer): The ID of the user.
    - Request Body: `User`

- `DELETE /api/v1/users/{id}`
    - Description: Deletes a user by their ID.
    - Authentication: Yes
    - Parameters:
        - `id` (integer): The ID of the user.

- `GET /api/v1/users/ping`
    - Description: Pings the users service.
    - Authentication: Yes
    - Response: `{"message": "pong"}`

### Accounts

- `POST /api/v1/account`
    - Description: Creates a new account.
    - Authentication: Yes
    - Request Body: `Account`

- `GET /api/v1/accounts/{id}`
    - Description: Retrieves an account by its ID.
    - Authentication: Yes
    - Parameters:
        - `id` (string): The ID of the account.
    - Response: `Account`

- `GET /api/v1/accounts/number/{number}`
    - Description: Retrieves an account by its account number.
    - Authentication: Yes
    - Parameters:
        - `number` (string): The account number of the account.
    - Response: `Account`

- `PUT /api/v1/accounts/{id}`
    - Description: Updates account details by its ID.
    - Authentication: Yes
    - Parameters:
        - `id` (string): The ID of the account.
    - Request Body: `Account`

- `PUT /api/v1/accounts/balance/{account_number}/{amount}`
    - Description: Updates the account balance by its account number.
    - Authentication: Yes
    - Parameters:
        - `account_number` (string): The account number of the account.
        - `amount` (number): The new balance of the account.

- `DELETE /api/v1/accounts/{id}`
    - Description: Deletes an account by its ID.
    - Authentication: Yes
    - Parameters:
        - `id` (string): The ID of the account.

### Transactions

- `POST /api/v1/transaction`
    - Description: Creates a new transaction.
    - Authentication: Yes
    - Request Body: `Transaction`

- `GET /api/v1/transactions/{transaction_id}`
    - Description: Retrieves a transaction by its transaction ID.
    - Authentication: Yes
    - Parameters:
        - `transaction_id` (integer): The transaction ID of the transaction.
    - Response: `Transaction`

- `GET /api/v1/transactions/sender/{accountNumber}`
    - Description: Retrieves transactions by the sender's account number.
    - Authentication: Yes
    - Parameters:
        - `accountNumber` (string): The account number of the sender.
    - Response: `Array of Transaction`

- `GET /api/v1/transactions/receiver/{accountNumber}`
    - Description: Retrieves transactions by the receiver's account number.
    - Authentication: Yes
    - Parameters:
        - `accountNumber` (string): The account number of the receiver.
    - Response: `Array of Transaction`

- `PUT /api/v1/transaction/`
    - Description: Updates a transaction.
    - Authentication: Yes
    - Request Body: `Transaction`

- `DELETE /api/v1/transaction/{transaction_id}`
    - Description: Deletes a transaction by its transaction ID.
    - Authentication: Yes
    - Parameters:
        - `transaction_id` (integer): The transaction ID of the transaction.

- `GET /api/v1/transactions/{transaction_reference}`
    - Description: Retrieves a transaction by its transaction reference.
    - Authentication: Yes
    - Parameters:
        - `transaction_reference` (integer): The transaction reference of the transaction.
    - Response: `Transaction`

## Middlewares

The API uses the following middlewares:

- **JSONMiddleware**: Sets the content type to `application/json`.
- **LoggingMiddleware**: Logs every incoming request.
- **TimeoutMiddleware**: Times out all requests that take longer than 15 seconds.

## Running the Server

To run the server, you need to call the `Serve` method on the `Handler` struct. This method starts the server and listens for incoming requests. It also sets up a signal listener for `SIGTERM` signals to gracefully shut down the server when needed.

## Error Handling

The API uses standard HTTP status codes to indicate the success or failure of an API request. In general, codes in the `2xx` range indicate success, codes in the `4xx` range indicate an error that failed given the information provided (e.g., a required parameter was omitted, etc.), and codes in the `5xx` range indicate an error with the server.

If the API returns an error, the response will also include an error message in the `message` field of the response object.

You can copy and paste the above content into a `README.md` file on your GitHub repository.