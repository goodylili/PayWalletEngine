# PayWalletEngine API Error Reference

Welcome to the PayWalletEngine API Error Reference. This document provides a comprehensive list of error codes,
descriptions, and recommended actions for handling each error. We aim to ensure a smooth and error-free experience when
using our API.

## Table of Contents

1. [Introduction](#introduction)
2. [Client Errors](#client-errors)
3. [Server Errors](#server-errors)
4. [Custom Errors](#custom-errors)
5. [Conclusion](#conclusion)

## 1. Introduction

The PayWalletEngine API may return various error codes to help you understand and handle unexpected situations. Below,
we've documented common error codes specific to our project, their descriptions, and suggestions on how to proceed when
encountering them.

## 2. Client Errors

| Error Code | Description  | Suggested Action                                                                                      |
|------------|--------------|-------------------------------------------------------------------------------------------------------|
| 400        | Bad Request  | Check the request parameters for correctness. Refer to the API documentation for the correct format.  |
| 401        | Unauthorized | Ensure you have valid authentication credentials. Contact support if authorization issues persist.    |
| 403        | Forbidden    | Confirm you have the necessary permissions for the requested operation. Contact your admin if needed. |
| 404        | Not Found    | Verify the requested resource exists and the URL is correct. Check the resource identifiers.          |
| 409        | Conflict     | Ensure the resource is in the correct state. Avoid duplicate resource creation.                       |

## 3. Server Errors

| Error Code | Description           | Suggested Action                                                                                                |
|------------|-----------------------|-----------------------------------------------------------------------------------------------------------------|
| 500        | Internal Server Error | Report the issue to our support team for investigation. Avoid repeated requests until the issue is resolved.    |
| 503        | Service Unavailable   | Retry the request after some time or check for maintenance announcements. Implement appropriate error handling. |

## 4. Custom Errors

| Error Code | Description           | Suggested Action                                                                                              |
|------------|-----------------------|---------------------------------------------------------------------------------------------------------------|
| 422        | Invalid Email Format  | Verify that the email address follows the correct format.                                                     |
| 423        | User Not Found        | Double-check the user identifier and ensure it corresponds to an existing user. Handle this error gracefully. |
| 500        | Password Reset Failed | Report the issue to our support team for resolution. Avoid repeated password reset attempts.                  |

## 5. Conclusion

We hope this error reference helps you effectively handle errors encountered while using the PayWalletEngine API. If you
encounter an error not listed here or require further assistance, please refer to the API documentation or contact our
support team.

