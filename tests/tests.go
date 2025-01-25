package tests

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

	// Test case: Missing required fields
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
