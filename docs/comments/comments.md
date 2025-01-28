## API Endpoints Documentation

# API Endpoints Documentation

## Create Comment
**Path:** `host/api/v1/comments/:lesson_id`  
**Method:** `POST`  
**Description:** Create a new comment for a lesson.

### Parameters
- **lesson_id** (integer, required): ID of the lesson.

### Request Body
- **content** (string, required): The content of the comment.

### Responses
- **201 Created**
    ```json
    {
        "id": 1,
        "content": "This is a comment",
        "lesson_id": 101,
        "user_id": "user123"
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "Invalid input"
    }
    ```

## Retrieve Comment
**Path:** `host/api/v1/comments/:id`  
**Method:** `GET`  
**Description:** Retrieve a comment by its ID.

### Parameters
- **id** (integer, required): ID of the comment.

### Responses
- **200 OK**
    ```json
    {
        "id": 1,
        "content": "This is a comment",
        "lesson_id": 101,
        "user_id": "user123"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "Comment not found"
    }
    ```

## Retrieve Comments for Lesson
**Path:** `host/api/v1/comments/lesson/:lesson_id`  
**Method:** `GET`  
**Description:** Retrieve all comments for a lesson.

### Parameters
- **lesson_id** (integer, required): ID of the lesson.

### Responses
- **200 OK**
    ```json
    [
        {
            "id": 1,
            "content": "This is a comment",
            "lesson_id": 101,
            "user_id": "user123"
        },
        {
            "id": 2,
            "content": "This is another comment",
            "lesson_id": 101,
            "user_id": "user456"
        }
    ]
    ```

## Retrieve User Comments
**Path:** `host/api/v1/comments/user`  
**Method:** `GET`  
**Description:** Retrieve comments by the authenticated user.

### Responses
- **200 OK**
    ```json
    [
        {
            "id": 1,
            "content": "This is a comment",
            "lesson_id": 101,
            "user_id": "user123"
        },
        {
            "id": 3,
            "content": "This is a third comment",
            "lesson_id": 102,
            "user_id": "user123"
        }
    ]
    ```

## Delete Comment
**Path:** `host/api/v1/comments/:id`  
**Method:** `DELETE`  
**Description:** Delete a comment.

### Parameters
- **id** (integer, required): ID of the comment to delete.

### Responses
- **200 OK**
    ```json
    {
        "message": "Comment deleted successfully"
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "Invalid comment ID"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "Server error"
    }
    ```

## Retrieve All Notifications
**Path:** `host/api/v1/notifications/all_notification`  
**Method:** `GET`  
**Description:** Retrieve all notifications for the authenticated user.

### Responses
- **200 OK**
    ```json
    [
        {
            "id": 1,
            "title": "New Comment",
            "description": "You have a new comment on your lesson",
            "user_id": "user123"
        },
        {
            "id": 2,
            "title": "Lesson Updated",
            "description": "A lesson you are following has been updated",
            "user_id": "user123"
        }
    ]
    ```