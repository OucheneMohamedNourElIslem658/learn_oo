# Users Service:

## Endpoint: `Register User With Email And Password`

##### **Method**: `POST`

##### **URL**: `http://localhost:8000/api/v1/users/auth/register-with-email-and-password`

##### **Request**

###### **Headers**:
- `Content-Type`: `application/json`

###### **Request Body**:
```json
{
  "email": "example@domain.extention",
  "password": "its lenght must be greater than 5",
  "full_name": "its lenght must be greater than 0"
}
```

##### **Response**:

###### **Success Responses**:

- **201 Created**  
  **Description:** User successfully registered.  
  **Response Body**:
  ```json
  {
    "id": "123456",
    "email": "example@domain.extention",
    "full_name": "John Doe",
    "status": "Registered"
  }
  ```

###### **Fail Responses**:

- **400 Bad Request**  
  **Description:** User successfully registered.  
  **Response Body**:
  ```json
  {
    "email": "invalid,required",
    "password": "invalid",
    "full_name": "required",
  }
  ```

  ```json
  {
    "message": "email already in use",
  }
  ```