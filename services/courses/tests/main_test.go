package main

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateCourseEndpoint(t *testing.T) {
	fmt.Println("Testing course creation :")
	client := resty.New()

	baseURL := "https://learn-oo-api.onrender.com/api/v1"

	idToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JfaWQiOiI3MjEyNWNhMC05NGYzLTQyZTUtYWMzYS0wZDc5N2VjOTA3OGIiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZXhwIjoxNzM0NzIxNDI5LCJpZCI6IjkyZDg1NmQwLWEwZWUtNDIyMC04MWI0LTY2ZThhZGMwNzNmYyJ9.cYkOAhyYW5w9DSGHr5yxtpk06eePIpsoZs9pc2llTFk"

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

		fmt.Println(resp.String())

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

	// Test case: Unauthorized (missing token)
	t.Run("create-course-unauthorized", func(t *testing.T) {
		idToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JfaWQiOiJhYmQ3YTI1NS1jOWJhLTRiNDgtYWI4My01MzBiY2EyMmZjZTUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZXhwIjoxNzMzMzgzNTMyLCJpZCI6IjQ0MjE5OTk5LWE5NTEtNDI2Ny05MDI0LWIxMTQ3YTk3NDkwNiJ9.B5FlLPOcZBquK0HtDDJyKpY1vHqLWmsNnD6lRK40Zi8"
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %v", idToken)).
			SetMultipartFormData(map[string]string{
				"title":       "Counting",
				"description": "Learn How To Count",
				"price":       "2000",
				"language":    "en",
				"level":       "beginner",
				"duration":    "5",
			}).
			SetFile("video", "./assets/123.mp4").
			SetFile("image", "./assets/123.jpg").
			Post(baseURL + "/courses/")

		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
		assert.Contains(t, resp.String(), "invalid id token")
	})
}
