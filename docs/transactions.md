# Transactions API Documentation

## Overview

The Transactions API provides functionalities for executing and retrieving transactions within the platform. This
includes credit, debit, and fund transfers as well as retrieving transactions based on various parameters.

## Index

- **[Endpoints](#endpoints)**
    - [Get Transactions from Account](#1-get-transactions-from-account)
    - [Get Transaction by Reference](#2-get-transaction-by-reference)
    - [Credit Account](#3-credit-account)
    - [Debit Account](#4-debit-account)
    - [Transfer Funds](#5-transfer-funds)
    - [Get User, Account and Transaction by Transaction ID](#6-get-user-account-and-transaction-by-transaction-id)
    - [Get Account by Transaction ID](#7-get-account-by-transaction-id)


### **Base URL**: `/api/v1/transactions`

---

### **Models**

### <a name="the-transaction-object"></a>**The Transaction Object**

| Field            | Type      | Description                                |
|------------------|-----------|--------------------------------------------|
| `transaction_id` | uuid.UUID | Unique identifier of the transaction.      |
| `amount`         | float   | Amount involved in the transaction.        |
| `paymentMethod`  | string    | Method of payment used in the transaction. |
| `type`           | string    | Type of the transaction (credit, debit).   |
| `status`         | string    | Status of the transaction.                 |
| `description`    | string    | Description or reason for the transaction. |

---

## <a name="endpoints"></a>**Endpoints**:

### <a name="1-get-transactions-from-account"></a>**1. Get Transactions from Account**

- **Endpoint**: `/account/{account_number}`
- **HTTP Method**: `GET`
- **Description**: Retrieves all transactions made from a specific account.

| Parameter      | Type | Description                   | Required |
|----------------|------|-------------------------------|----------|
| account_number | int  | Account number to fetch from. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the transaction data.
- `400 Bad Request`: Invalid account number format.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="2-get-transaction-by-reference"></a>**2. Get Transaction by Reference**

- **Endpoint**: `/reference/{transaction_reference}`
- **HTTP Method**: `GET`
- **Description**: Fetches a specific transaction using its reference number.

| Parameter             | Type | Description                                   | Required |
|-----------------------|------|-----------------------------------------------|----------|
| transaction_reference | int  | Reference number of the transaction to fetch. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the transaction data.
- `400 Bad Request`: Invalid reference format.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="3-credit-account"></a>**3. Credit Account**

- **Endpoint**: `/credit`
- **HTTP Method**: `POST`
- **Description**: Credits a specific account with the provided details.

**Request Body**:

```json
{
  "receiver_account_number": 5867466691,
  "amount": 100.00,
  "description": "Credit transaction",
  "payment_method": "Credit Card"
}
```

**Responses**:

- `201 Created`: Successfully credited the account.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="4-debit-account"></a>**4. Debit Account**

- **Endpoint**: `/debit`
- **HTTP Method**: `POST`
- **Description**: Debits a specific account based on the provided details.

**Request Body**:

```json
{
  "sender_account_number": 5867466691,
  "amount": 10,
  "description": "Fundsssss",
  "payment_method": "Card"
}
```

**Responses**:

- `201 Created`: Successfully debited the account.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="5-transfer-funds"></a>**5. Transfer Funds**

- **Endpoint**: `/transfer`
- **HTTP Method**: `POST`
- **Description**: Transfers funds from a sender to a receiver based on provided details.

**Request Body**:

```json
{
  "sender_account_number": 5867466691,
  "receiver_account_number": 1677203234,
  "amount": 50,
  "description": "Cha Ching",
  "payment_method": "app transfer"
}

```

**Responses**:

- `201 Created`: Successfully transferred the funds.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="6-get-user-account-and-transaction-by-transaction-id"></a>**6. Get User,

Account, and Transaction by Transaction ID**

- **Endpoint**: `/details/{transaction_id}`
- **HTTP Method**: `GET`
- **Description**: Fetches details of the user, account, and transaction based on a transaction ID.

| Parameter      | Type | Description                               | Required |
|----------------|------|-------------------------------------------|----------|
| transaction_id | int  | Transaction ID used to fetch the details. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the details.
- `400 Bad Request`: Invalid transaction ID format.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="7-get-account-by-transaction-id"></a>**7. Get Account by Transaction ID**

- **Endpoint**: `/account-details/{transaction_id}`
- **HTTP Method**: `GET`
- **Description**: Fetches account details based on a transaction ID.

| Parameter      | Type | Description                               | Required |
|----------------|------|-------------------------------------------|----------|
| transaction_id | int  | Transaction ID used to fetch the account. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the account details.
- `400 Bad Request`: Invalid transaction ID format.
- `500 Internal Server Error`: Unexpected server error.

---
