# Database Tables and Attributes

## Users
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | text         | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| email              | text         | NOT NULL, UNIQUE          |
| password           | text         |                           |
| full_name          | text         | NOT NULL                  |
| email_verified     | boolean      |                           |
| payment_customer_id| text         |                           |

## Authors
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | text         | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| bio                | text         |                           |
| user_id            | text         | FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Courses
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| title              | text         | NOT NULL                  |
| description        | text         |                           |
| price              | decimal      |                           |
| payment_price_id   | text         | UNIQUE                    |
| payment_product_id | text         | UNIQUE                    |
| language           | text         | DEFAULT 'en'              |
| level              | text         | DEFAULT 'beginner'        |
| duration           | bigint       |                           |
| is_completed       | boolean      |                           |
| author_id          | text         | FOREIGN KEY REFERENCES authors(id) ON DELETE SET NULL ON UPDATE CASCADE |

## Chapters
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| title              | text         |                           |
| description        | text         |                           |
| course_id          | bigint       | FOREIGN KEY REFERENCES courses(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Tests
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| max_chances        | bigint       | DEFAULT 1                 |
| chapter_id         | bigint       | FOREIGN KEY REFERENCES chapters(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Test Results
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| test_id            | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES tests(id) ON DELETE CASCADE ON UPDATE CASCADE |
| learner_id         | varchar(64)  | PRIMARY KEY, FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE |
| has_succeed        | boolean      |                           |
| current_chance     | bigint       |                           |

## Lessons
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| order              | bigint       |                           |
| title              | text         | NOT NULL                  |
| description        | text         |                           |
| content            | text         |                           |
| chapter_id         | bigint       | FOREIGN KEY REFERENCES chapters(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Lesson Learners
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| lesson_id          | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES lessons(id) ON DELETE CASCADE ON UPDATE CASCADE |
| learner_id         | varchar(64)  | PRIMARY KEY, FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE |
| learned            | boolean      |                           |

## Course Learners
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| course_id          | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES courses(id) ON DELETE CASCADE ON UPDATE CASCADE |
| learner_id         | text         | PRIMARY KEY, FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE |
| leaning_status     | text         | DEFAULT 'learning'        |
| rate               | decimal      |                           |
| checkout_id        | text         |                           |

## Questions
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| content            | text         |                           |
| description        | text         |                           |
| duration           | bigint       |                           |
| deleted_at         | timestamptz  |                           |
| test_id            | bigint       | FOREIGN KEY REFERENCES tests(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Question Answers
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| question_id        | bigint       | FOREIGN KEY REFERENCES questions(id) ON DELETE CASCADE ON UPDATE CASCADE |
| learner_id         | varchar(64)  | FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Categories
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| name               | text         | NOT NULL, UNIQUE          |

## Course Categories
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| course_id          | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES courses(id) ON DELETE CASCADE ON UPDATE CASCADE |
| category_id        | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES categories(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Objectives
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| content            | text         |                           |
| course_id          | bigint       | FOREIGN KEY REFERENCES courses(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Requirements
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| content            | text         |                           |
| course_id          | bigint       | FOREIGN KEY REFERENCES courses(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Options
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| content            | text         |                           |
| is_correct         | boolean      |                           |
| deleted_at         | timestamptz  |                           |
| question_id        | bigint       | FOREIGN KEY REFERENCES questions(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Option Selections
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| option_id          | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES options(id) ON DELETE CASCADE ON UPDATE CASCADE |
| question_answer_id | bigint       | PRIMARY KEY, FOREIGN KEY REFERENCES question_answers(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Files
| Attribute          | Type         | Constraints               |
|--------------------|--------------|---------------------------|
| id                 | bigserial    | PRIMARY KEY               |
| created_at         | timestamptz  |                           |
| updated_at         | timestamptz  |                           |
| deleted_at         | timestamptz  |                           |
| key                | text         | NOT NULL                  |
| url                | text         | NOT NULL                  |
| course_id          | bigint       | FOREIGN KEY REFERENCES courses(id) ON DELETE CASCADE ON UPDATE CASCADE |

## Indexes
- `idx_users_deleted_at` on `users(deleted_at)`
- `idx_authors_deleted_at` on `authors(deleted_at)`
- `idx_courses_deleted_at` on `courses(deleted_at)`
- `idx_chapters_deleted_at` on `chapters(deleted_at)`
- `idx_tests_deleted_at` on `tests(deleted_at)`
- `idx_test_results_deleted_at` on `test_results(deleted_at)`
- `idx_lessons_deleted_at` on `lessons(deleted_at)`
- `idx_lesson_learners_deleted_at` on `lesson_learners(deleted_at)`
- `idx_course_learners_deleted_at` on `course_learners(deleted_at)`
- `idx_questions_deleted_at` on `questions(deleted_at)`
- `idx_question_answers_deleted_at` on `question_answers(deleted_at)`
- `idx_categories_deleted_at` on `categories(deleted_at)`
- `idx_course_categories_deleted_at` on `course_categories(deleted_at)`
- `idx_objectives_deleted_at` on `objectives(deleted_at)`
- `idx_requirements_deleted_at` on `requirements(deleted_at)`
- `idx_options_deleted_at` on `options(deleted_at)`
- `idx_option_selections_deleted_at` on `option_selections(deleted_at)`
- `idx_files_deleted_at` on `files(deleted_at)`
- `idx_comments_deleted_at` on `comments(deleted_at)`
- `idx_notifications_deleted_at` on `notifications(deleted_at)`