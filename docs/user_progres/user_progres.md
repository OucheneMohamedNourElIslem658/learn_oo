# API Endpoints Documentation

## Start Course
**Path:** `host/api/v1/user-courses/start-course/:course_id/`  
**Method:** `POST`  
**Description:** Create course session for the user.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **payment_success_url** (string, required)
- **payment_fail_url** (string, required)

### Responses
- **202 Accepted** (only if course is paid)
    ```json
    {
        "payment_url": "course payment url"
    }
    ```
- **201 Created** (only if course is free)
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
    }
    ```
    ```json
    {
        "error": "user is already a learner"
    }
    ```
    ```json
    {
        "error": "this user is the author of this course"
    }
    ```
    ```json
    {
        "error": "bad request",
        "field 1": "validation of field 1",
        "...": "..."
    }
    ```
- **401 Unauthorized**
    ```json
    {
        "error": "id token expired"
    }
    ```
    ```json
    {
        "error": "requester is not a user"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "course not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```