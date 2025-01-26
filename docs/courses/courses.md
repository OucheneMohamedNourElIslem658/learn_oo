# API Endpoints Documentation

## Get Course
**Path:** `host/api/v1/courses/:course_id`  
**Method:** `GET`  
**Description:** Get course details.

### Headers
- **Authorization** (string, optional): Bearer id_token. Add it for author so he can get his uncompleted courses.

### Query Parameters
- **append_with** (string, optional): Include some infos associated to the course. Possible values: `image,author,video,requirements,objectives,categories,chapters,learners`.

### Responses
- **200 OK**
    ```json
    {
        "id": 13,
        "created_at": "2024-12-14T12:21:52.485722Z",
        "updated_at": "2024-12-15T07:17:26.217226Z",
        "deleted_at": null,
        "title": "test",
        "description": "test",
        "price": 0,
        "payment_price_id": null,
        "payment_product_id": null,
        "language": "en",
        "level": "bigener",
        "duration": 15,
        "rate": 3.5,
        "raters_count": 0,
        "is_completed": true,
        "requirements": [...],
        "objectives": [...],
        "video": {...},
        "image": {...},
        "author_id": "72125ca0-94f3-42e5-ac3a-0d797ec9078b",
        "author": {...},
        "categories": [...],
        "chapters": [...],
        "learners": [...]
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
    }
    ```
- **401 Unauthorized**
    ```json
    {
        "error": "id token expired"
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

## Get Courses
**Path:** `host/api/v1/courses`  
**Method:** `GET`  
**Description:** Get list of courses.

### Headers
- **Authorization** (string, optional): Bearer id_token. Add it for author so he can get his uncompleted courses.

### Query Parameters
- **append_with** (string, optional): Include some infos associated to the course. Possible values: `image,author,video,categories`.
- **title** (string, optional): Write part of desired title.
- **free_or_paid** (string, optional): Takes value `free` or `paid`.
- **language** (string, optional): Takes `en`, `fr`, or `ar`.
- **min_duration** (integer, optional): Minimum course duration in minutes (must be greater or equal to 5).
- **max_duration** (integer, optional): Maximum course duration in minutes (must be greater or equal to 5).
- **level** (string, optional): Must be one of `bigener`, `advanced`, or `medium`.
- **page_size** (integer, optional): Number of courses per page (at least 1, default is 10).
- **page** (integer, optional): Page number (at least 1, default is 1).
- **categories_names** (string, optional): List of category names, e.g., `Golang,dev,Acting`.

### Responses
- **200 OK**
    ```json
    {
        "count": 3,
        "courses": [...],
        "current_page": 1,
        "max_pages": 1
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
    }
    ```
- **401 Unauthorized**
    ```json
    {
        "error": "id token expired"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Delete Course
**Path:** `host/api/v1/courses/:course_id`  
**Method:** `DELETE`  
**Description:** Delete a course.

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
- **401 Unauthorized**
    ```json
    {
        "error": "id token expired"
    }
    ```
    ```json
    {
        "error": "user is not author"
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

## Update Course
**Path:** `host/api/v1/courses/:course_id`  
**Method:** `PUT`  
**Description:** Update course details.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **title** (string)
- **description** (string)
- **price** (number, equal or greater than 0)
- **language** (string, `en` or `ar` or `fr`)
- **level** (string, `advanced` or `bigener` or `medium`)
- **duration** (integer, min 5 minutes)
- **categories_names** (array of strings, e.g., `["Golang", "dev", "Acting"]`)
- **is_completed** (boolean)

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
        "error": "user is not author"
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

## Create Course
**Path:** `host/api/v1/courses`  
**Method:** `POST`  
**Description:** Create a new course.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **type**: `multipart/form-data`
- **body**:
    - **title** (string, required)
    - **description** (string, required)
    - **price** (number, equal or greater than 0, required)
    - **language** (string, `en` or `ar` or `fr`, required)
    - **level** (string, `advanced` or `bigener` or `medium`, required)
    - **duration** (integer, min 5 minutes, required)
    - **video** (file, required)
    - **image** (file, required)

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
        "error": "user is not author"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update Course Image
**Path:** `host/api/v1/courses/:course_id/image`  
**Method:** `PUT`  
**Description:** Update course image.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **type**: `multipart/form-data`
- **body**:
    - **image** (file, required)

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
        "image": "required",
        "request": "error indicated that you malformed your request"
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
        "error": "user is not author"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update Course Video
**Path:** `host/api/v1/courses/:course_id/video`  
**Method:** `PUT`  
**Description:** Update course video.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **type**: `multipart/form-data`
- **body**:
    - **video** (file, required)

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
        "video": "required",
        "request": "error indicated that you malformed your request"
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
        "error": "user is not author"
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

## Create Category
**Path:** `host/api/v1/courses/categories`  
**Method:** `POST`  
**Description:** Create a new category.

### Request Body
- **name** (string, required)

### Responses
- **201 Created**
- **400 Bad Request**
    ```json
    {
        "error": "bad request",
        "field 1": "validation of field 1",
        "...": "..."
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Delete Category
**Path:** `host/api/v1/courses/categories/:category_id`  
**Method:** `DELETE`  
**Description:** Delete a category.

### Responses
- **200 OK**
- **404 Not Found**
    ```json
    {
        "error": "category not found"
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
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Get Categories
**Path:** `host/api/v1/courses/categories`  
**Method:** `GET`  
**Description:** Get list of categories.

### Responses
- **200 OK**
    ```json
    {
        "categories": [...],
        "count": 9
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
    ```json
    {
        "error": "category name already exists"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Create Objective
**Path:** `host/api/v1/courses/:course_id/objectives`  
**Method:** `POST`  
**Description:** Create a new objective for a course.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **content** (string, required)

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

## Delete Objective
**Path:** `host/api/v1/courses/:course_id/objectives/:objective_id`  
**Method:** `DELETE`  
**Description:** Delete an objective from a course.

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
        "error": "objective not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Create Requirement
**Path:** `host/api/v1/courses/:course_id/requirements`  
**Method:** `POST`  
**Description:** Create a new requirement for a course.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **content** (string, required)

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

## Delete Requirement
**Path:** `host/api/v1/courses/:course_id/requirements/:requirement_id`  
**Method:** `DELETE`  
**Description:** Delete a requirement from a course.

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
        "error": "requirement not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```