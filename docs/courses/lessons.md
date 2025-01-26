# API Endpoints Documentation

## Create Lesson with Content
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/lessons/create-with-content`  
**Method:** `POST`  
**Description:** Create lessons with document content.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **title** (string, required)
- **description** (string, required)
- **content** (map[string]any, required): Representation of Quill doc JSON.

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
        "error": "chapter not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Create Lesson with Video
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/lessons/create-with-video`  
**Method:** `POST`  
**Description:** Create lessons with video content.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **type** (string, required): `multipart/form-data`
- **body** (object, required)
  - **title** (string, required)
  - **description** (string, required)
  - **video** (file, required)

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
        "error": "chapter not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update Lesson
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/lessons/:lesson_id`  
**Method:** `PUT`  
**Description:** Update lesson details.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **title** (string)
- **description** (string)
- **content** (map[string]any, required): Representation of Quill doc JSON.

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
    ```json
    {
        "error": "lesson is a video"
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
        "error": "chapter not found"
    }
    ```
    ```json
    {
        "error": "lesson not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Get Lesson
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id`  
**Method:** `GET`  
**Description:** Get lesson details.

### Query Parameters
- **append_with** (string, optional): Include some infos associated to the chapter. Possible values: `chapter,learners`.

### Responses
- **200 OK**
    ```json
    {
        "id": 23,
        "created_at": "2024-12-15T09:51:51.336273Z",
        "updated_at": "2024-12-15T09:51:51.336273Z",
        "deleted_at": null,
        "title": "Variables and Constants",
        "description": "Learn how to declare variables and constants in Go.",
        "is_video": false,
        "content": {
            "ops": [
                {
                    "insert": "Variables and constants are fundamental in Go.\n"
                }
            ]
        },
        "video": null,
        "chapter_id": 15,
        "chapter": {
            "id": 15,
            "created_at": "2024-12-15T09:51:51.236287Z",
            "updated_at": "2024-12-15T09:51:51.236287Z",
            "deleted_at": null,
            "title": "Go Basics",
            "description": "Learn the basic syntax and data structures of Go.",
            "course_id": 18,
            "course": null,
            "lessons": null,
            "test": null
        },
        "learners": [
            {
                "id": "92d856d0-a0ee-4220-81b4-66e8adc073fc",
                "created_at": "2024-12-11T13:11:39.167269Z",
                "updated_at": "2024-12-13T13:32:04.17469Z",
                "deleted_at": null,
                "email": "m_ouchene@estin.dz",
                "password": "$2a$10$0pImKZSJKjGFcIKcpz27zuhNXwolVaiUbfBPqZZJJxLQNCGqRQjzq",
                "full_name": "mohamed OUCHENE",
                "email_verified": true,
                "image": null,
                "author_profile": null,
                "courses": null,
                "lessons": null,
                "tests": null
            }
        ]
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
        "error": "chapter not found"
    }
    ```
    ```json
    {
        "error": "lesson not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Delete Lesson
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/lessons/:lesson_id`  
**Method:** `DELETE`  
**Description:** Delete a lesson.

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
        "error": "chapter not found"
    }
    ```
    ```json
    {
        "error": "lesson not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Create Lesson Video
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/lessons/:lesson_id/video`  
**Method:** `POST`  
**Description:** Create lessons with video content.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **type** (string, required): `multipart/form-data`
- **body** (object, required)
  - **title** (string, required)
  - **description** (string, required)
  - **video** (file, required)

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
        "error": "lesson is not a video"
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
        "error": "chapter not found"
    }
    ```
    ```json
    {
        "error": "lesson not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```