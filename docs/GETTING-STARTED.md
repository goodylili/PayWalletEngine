






### <a name="health-check-ping"></a>**7. Health Check (Ping)**

- **Endpoint**: `/ping`
- **HTTP Method**: `GET`
- **Description**: Checks the health of the service.

| Parameter | Type | Description | Required |
|-----------|------|-------------|----------|
| -         | -    | -           | -        |

**Responses**:

- `200 OK`: Service is healthy.
- `500 Internal Server Error`: Service is down or facing issues.

---

### <a name="reset-password"></a>**8. Reset Password**

- **Endpoint**: `/password/reset`
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