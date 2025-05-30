# API Endpoints Documentation

## Get User Profile
**Path:** `host/api/v1/users/profiles/profile`  
**Method:** `GET`  
**Description:** Get user profile.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Query Parameters
- **append_with** (string, optional): Possible values: `image,author_profile`. Include additional information associated with the user.

### Responses
- **200 OK**
    ```json
    {
        "id": "61f244aa-a912-4fd1-98ac-b1df7c59a484",
        "created_at": "2024-11-22T00:59:56.311557+08:00",
        "updated_at": "2024-11-22T01:53:47.935183+08:00",
        "deleted_at": null,
        "email": "mohamedouchene996@gmail.com",
        "password": "$2a$10$oSAXVysFdiiqlzyktuTii.9IzhdWeSwwKAHiWjUSe8EtnwIAqHIaO",
        "full_name": "ouchene mohamed nour el islam",
        "email_verified": true,
        "image": {
            "id": 7,
            "created_at": "2024-11-22T22:08:42.484775+08:00",
            "updated_at": "2024-11-22T22:08:42.484775+08:00",
            "deleted_at": null,
            "url": "https://ik.imagekit.io/cdejmhtxd/learn_oo/images/users/61f244aa-a912-4fd1-98ac-b1df7c59a484/image_B6s3pGhUm",
            "thumbnail_url": "",
            "image_kit_id": "67409063e375273f6000a1f8",
            "user_id": "61f244aa-a912-4fd1-98ac-b1df7c59a484"
        },
        "author_profile": {
            "id": "132e85c5-b0de-4516-997d-48ad58ce4583",
            "created_at": "2024-11-22T01:55:35.866383+08:00",
            "updated_at": "2024-11-22T01:55:35.866383+08:00",
            "deleted_at": null,
            "user_id": "61f244aa-a912-4fd1-98ac-b1df7c59a484"
        }
    }
    ```
- **401 Unauthorized**
    ```json
    {
        "error": "id token expired"
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "user not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update User Profile
**Path:** `host/api/v1/users/profiles/profile`  
**Method:** `PUT`  
**Description:** Update user profile.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **full_name** (string, required): The length needs to be greater than 0.

### Responses
- **200 OK**
- **401 Unauthorized**
    ```json
    {
        "error": "id token expired"
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "invalid id token"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "user not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update User Profile Image
**Path:** `host/api/v1/users/profiles/profile/image`  
**Method:** `PUT`  
**Description:** Update user profile image.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **image** (file, required): Profile image as file.

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "error": "multipart request error"
    }
    ```
    ```json
    {
        "error": "the file is not an image"
    }
    ```
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

## Add Author Profile to the Current User
**Path:** `host/api/v1/users/authors/upgrade`  
**Method:** `PUT`  
**Description:** Add author profile to the current user.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "error": "undefined authorization"
    }
    ```
    ```json
    {
        "error": "user is already an author"
    }
    ```
- **404 Not Found**
    ```json
    {
        "error": "user not found"
    }
    ```
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

## Soft Delete Author Profile of the Current User
**Path:** `host/api/v1/users/authors/downgrade`  
**Method:** `DELETE`  
**Description:** Soft delete author profile of the current user.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Responses
- **200 OK**
- **400 Bad Request**
    ```json
    {
        "error": "undefined authorization"
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

## Get Author Profile
**Path:** `host/api/v1/users/profiles/authors/profile`  
**Method:** `GET`  
**Description:** Get author profile.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Query Parameters
- **append_with** (string, optional): Possible values: `user,accomplishments`. Include additional information associated with the author.

### Responses
- **200 OK**
    ```json
    {
        "id": "132e85c5-b0de-4516-997d-48ad58ce4583",
        "created_at": "2024-11-22T01:55:35.866383+08:00",
        "updated_at": "2024-11-22T01:55:35.866383+08:00",
        "deleted_at": null,
        "user_id": "61f244aa-a912-4fd1-98ac-b1df7c59a484",
        "user": {
            "id": "61f244aa-a912-4fd1-98ac-b1df7c59a484",
            "created_at": "2024-11-22T00:59:56.311557+08:00",
            "updated_at": "2024-11-22T01:53:47.935183+08:00",
            "deleted_at": null,
            "email": "mohamedouchene996@gmail.com",
            "password": "$2a$10$oSAXVysFdiiqlzyktuTii.9IzhdWeSwwKAHiWjUSe8EtnwIAqHIaO",
            "full_name": "ouchene mohamed nour el islam",
            "email_verified": true
        },
        "accomplishments": [
            {
                "id": 8,
                "created_at": "2024-11-23T15:10:59.355227+08:00",
                "updated_at": "2024-11-23T15:10:59.355227+08:00",
                "deleted_at": null,
                "url": "https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/accomplisments/132e85c5-b0de-4516-997d-48ad58ce4583/image_x0UhcF0v8",
                "thumbnail_url": "",
                "image_kit_id": "67418002e375273f60a9884c",
                "author_id": "132e85c5-b0de-4516-997d-48ad58ce4583"
            }
        ]
    }
    ```
- **400 Bad Request**
    ```json
    {
        "error": "undefined authorization"
    }
    ```
    ```json
    {
        "error": "user is not an author"
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
        "error": "author not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Update Author Profile
**Path:** `host/api/v1/users/profiles/authors/profile`  
**Method:** `PUT`  
**Description:** Update author profile.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **bio** (json, required): Needs to be in JSON format (description of the rich text).

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
        "error": "user is not an author"
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
        "error": "author not found"
    }
    ```
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## Add Accomplishments to the Author Profile
**Path:** `host/api/v1/users/profiles/author/profile/accomplishments`  
**Method:** `POST`  
**Description:** Add accomplishments to the author profile.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Request Body
- **accomplishments** (list of files, required): List of files as accomplishments for this author.

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
        "error": "user is not an author"
    }
    ```
    ```json
    {
        "error": "multipart request error"
    }
    ```
    ```json
    {
        "error": "accomplishments list cannot be empty"
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

## Delete an Author Accomplishment by ID
**Path:** `host/api/v1/users/profiles/author/profile/accomplishments/:id`  
**Method:** `DELETE`  
**Description:** Delete an author accomplishment by ID.

### Headers
- **Authorization** (string, required): Bearer id_token.

### Path Parameters
- **id** (string, required): ID of the accomplishment to delete.

### Responses
- **200 OK**
- **404 Not Found**
    ```json
    {
        "error": "file not found"
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
- **500 Internal Server Error**
    ```json
    {
        "error": "error message (contact me when you see one)"
    }
    ```

## gRPC Service Documentation

### Service: `ProfilesService`

#### RPC Methods

- **GetProfile**
    - **Request:** `google.protobuf.Empty`
    - **Response:** [`Profile`](#profile)
    - **Description:** Retrieves the current user's profile.

- **UpgradeToAuthor**
    - **Request:** `google.protobuf.Empty`
    - **Response:** `google.protobuf.Empty`
    - **Description:** Upgrades the current user to an author.

- **DowngradeToUser**
    - **Request:** `google.protobuf.Empty`
    - **Response:** `google.protobuf.Empty`
    - **Description:** Downgrades the current author to a regular user.

- **GetAuthor**
    - **Request:** [`GetAuthorRequest`](#getauthorrequest)
    - **Response:** [`Author`](#author)
    - **Description:** Retrieves an author profile by ID.

---

### Messages

#### Profile

| Field           | Type      | Description                                      |
|-----------------|-----------|--------------------------------------------------|
| id              | string    | Unique identifier for the user                   |
| full_name       | string    | Full name of the user                            |
| email           | string    | Email address                                    |
| email_verified  | bool      | Whether the email is verified                    |
| image           | File      | (optional) Profile image                         |
| author_profile  | Author    | (optional) Author profile if user is an author   |
| courses         | Course[]  | List of courses associated with the user         |

#### Author

| Field           | Type      | Description                                      |
|-----------------|-----------|--------------------------------------------------|
| id              | string    | Unique identifier for the author                 |
| bio             | string    | Author's biography                              |
| balance         | int32     | Author's balance                                |
| user_profile    | Profile   | Linked user profile                             |
| accomplishments | File[]    | List of accomplishment files                    |

#### Course

| Field             | Type      | Description                                    |
|-------------------|-----------|------------------------------------------------|
| id                | uint64    | Course ID                                      |
| title             | string    | Course title                                   |
| description       | string    | Course description                             |
| price             | double    | Course price                                   |
| payment_price_id  | string    | Payment price ID                               |
| payment_product_id| string    | Payment product ID                             |
| language          | string    | Course language                                |
| level             | string    | Course level                                   |
| duration          | uint64    | Duration (in minutes or seconds)               |
| rate              | double    | Average rating                                 |
| raters_count      | uint64    | Number of raters                               |
| is_completed      | bool      | Completion status                              |
| video             | File      | Course video file                              |
| image             | File      | Course image file                              |
| author_id         | string    | Author ID                                      |
| author            | Author    | Author details                                 |

#### File

| Field         | Type      | Description                                      |
|---------------|-----------|--------------------------------------------------|
| id            | uint64    | File ID                                          |
| url           | string    | File URL                                         |
| height        | int32     | Image height (if applicable)                     |
| width         | int32     | Image width (if applicable)                      |
| thumbnail_url | string    | (optional) Thumbnail URL                         |

#### GetAuthorRequest

| Field | Type   | Description                |
|-------|--------|----------------------------|
| id    | string | Author ID to retrieve      |

---