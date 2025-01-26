# API Endpoints Documentation

## Create Chapter
**Path:** `host/api/v1/courses/:course_id/chapters`  
**Method:** `POST`  
**Description:** Create a new chapter.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **title** (string, required)
- **description** (string, required)

### Responses
- **201 Created**
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
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
        "error": "user is not an author"
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

## Update Chapter
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id`  
**Method:** `PUT`  
**Description:** Update chapter details.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **title** (string)
- **description** (string)

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
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
        "error": "user is not an author"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "course not found"
    }
    ```
    ```json
    {
        "error": "chapter not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Get Chapter
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id`  
**Method:** `GET`  
**Description:** Get chapter details.

### Query Parameters
- **append_with** (string, optional): Include some infos associated to the chapter. Possible values: `course,lessons,test`.

### Responses
- **200 OK**
    ```json
    {
        "id": 15,
        "created_at": "2024-12-15T09:51:51.236287Z",
        "updated_at": "2024-12-15T09:51:51.236287Z",
        "deleted_at": null,
        "title": "Go Basics",
        "description": "Learn the basic syntax and data structures of Go.",
        "course_id": 18,
        "course": {...},
        "lessons": [...],
        "test": {...}
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "bad request",
        "field 1": "validation of field 1",
        "...": "..."
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "course not found"
    }
    ```
    ```json
    {
        "error": "chapter not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Delete Chapter
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id`  
**Method:** `DELETE`  
**Description:** Delete a chapter.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
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
        "error": "user is not an author"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "course not found"
    }
    ```
    ```json
    {
        "error": "chapter not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```