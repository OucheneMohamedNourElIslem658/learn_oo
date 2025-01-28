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

## Check Course Completion
**Path:** `host/api/v1/check-course-completion/:courseID`  
**Method:** `GET`  
**Description:** Checks if the authenticated user has completed all tests in a course.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Parameters
- **courseID** (integer, required): The ID of the course to check for completion.

### Responses
- **200 OK**
    ```json
    {
        "course": {
            "id": 1,
            "name": "Go Programming"
        },
        "user": {
            "id": 1,
            "name": "John Doe"
        },
        "testResults": [
            {
                "test_id": 1,
                "has_succeed": true
            },
            {
                "test_id": 2,
                "has_succeed": true
            }
        ],
        "date": "2024-12-31T00:00:00Z"
    }
    ```

## Mark Lessons Learned
**Path:** `host/api/v1/mark-lessons-learned/:chapterID`  
**Method:** `POST`  
**Description:** Marks all lessons within a specified chapter as learned for a given user.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Parameters
- **chapterID** (integer, required): The ID of the chapter whose lessons are to be marked as learned.

### Responses
- **200 OK**
    ```json
    {
        "message": "Lessons marked as learned"
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "Invalid user ID type"
    }
    ```
    ```json
    {
        "error": "Invalid chapter ID"
    }
    ```
- **401 Unauthorized**
    ```json
    {
        "error": "User not authenticated"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "Internal server error message"
    }
    ```