# Getting Started with PayWalletEngine API

In this guide, we'll walk you through the initial steps to get the **PayWalletEngine** API up and running on your local
environment.

## Table of Contents

1. [Health Check (Ping)](#health-check-ping)
2. [Database Health Check](#database-health-check)
3. [API Endpoints Overview](#api-endpoints-overview)
4. [Further Reading](#further-reading)

---

### <a name="health-check-ping"></a>**1. Health Check (Ping)**

Once you have the server running, you can check the health of the service by using the following endpoint:

- **Endpoint**: `/ping`
- **HTTP Method**: `GET`
- **Description**: Checks the health of the service.

**Responses**:

- `200 OK`: Service is healthy.
- `500 Internal Server Error`: Service is down or facing issues.

---

### <a name="database-health-check"></a>**2. Database Health Check**

To ensure your database connectivity is intact:

- **Endpoint**: `/alive`
- **HTTP Method**: `GET`
- **Description**: Checks if the database is alive and responding.

| Parameter | Type | Description | Required |
|-----------|------|-------------|----------|
| -         | -    | -           | -        |

**Responses**:

- `200 OK`: Database is alive and responding.
- `500 Internal Server Error`: Database connection issues or the service is down.

---

### <a name="api-endpoints-overview"></a>**3. API Endpoints Overview**

For a deep dive into each category of endpoints, refer to the detailed documentation:


- [Authentication](./auth.md)
- [Users](./users.md)
- [Accounts](./accounts.md)
- [Transactions](./transactions.md)
- [Error Codes](./errors.md)

---

### <a name="further-reading"></a>**4. Further Reading**

As you delve into **PayWalletEngine** API, we recommend familiarizing yourself with the main [README.md](../README.md)
for a holistic understanding of the system, its architecture, and functionalities. For any issues or contributions,
please follow the guidelines specified in the main README.

---
