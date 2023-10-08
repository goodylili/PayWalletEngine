# **Users API Documentation**

--- 

## **Overview**
The Users API allows the management of user accounts within the platform. It provides endpoints for creating, retrieving, updating, and managing user accounts.

---

### **Base URL**: `/api/v1`

---


### **Models**

#### **The User Object**

| Field       | Type    | Description                            | Restrictions |
|-------------|---------|----------------------------------------|--------------|
| `username`  | string  | Username assigned to the user.         | Unique      |
| `email`     | string  | Email address linked to the user.      | Unique, Valid Email Format |
| `password`  | string  | User's hashed password (write-only).   | At least 8 characters |
| `is_active` | boolean | Indicates if the user is active.       | true or false |

---

## **Endpoints**:

### **1. Create User**

- **Endpoint**: `/user/create`
- **HTTP Method**: `POST`
- **Description**: Registers a new user with the provided details.

| Parameter | Type   | Description | Required |
|-----------|--------|-------------|----------|
| -         | -      | -           | -        |

**Request Body**:
```json
{
   "username": "string",
   "email": "string",
   "password": "string"
}
```

**Responses**:
- `201 Created`: Successfully created a user. Returns the user data.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### **2. Retrieve User by ID**

- **Endpoint**: `/users/{id}`
- **HTTP Method**: `GET`
- **Description**: Fetches details of a specific user using their unique ID.

| Parameter | Type | Description         | Required |
|-----------|------|---------------------|----------|
| id        | int  | Unique identifier of the user | Yes      |

**Responses**:
- `200 OK`: Successfully fetched the user data.
- `400 Bad Request`: Invalid ID format.
- `404 Not Found`: User with the provided ID doesn't exist.

---

### **3. Update User**

- **Endpoint**: `/users/{id}/update`
- **HTTP Method**: `PUT`
- **Description**: Modifies the details of an existing user.

| Parameter | Type | Description         | Required |
|-----------|------|---------------------|----------|
| id        | int  | Unique identifier of the user | Yes      |

**Request Body**:
```json
{
   "username": "string",
   "email": "string",
   "password": "string",
   "is_active": true
}
```

**Responses**:
- `200 OK`: Successfully updated the user data.
- `400 Bad Request`: Invalid input or malformed request.
- `404 Not Found`: User with the provided ID doesn't exist.
- `500 Internal Server Error`: Unexpected server error.

Certainly! Let's continue documenting the remaining endpoints:

---

### **4. Change User Status**

- **Endpoint**: `/users/{id}/status`
- **HTTP Method**: `PUT`
- **Description**: Updates the status of a user (activate/deactivate).

| Parameter | Type | Description                         | Required |
|-----------|------|-------------------------------------|----------|
| id        | int  | Unique identifier of the user.      | Yes      |

**Request Body**:
```json
{
   "is_active": true
}
```

**Responses**:
- `200 OK`: Successfully updated the user's status.
- `400 Bad Request`: Invalid input or malformed request.
- `404 Not Found`: User with the provided ID doesn't exist.
- `500 Internal Server Error`: Unexpected server error.

---

### **5. Retrieve User by Email**

- **Endpoint**: `/users/email/{email}`
- **HTTP Method**: `GET`
- **Description**: Fetches user details based on their email address.

| Parameter | Type   | Description                  | Required |
|-----------|--------|------------------------------|----------|
| email     | string | Email address of the user.   | Yes      |

**Responses**:
- `200 OK`: Successfully fetched the user data.
- `400 Bad Request`: Invalid email format.
- `404 Not Found`: User with the provided email doesn't exist.
- `500 Internal Server Error`: Unexpected server error.

---

### **6. Retrieve User by Username**

- **Endpoint**: `/users/username/{username}`
- **HTTP Method**: `GET`
- **Description**: Fetches user details based on their username.

| Parameter | Type   | Description                     | Required |
|-----------|--------|---------------------------------|----------|
| username  | string | Username of the user.           | Yes      |

**Responses**:
- `200 OK`: Successfully fetched the user data.
- `404 Not Found`: User with the provided username doesn't exist.
- `500 Internal Server Error`: Unexpected server error.

---

### **7. Health Check (Ping)**

- **Endpoint**: `/users/ping`
- **HTTP Method**: `GET`
- **Description**: Checks the health of the service.

| Parameter | Type | Description | Required |
|-----------|------|-------------|----------|
| -         | -    | -           | -        |

**Responses**:
- `200 OK`: Service is healthy.
- `500 Internal Server Error`: Service is down or facing issues.

---

### **8. Reset Password**

- **Endpoint**: `/users/password/reset`
- **HTTP Method**: `PUT`
- **Description**: Resets the user's password.

**Request Body**:
```json
{
   "username": "string",
   "email": "string",
   "password": "new_password"
}
```

**Responses**:
- `200 OK`: Password successfully reset.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

