# API Endpoints Documentation

## Create Test
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/tests`  
**Method:** `POST`  
**Description:** Create a new test.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **max_chances** (integer, min=1, default=1)

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

## Update Test
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/tests/:test_id`  
**Method:** `PUT`  
**Description:** Update test details.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **max_chances** (integer, min=1, default=1)

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
        "error": "test not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Get Test
**Path:** `host/api/v1/courses/:course_id/tests/:test_id`  
**Method:** `GET`  
**Description:** Get test details.

### Query Parameters
- **append_with** (string, optional): Include some infos associated to the test. Possible values: `questions,chapter`.

### Responses
- **200 OK**
    ```json
    {
        "id": 2,
        "questions": [
            {
                "id": 1,
                "content": "What is the organization that created Golang?",
                "description": "it is not that hard",
                "duration": 30,
                "test_id": 2,
                "options": [
                    {
                        "id": 1,
                        "content": "google",
                        "is_correct": true,
                        "question_id": 1
                    },
                    {
                        "id": 2,
                        "content": "facebook",
                        "is_correct": false,
                        "question_id": 1
                    },
                    {
                        "id": 3,
                        "content": "twitter",
                        "is_correct": false,
                        "question_id": 1
                    },
                    {
                        "id": 4,
                        "content": "microsoft",
                        "is_correct": false,
                        "question_id": 1
                    }
                ]
            },
            {
                "id": 2,
                "content": "What is the organization that created Golang?",
                "description": "it is not that hard",
                "duration": 30,
                "test_id": 2,
                "options": [
                    {
                        "id": 5,
                        "content": "google",
                        "is_correct": true,
                        "question_id": 2
                    },
                    {
                        "id": 6,
                        "content": "google",
                        "is_correct": false,
                        "question_id": 2
                    },
                    {
                        "id": 7,
                        "content": "twitter",
                        "is_correct": false,
                        "question_id": 2
                    },
                    {
                        "id": 8,
                        "content": "microsoft",
                        "is_correct": false,
                        "question_id": 2
                    }
                ]
            }
        ],
        "max_chances": 3,
        "chapter_id": 24,
        "chapter": {
            "id": 24,
            "title": "Getting Started",
            "description": "Introduction to acting and the world of performance.",
            "course_id": 25
        }
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

## Delete Test
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/tests/:test_id`  
**Method:** `DELETE`  
**Description:** Delete a test.

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
        "error": "test not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Create Question
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/tests/:test_id/questions`  
**Method:** `POST`  
**Description:** Create a new question.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **content** (string, required)
- **description** (string, optional)
- **duration** (integer, min=10, required)
- **options** (array, required): Must contain at least two elements, without duplicates, and at least one correct option.
    - **option** (string, required)
    - **is_correct** (boolean, required)

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
        "error": "test not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update Question
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/tests/:test_id/questions/:question_id`  
**Method:** `PUT`  
**Description:** Update question details.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **content** (string)
- **description** (string)
- **duration** (integer, min=10)
- **options** (array): Must contain at least two elements, without duplicates, and at least one correct option.
    - **option** (string, required)
    - **is_correct** (boolean, required)

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
        "error": "test not found"
    }
    ```
    ```json
    {
        "error": "question not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Get Question
**Path:** `host/api/v1/courses/:course_id/tests/:test_id`  
**Method:** `GET`  
**Description:** Get question details.

### Query Parameters
- **append_with** (string, optional): Include some infos associated to the question. Possible values: `test`.

### Responses
- **200 OK**
    ```json
    {
        "id": 2,
        "content": "What is the organization that created Golang?",
        "description": "it is not that hard",
        "duration": 30,
        "test_id": 2,
        "test": {
            "id": 2,
            "max_chances": 3,
            "chapter_id": 24
        },
        "options": [
            {
                "id": 5,
                "content": "google",
                "is_correct": true,
                "question_id": 2
            },
            {
                "id": 6,
                "content": "google",
                "is_correct": false,
                "question_id": 2
            },
            {
                "id": 7,
                "content": "twitter",
                "is_correct": false,
                "question_id": 2
            },
            {
                "id": 8,
                "content": "microsoft",
                "is_correct": false,
                "question_id": 2
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
        "error": "test not found"
    }
    ```
    ```json
    {
        "error": "question not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Delete Question
**Path:** `host/api/v1/courses/:course_id/chapters/:chapter_id/tests/:test_id/questions/:question_id`  
**Method:** `DELETE`  
**Description:** Delete a question.

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
        "error": "user is not an author"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "test not found"
    }
    ```
    ```json
    {
        "error": "question not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```