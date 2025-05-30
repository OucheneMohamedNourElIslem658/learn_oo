# API Endpoints Documentation

## Register User with Email and Password
**Path:** `host/api/v1/users/auth/register-with-email-and-password`  
**Method:** `POST`  
**Description:** Register user with email and password.

### Request Body
- **full_name** (string, required): The length needs to be greater than 0.
- **email** (string, required): Example: `m_ouchene@estin.dz`.
- **password** (string, required): The length needs to be greater than 4.

### Responses
- **400 Bad Request**
    ```json
    {
        "message": {
            "full_name": "required",
            "password": "required,invalid",
            "email": "required,invalid"
        }
    }
    ```
    ```json
    {
        "message": "email already in use"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "message": "error message (contact me when you see one)"
    }
    ```

## Login User with Email and Password
**Path:** `host/api/v1/users/auth/login-with-email-and-password`  
**Method:** `POST`  
**Description:** Login user with email and password.

### Request Body
- **email** (string, required): Example: `m_ouchene@estin.dz`.
- **password** (string, required): The length needs to be greater than 4.

### Responses
- **200 OK**
    ```json
    {
        "cookies": [
            {
                "name": "id_token",
                "description": "id token to authorize user to do some actions"
            },
            {
                "name": "refresh_token",
                "description": "refresh token to refresh the id token when it expires or when user role changes"
            }
        ]
    }
    ```
- **400 Bad Request**
    ```json
    {
        "message": {
            "password": "required,invalid",
            "email": "required,invalid"
        }
    }
    ```
    ```json
    {
        "message": "email not verified"
    }
    ```
    ```json
    {
        "message": "incorrect password"
    }
    ```
- **404 Not Found**
    ```json
    {
        "message": "email not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "message": "error message (contact me when you see one)"
    }
    ```

## Send Email Verification Link
**Path:** `host/api/v1/users/auth/send-email-verification-link`  
**Method:** `POST`  
**Description:** Link sent by email to verify user email.

### Request Body
- **email** (string, required): Example: `m_ouchene@estin.dz`.

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "message": {
            "email": "required,invalid"
        }
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "message": "error message (contact me when you see one)"
    }
    ```

## Send Password Reset Link
**Path:** `host/api/v1/users/auth/send-password-reset-link`  
**Method:** `POST`  
**Description:** Link sent by email to reset user password.

### Request Body
- **email** (string, required): Example: `m_ouchene@estin.dz`.

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "message": {
            "email": "required,invalid"
        }
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "message": "error message (contact me when you see one)"
    }
    ```

## OAuth Login
**Path:** `host/api/v1/users/auth/oauth/:provider/login`  
**Method:** `GET`  
**Description:** Navigate to this endpoint to login with your Facebook or Google account.

### Path Parameters
- **provider** (string, required): Possible values: `google,facebook`.

### Query Parameters
- **success_url** (string, required): URL to navigate to when login is successful.
- **failure_url** (string, required): URL to navigate to when login fails.

### Responses
- **200 OK**
    ```json
    {
        "cookies": [
            {
                "name": "id_token",
                "description": "id token to authorize user to do some actions"
            },
            {
                "name": "refresh_token",
                "description": "refresh token to refresh the id token when it expires or when user role changes"
            }
        ]
    }
    ```
- **400 Bad Request**
    ```json
    {
        "message": {
            "success_url": "required",
            "failure_url": "required"
        }
    }
    ```
    ```json
    {
        "message": "provider not supported"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "message": "error message (contact me when you see one)"
    }
    ```

## Refresh ID Token
**Path:** `host/api/v1/users/refresh-id-token`  
**Method:** `GET`  
**Description:** Refresh your ID token.

### Headers
- **Authorization** (string, required): Bearer refresh_token.

### Responses
- **200 OK**
    ```json
    {
        "cookies": [
            {
                "name": "id_token",
                "description": "new id token to authorize user to do some actions"
            }
        ]
    }
    ```
- **400 Bad Request**
    ```json
    {
        "message": "undefined authorization"
    }
    ```
    ```json
    {
        "message": "undefined refresh token"
    }
    ```
- **401 Unauthorized**
    ```json
    {
        "message": "refresh token expired"
    }
    ```
- **404 Not Found**
    ```json
    {
        "message": "user not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "message": "error message (contact me when you see one)"
    }
    ```

## gRPC Service Documentation

### Service: `AuthService`

#### RegisterWithEmailAndPassword
- **rpc:** `RegisterWithEmailAndPassword(RegisterRequest) returns (RegisterResponse)`
- **Description:** Register a user with email and password.

##### RegisterRequest
| Field      | Type   | Required | Description                              |
|------------|--------|----------|------------------------------------------|
| full_name  | string | Yes      | Full name of the user (non-empty)        |
| email      | string | Yes      | User email (e.g., `m_ouchene@estin.dz`)  |
| password   | string | Yes      | Password (length > 4)                    |

##### RegisterResponse
| Field   | Type   | Description                |
|---------|--------|---------------------------|
| message | string | Registration status message|

---

#### LoginWithEmailAndPassword
- **rpc:** `LoginWithEmailAndPassword(LoginRequest) returns (LoginResponse)`
- **Description:** Login a user with email and password.

##### LoginRequest
| Field    | Type   | Required | Description                              |
|----------|--------|----------|------------------------------------------|
| email    | string | Yes      | User email (e.g., `m_ouchene@estin.dz`)  |
| password | string | Yes      | Password (length > 4)                    |

##### LoginResponse
| Field         | Type   | Description                                 |
|---------------|--------|---------------------------------------------|
| id_token      | string | ID token for user authorization             |
| refresh_token | string | Token to refresh the ID token               |

---

#### SendEmailVerificationLink
- **rpc:** `SendEmailVerificationLink(EmailLinkRequest) returns (google.protobuf.Empty)`
- **Description:** Send an email verification link to the user.

##### EmailLinkRequest
| Field | Type   | Required | Description                              |
|-------|--------|----------|------------------------------------------|
| email | string | Yes      | User email (e.g., `m_ouchene@estin.dz`)  |

---

#### SendPasswordResetLink
- **rpc:** `SendPasswordResetLink(EmailLinkRequest) returns (google.protobuf.Empty)`
- **Description:** Send a password reset link to the user.

##### EmailLinkRequest
| Field | Type   | Required | Description                              |
|-------|--------|----------|------------------------------------------|
| email | string | Yes      | User email (e.g., `m_ouchene@estin.dz`)  |

---

#### RefreshIDToken
- **rpc:** `RefreshIDToken(google.protobuf.Empty) returns (RefreshIDTokenReponse)`
- **Description:** Refresh the user's ID token.

##### RefreshIDTokenReponse
| Field    | Type   | Description                      |
|----------|--------|----------------------------------|
| id_token | string | New ID token for authorization   |