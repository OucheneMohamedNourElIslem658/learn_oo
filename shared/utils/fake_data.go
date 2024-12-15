package utils

import (
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/database"
	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

func GenerateExampleData() {
	courses := []models.Course{
		// {
		// 	Title:       "Introduction to Go",
		// 	Description: "A beginner's guide to learning Go programming.",
		// 	Price:       29.99,
		// 	Language:    "en",
		// 	Level:       "advanced",
		// 	Duration:    120,
		// 	IsCompleted: true,
		// 	Objectives: []models.Objective{
		// 		{Content: "Understand variables, loops, and basic syntax."},
		// 		{Content: "Learn about goroutines and channels."},
		// 	},
		// 	Requirements: []models.Requirement{
		// 		{Content: "Basic understanding of programming concepts."},
		// 		{Content: "Familiarity with any programming language is helpful."},
		// 	},			
		// 	AuthorID:    "72125ca0-94f3-42e5-ac3a-0d797ec9078b",
		// 	Video: &models.File{
		// 		URL: "https://ik.imagekit.io/cdejmhtxd/files/authors/Go%20in%20100%20Seconds%20(1).mp4?updatedAt=1734254395913",
		// 		Height: 720,
		// 		Width: 1280,
		// 	},
		// 	Image: &models.File{
		// 		URL: "data:image/jpeg;base64,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",
		// 		Height: 320,
		// 		Width: 630,
		// 	},
		// 	Chapters: []models.Chapter{
		// 		{
		// 			Title:       "Getting Started",
		// 			Description: "Introduction to the Go language and setting up your development environment.",
		// 			Lessons: []models.Lesson{
		// 				{
		// 					Title:       "Installing Go",
		// 					Description: "Learn how to install Go on your system.",
		// 					IsVideo:     false,
		// 					Content:     gin.H{
		// 						"ops": []gin.H{
		// 							{
		// 								"insert": "Follow the instructions on the Go website.\n",
		// 							},
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Title:       "Hello World in Go",
		// 					Description: "Write and run your first Go program.",
		// 					IsVideo:     true,
		// 					Video: &models.File{
		// 						URL: "https://ik.imagekit.io/cdejmhtxd/files/Welcome%20to%20series%20on%20GO%20programming%20language.mp4?updatedAt=1734255071005",
		// 						Height: 720,
		// 						Width: 1280,
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Title:       "Go Basics",
		// 			Description: "Learn the basic syntax and data structures of Go.",
		// 			Lessons: []models.Lesson{
		// 				{
		// 					Title:       "Variables and Constants",
		// 					Description: "Learn how to declare variables and constants in Go.",
		// 					IsVideo:     false,
		// 					Content: gin.H{
		// 						"ops": []gin.H{
		// 							{
		// 								"insert": "Variables and constants are fundamental in Go.\n",
		// 							},
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Title:       "Control Flow",
		// 					Description: "Understand conditional statements and loops in Go.",
		// 					IsVideo:     true,
		// 					Content: gin.H{
		// 						"ops": []gin.H{
		// 							{
		// 								"insert": "Control flow fundamental in Go.\n",
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		{
			Title:       "Introduction to Cooking",
			Description: "A beginner's guide to mastering the art of cooking.",
			Price:       19.99,
			Language:    "en",
			Level:       "beginner",
			Duration:    90,
			IsCompleted: false,
			AuthorID:    "72125ca0-94f3-42e5-ac3a-0d797ec9078b",
			Categories: []models.Category{
				{Name: "cooking"},
				{Name: "food"},
			},
			Objectives: []models.Objective{
				{Content: "Learn basic cooking techniques."},
				{Content: "Understand kitchen safety and hygiene."},
			},
			Requirements: []models.Requirement{
				{Content: "No prior cooking experience required."},
				{Content: "A passion for food and cooking."},
			},
			Video: &models.File{
				URL: "https://ik.imagekit.io/cdejmhtxd/files/How%20to%20make%20Miniature%20Spaghetti%20Meatballs%20.mp4?updatedAt=1734257077649",
				Height: 720,
				Width: 1280,
			},
			Image: &models.File{
				URL: "https://media-cldnry.s-nbcnews.com/image/upload/t_nbcnews-fp-1200-630,f_auto,q_auto:best/newscms/2019_41/3044956/191009-cooking-vegetables-al-1422.jpg",
				Height: 630,
				Width: 1200,
			},
			Chapters: []models.Chapter{
				{
					Title:       "Getting Started",
					Description: "Introduction to cooking and kitchen setup.",
					Lessons: []models.Lesson{
						{
							Title:       "Kitchen Essentials",
							Description: "Learn about essential kitchen tools and equipment.",
							IsVideo:     true,
							Content:     gin.H{
								"ops": []gin.H{
									{
										"insert": "Familiarize yourself with the tools you'll need.\n",
									},
								},
							},
						},
						{
							Title:       "Basic Knife Skills",
							Description: "Learn how to properly use a knife in the kitchen.",
							IsVideo:     true,
							Video: &models.File{
								URL: "https://dummyurl.com/basic_knife_skills_video.mp4",
								Height: 720,
								Width: 1280,
							},
						},
					},
				},
				{
					Title:       "Cooking Techniques",
					Description: "Explore various cooking methods and techniques.",
					Lessons: []models.Lesson{
						{
							Title:       "Boiling and Steaming",
							Description: "Learn how to boil and steam vegetables.",
							IsVideo:     false,
							Content: gin.H{
								"ops": []gin.H{
									{
										"insert": "Boiling and steaming are fundamental cooking methods.\n",
									},
								},
							},
						},
						{
							Title:       "Sautéing and Stir-Frying",
							Description: "Understand the techniques of sautéing and stir-frying.",
							IsVideo:     true,
							Content: gin.H{
								"ops": []gin.H{
									{
										"insert": "Sautéing and stir-frying are quick cooking methods.\n",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Title:       "Introduction to Acting",
			Description: "A beginner's guide to the fundamentals of acting.",
			Price:       39.99,
			Language:    "en",
			Level:       "beginner",
			Duration:    120,
			IsCompleted: true,
			Categories: []models.Category{
				{Name: "Acting"},
				{Name: "Art"},
			},
			Objectives: []models.Objective{
				{Content: "Understand the basics of character development."},
				{Content: "Learn essential acting techniques and exercises."},
			},
			Requirements: []models.Requirement{
				{Content: "No prior acting experience required."},
				{Content: "A willingness to perform and express yourself."},
			},
			AuthorID:    "72125ca0-94f3-42e5-ac3a-0d797ec9078b",
			Video: &models.File{
				URL: "https://ik.imagekit.io/cdejmhtxd/files/What%20is%20Acting_%20__%20The%20Art%20of%20Acting.mp4?updatedAt=1734257479478",
				Height: 360,
				Width: 640,
			},
			Image: &models.File{
				URL: "https://miro.medium.com/v2/resize:fit:1400/0*lXgMeXvD39oe65AR",
				Height: 1334,
				Width: 768,
			},
			Chapters: []models.Chapter{
				{
					Title:       "Getting Started",
					Description: "Introduction to acting and the world of performance.",
					Lessons: []models.Lesson{
						{
							Title:       "Understanding Acting",
							Description: "Learn what acting is and its importance in storytelling.",
							IsVideo:     false,
							Content:     gin.H{
								"ops": []gin.H{
									{
										"insert": "Acting is about bringing characters to life.\n",
									},
								},
							},
						},
						{
							Title:       "Warm-Up Exercises",
							Description: "Discover essential warm-up exercises for actors.",
							IsVideo:     true,
							Video: &models.File{
								URL: "https://ik.imagekit.io/cdejmhtxd/files/What%20is%20Acting_%20__%20The%20Art%20of%20Acting.mp4?updatedAt=1734257479478",
								Height: 360,
				                Width: 640,
							},
						},
					},
				},
				{
					Title:       "Acting Techniques",
					Description: "Explore various techniques used in acting.",
					Lessons: []models.Lesson{
						{
							Title:       "Method Acting",
							Description: "Learn about the method acting technique and its applications.",
							IsVideo:     false,
							Content: gin.H{
								"ops": []gin.H{
									{
										"insert": "Method acting focuses on emotional authenticity.\n",
									},
								},
							},
						},
						{
							Title:       "Improvisation",
							Description: "Understand the basics of improvisational acting.",
							IsVideo:     true,
							Content: gin.H{
								"ops": []gin.H{
									{
										"insert": "Improvisation helps actors think on their feet.\n",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	
	// categories := []models.CourseCategory{
	// 	{
	// 		CourseID: 18,
	// 		CategoryID: 3,
	// 	},
	// 	{
	// 		CourseID: 18,
	// 		CategoryID: 4,
	// 	},
	// 	{
	// 		CourseID: 18,
	// 		CategoryID: 5,
	// 	},
	// }

	database.Instance.Create(&courses)
}