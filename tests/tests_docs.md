# Tests Documentation

This document provides detailed information and code snippets for the `tests` package. The package includes tests for course creation, user authentication, and user authorization.

## Package Imports

```go
import (
    "fmt"
    "math/rand"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/middlewares"
    "github.com/OucheneMohamedNourElIslem658/learn_oo/shared/utils"
    "github.com/gin-gonic/gin"
    "github.com/go-resty/resty/v2"
    "github.com/stretchr/testify/assert"
)
```

## TestCreateCourseEndpoint

This function tests the course creation endpoint.

### Test Cases

1. **create-course-success**: Tests successful course creation.
2. **create-course-bad-request**: Tests course creation with missing required fields.

```go
func TestCreateCourseEndpoint(t *testing.T) {
    fmt.Println("Testing course creation :")
    client := resty.New()

    baseURL := "https://learn-oo-api.onrender.com/api/v1"

    idToken := func() string {
        userID := "92d856d0-a0ee-4220-81b4-66e8adc073fc"
        authorID := "72125ca0-94f3-42e5-ac3a-0d797ec9078b"

        token, _ := utils.CreateIdToken(
            userID,
            &authorID,
            true,
        )

        return token
    }()

    t.Run("create-course-success", func(t *testing.T) {
        resp, err := client.R().
            SetHeader("Authorization", fmt.Sprintf("Bearer %v", idToken)).
            SetMultipartFormData(map[string]string{
                "title":       "Counting",
                "description": "Learn How To Count",
                "price":       "2000",
                "language":    "en",
                "level":       "bigener",
                "duration":    "5",
            }).
            SetFile("video", "./assets/123.mp4").
            SetFile("image", "./assets/123.jpg").
            Post(baseURL + "/courses/")

        assert.NoError(t, err)
        assert.Equal(t, 201, resp.StatusCode())
    })

    t.Run("create-course-bad-request", func(t *testing.T) {
        resp, err := client.R().
            SetHeader("Authorization", fmt.Sprintf("Bearer %v", idToken)).
            SetHeader("Content-Type", "application/json").
            SetMultipartFormData(map[string]string{
                "title":       "",
                "description": "Learn How To Count",
                "price":       "5000",
                "language":    "sp",
                "level":       "beg",
                "duration":    "2",
            }).
            SetFile("image", "./assets/123.txt").
            Post(baseURL + "/courses/")

        assert.NoError(t, err)
        assert.Equal(t, 400, resp.StatusCode())
        assert.Contains(t, resp.String(), "'ar' 'fr' 'en'")
        assert.Contains(t, resp.String(), "'bigener' 'medium' 'advanced'")
        assert.Contains(t, resp.String(), "at least: 5")
        assert.Contains(t, resp.String(), "video")
        assert.Contains(t, resp.String(), "required")
    })
}
```

## TestUserAuthentification

This function tests user authentication endpoints.

### Test Cases

1. **user-creation-success**: Tests successful user creation.
2. **user-creation-bad-request**: Tests user creation with missing required fields.
3. **user-login-success**: Tests successful user login.

```go
func TestUserAuthentification(t *testing.T) {
    fmt.Println("Testing user authentification :")
    client := resty.New()

    baseURL := "https://learn-oo-api.onrender.com/api/v1"

    t.Run("user-creation-success", func(t *testing.T) {
        randomEmail := func() string {
            rand.Seed(time.Now().UnixNano())
            randomInt := rand.Int()
            fmt.Println("Random Integer:", randomInt)
            return fmt.Sprintf("test%v@gmail.com", randomInt)
        }()

        resp, err := client.R().
            SetHeader("Content-Type", "application/json").
            SetBody(map[string]string{
                "email":     randomEmail,
                "password":  "123456",
                "full_name": "test",
            }).Post(baseURL + "/users/auth/register-with-email-and-password")

        fmt.Println(resp.String())

        assert.NoError(t, err)
        assert.Equal(t, 201, resp.StatusCode())
    })

    t.Run("user-creation-bad-request", func(t *testing.T) {
        resp, err := client.R().
            SetHeader("Content-Type", "application/json").
            SetBody(map[string]string{
                "email":     "m_ouchene@estin.dz",
                "password":  "",
                "full_name": "test",
            }).Post(baseURL + "/users/auth/register-with-email-and-password")

        assert.NoError(t, err)
        assert.Equal(t, 400, resp.StatusCode())
    })

    t.Run("user-login-success", func(t *testing.T) {
        resp, err := client.R().
            SetHeader("Content-Type", "application/json").
            SetBody(map[string]string{
                "email":    "m_ouchene@estin.dz",
                "password": "123456",
            }).Post(baseURL + "/users/auth/login-with-email-and-password")

        assert.NoError(t, err)
        assert.Equal(t, 200, resp.StatusCode())
    })
}
```

## TestUsersAuthorization

This function tests user authorization middleware.

### Test Case

1. **Authorization Middleware**: Tests the authorization middleware with ID and author checks.

```go
func TestUsersAuthorization(t *testing.T) {
    authorizationMiddlewares := middlewares.NewAuthorizationMiddlewares()

    authorization := authorizationMiddlewares.Authorization()
    authorizationWithIDCheck := authorizationMiddlewares.AuthorizationWithIDCheck()
    authorizationWithAuthorCheck := authorizationMiddlewares.AuthorizationWithAuthorCheck()

    randomIDToken := func() string {
        userID := "92d856d0-a0ee-4220-81b4-66e8adc073fc"
        authorID := "72125ca0-94f3-42e5-ac3a-0d797ec9078b"

        token, _ := utils.CreateIdToken(
            userID,
            &authorID,
            true,
        )

        return token
    }()

    fmt.Println(randomIDToken)

    router := gin.New()
    router.GET("/",
        authorization,
        authorizationWithIDCheck,
        authorizationWithAuthorCheck,
        func(ctx *gin.Context) {
            ctx.Status(200)
        },
    )

    w := httptest.NewRecorder()

    req, _ := http.NewRequest("GET", "/", nil)
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", randomIDToken))
    router.ServeHTTP(w, req)

    fmt.Println(w.Body.String())

    assert.Equal(t, 200, w.Code)
}
```
## cURL Commands

For non-Go developers, here are the equivalent cURL commands for each test case.

### TestCreateCourseEndpoint

1. **create-course-success**:

```sh
curl -X POST "https://learn-oo-api.onrender.com/api/v1/courses/" \
-H "Authorization: Bearer <idToken>" \
-F "title=Counting" \
-F "description=Learn How To Count" \
-F "price=2000" \
-F "language=en" \
-F "level=bigener" \
-F "duration=5" \
-F "video=@./assets/123.mp4" \
-F "image=@./assets/123.jpg"
```

2. **create-course-bad-request**:

```sh
curl -X POST "https://learn-oo-api.onrender.com/api/v1/courses/" \
-H "Authorization: Bearer <idToken>" \
-H "Content-Type: application/json" \
-F "title=" \
-F "description=Learn How To Count" \
-F "price=5000" \
-F "language=sp" \
-F "level=beg" \
-F "duration=2" \
-F "image=@./assets/123.txt"
```

### TestUserAuthentification

1. **user-creation-success**:

```sh
curl -X POST "https://learn-oo-api.onrender.com/api/v1/users/auth/register-with-email-and-password" \
-H "Content-Type: application/json" \
-d '{
    "email": "test<randomInt>@gmail.com",
    "password": "123456",
    "full_name": "test"
}'
```

2. **user-creation-bad-request**:

```sh
curl -X POST "https://learn-oo-api.onrender.com/api/v1/users/auth/register-with-email-and-password" \
-H "Content-Type: application/json" \
-d '{
    "email": "m_ouchene@estin.dz",
    "password": "",
    "full_name": "test"
}'
```

3. **user-login-success**:

```sh
curl -X POST "https://learn-oo-api.onrender.com/api/v1/users/auth/login-with-email-and-password" \
-H "Content-Type: application/json" \
-d '{
    "email": "m_ouchene@estin.dz",
    "password": "123456"
}'
```

### TestUsersAuthorization

1. **Authorization Middleware**:

```sh
curl -X GET "http://localhost:8080/" \
-H "Authorization: Bearer <randomIDToken>"
```