# Accounts API Documentation
---

## Overview

The Accounts API allows the management of user accounts within the platform. It provides endpoints for creating,
retrieving, updating, and managing bank accounts.

---

## Index

- **[Endpoints](#endpoints)**
    - [Create Account](#1-create-account)
    - [Retrieve Account by ID](#2-retrieve-account-by-id)
    - [Update Account Details](#3-update-account-details)
    - [Retrieve User Details by Account Number](#4-retrieve-user-details-by-account-number)
    - [Retrieve Account by Account Number](#5-retrieve-account-by-account-number)
    - [Retrieve Accounts by User ID](#6-retrieve-accounts-by-user-id)

---

### **Base URL**: `/accounts/api/v1`

---

### **Models**

### <a name="the-account-object"></a>**The Account Object**

| Field            | Type    | Description                                  |
|------------------|---------|----------------------------------------------|
| `id`             | uint    | Unique identifier of the account.            |
| `account_number` | int64   | Unique account number.                       |
| `account_type`   | string  | Type of the bank account.                    |
| `balance`        | float64 | Current balance of the account.              |
| `user_id`        | uint    | ID of the user associated with this account. |

---

## <a name="endpoints"></a>**Endpoints**:

### <a name="create-account"></a>**1. Create Account**

- **Endpoint**: `create`
- **HTTP Method**: `POST`
- **Description**: Creates a new bank account with the provided details.

**Request Body**:

```json
{
  "account_type": "Checking",
  "balance": 500.00,
  "user_id": 1
}
```

**Responses**:

- `201 Created`: Successfully created an account.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="retrieve-account-by-id"></a>**2. Retrieve Account by ID**

- **Endpoint**: `/{id}`
- **HTTP Method**: `GET`
- **Description**: Fetches details of a specific bank account using its unique ID.

| Parameter | Type | Description                      | Required |
|-----------|------|----------------------------------|----------|
| id        | int  | Unique identifier of the account | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the account data.
- `400 Bad Request`: Invalid ID format.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="update-account-details"></a>**3. Update Account Details**

- **Endpoint**: `/{id}/update`
- **HTTP Method**: `PUT`
- **Description**: Modifies the details of an existing bank account.

| Parameter | Type | Description                      | Required |
|-----------|------|----------------------------------|----------|
| id        | int  | Unique identifier of the account | Yes      |

**Request Body**:

```json
{
  "id": 1,
  "account_type": "Investment",
  "balance": 2000.00,
  "user_id": "uint"
}
```

**Responses**:

- `200 OK`: Successfully updated the account data.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="retrieve-user-details-by-account-number"></a>**4. Retrieve User Details by Account Number**

- **Endpoint**: `/{account_number}/user`
- **HTTP Method**: `GET`
- **Description**: Fetches user details associated with a given account number.

| Parameter      | Type | Description                              | Required |
|----------------|------|------------------------------------------|----------|
| account_number | int  | Account number associated with the user. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the user details.
- `400 Bad Request`: Invalid account number format.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="retrieve-account-by-account-number"></a>**5. Retrieve Account by Account Number**

- **Endpoint**: `/number/{number}`
- **HTTP Method**: `GET`
- **Description**: Fetches bank account details using a specific account number.

| Parameter | Type | Description              | Required |
|-----------|------|--------------------------|----------|
| number    | int  | Specific account number. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the account data.
- `400 Bad Request`: Invalid account number format.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="retrieve-accounts-by-user-id"></a>**6. Retrieve Accounts by User ID**

- **Endpoint**: `/user/{user_id}`
- **HTTP Method**: `GET`
- **Description**: Fetches all bank accounts associated with a given user ID.

| Parameter | Type | Description               | Required |
|-----------|------|---------------------------|----------|
| user_id   | int  | ID of the associated user | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the accounts data.
- `400 Bad Request`: Invalid user ID format.
- `500 Internal Server Error`: Unexpected server error.

---
