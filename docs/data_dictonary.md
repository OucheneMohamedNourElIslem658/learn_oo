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

## Relations

### User - Author
**Relation Name:** user_author  
**Description:** One-to-One relationship between Users and Authors.  

### Author - Course
**Relation Name:** author_course  
**Description:** One-to-Many relationship between Authors and Courses.  

### Course - Chapter
**Relation Name:** course_chapter  
**Description:** One-to-Many relationship between Courses and Chapters.  

### Chapter - Lesson
**Relation Name:** chapter_lesson  
**Description:** One-to-Many relationship between Chapters and Lessons.  

### Chapter - Test
**Relation Name:** chapter_test  
**Description:** One-to-One relationship between Chapters and Tests.  

### Test - Question
**Relation Name:** test_question  
**Description:** One-to-Many relationship between Tests and Questions.  

### Question - Option
**Relation Name:** question_option  
**Description:** One-to-Many relationship between Questions and Options.  

### Question - Question Answer
**Relation Name:** question_question_answer  
**Description:** One-to-Many relationship between Questions and Question Answers.  

### Question Answer - Option Selection
**Relation Name:** question_answer_option_selection  
**Description:** One-to-Many relationship between Question Answers and Option Selections.  

### Course - Objective
**Relation Name:** course_objective  
**Description:** One-to-Many relationship between Courses and Objectives.  

### Course - Requirement
**Relation Name:** course_requirement  
**Description:** One-to-Many relationship between Courses and Requirements.  

### Course - File (Image)
**Relation Name:** course_image  
**Description:** One-to-One relationship between Courses and Files.  

### Course - File (Video)
**Relation Name:** course_video
**Description:** One-to-One relationship between Courses and Files.  

### Course - Category
**Relation Name:** course_category  
**Description:** Many-to-Many relationship between Courses and Categories.  
**Attributes:**
- `course_id` (bigint): PRIMARY KEY, FOREIGN KEY REFERENCES courses(id)
- `category_id` (bigint): PRIMARY KEY, FOREIGN KEY REFERENCES categories(id)

### Course - Learner
**Relation Name:** course_learner  
**Description:** Many-to-Many relationship between Courses and Learners.  
**Attributes:**
- `course_id` (bigint): PRIMARY KEY, FOREIGN KEY REFERENCES courses(id)
- `learner_id` (text): PRIMARY KEY, FOREIGN KEY REFERENCES users(id)

### Lesson - Learner
**Relation Name:** lesson_learner  
**Description:** Many-to-Many relationship between Lessons and Learners.  
**Attributes:**
- `lesson_id` (bigint): PRIMARY KEY, FOREIGN KEY REFERENCES lessons(id)
- `learner_id` (varchar(64)): PRIMARY KEY, FOREIGN KEY REFERENCES users(id)

### Test - Learner
**Relation Name:** test_learner  
**Description:** Many-to-Many relationship between Tests and Learners.  
**Attributes:**
- `test_id` (bigint): PRIMARY KEY, FOREIGN KEY REFERENCES tests(id)
- `learner_id` (varchar(64)): PRIMARY KEY, FOREIGN KEY REFERENCES users(id)
- `has_succeed` (boolean)
- `current_chance` (bigint)
- `created_at` (timestamptz)
- `updated_at` (timestamptz)
- `deleted_at` (timestamptz)

## Business Rules and Validations

### Users
- **Email:** Must be a valid email format.
- **Password:** Must be at least 8 characters long.
- **Full Name:** Cannot be empty.

### Courses
- **Price:** Must be a positive value.
- **Language:** Must be one of the predefined languages (e.g., 'en', 'es', 'fr').
- **Level:** Must be one of the predefined levels (e.g., 'beginner', 'intermediate', 'advanced').

### Lessons
- **Order:** Must be a positive integer.
- **Title:** Cannot be empty.

### Tests
- **Max Chances:** Must be a positive integer.

### Calculated Values
- **Course Rate:** Calculated as the average of all learner rates for the course.
- **Lesson Learned:** Determined if the learner has completed all the lessons in a chapter.
- **Test Passed:** Determined if the learner has passed the test based on the `has_succeed` field.

### Additional Rules
- **User Email Verification:** Users must verify their email before accessing certain features.
- **Course Completion:** A course is marked as completed if all chapters and tests are completed by the learner.
- **Content Uniqueness:** Titles and descriptions for courses, lessons, and chapters should be unique to avoid duplication.
- **Cascade Deletion:** Deleting a user will cascade delete related authors, courses, chapters, lessons, tests, and results.
- **Soft Deletion:** Records are not permanently deleted but marked with a `deleted_at` timestamp for potential recovery.

### Hooks Used
- **Before Save Hooks:** Used to enforce business rules before saving records to the database.
- **After Find Hooks:** Used to calculate and populate derived fields after retrieving records from the database.
- **Custom Validators:** Implemented to ensure data integrity and adherence to business rules.

## SQL Queries

### Table Creation

```sql
CREATE TABLE "users" (
    "id" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "email" text NOT NULL,
    "password" text,
    "full_name" text NOT NULL,
    "email_verified" boolean,
    "payment_customer_id" text,
    PRIMARY KEY ("id"),
    CONSTRAINT "uni_users_email" UNIQUE ("email")
);

CREATE TABLE "authors" (
    "id" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "bio" text,
    "user_id" text,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_users_author_profile" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "courses" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "title" text NOT NULL,
    "description" text,
    "price" decimal,
    "payment_price_id" text,
    "payment_product_id" text,
    "language" text DEFAULT 'en',
    "level" text DEFAULT 'beginner',
    "duration" bigint,
    "is_completed" boolean,
    "author_id" text,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_authors_courses" FOREIGN KEY ("author_id") REFERENCES "authors"("id") ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT "uni_courses_payment_price_id" UNIQUE ("payment_price_id"),
    CONSTRAINT "uni_courses_payment_product_id" UNIQUE ("payment_product_id")
);

CREATE TABLE "chapters" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "title" text,
    "description" text,
    "course_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_courses_chapters" FOREIGN KEY ("course_id") REFERENCES "courses"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "tests" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "max_chances" bigint DEFAULT 1,
    "chapter_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_chapters_test" FOREIGN KEY ("chapter_id") REFERENCES "chapters"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "test_results" (
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "test_id" bigint,
    "learner_id" varchar(64),
    "has_succeed" boolean,
    "current_chance" bigint,
    PRIMARY KEY ("test_id", "learner_id"),
    CONSTRAINT "fk_test_results_test" FOREIGN KEY ("test_id") REFERENCES "tests"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_test_results_learner" FOREIGN KEY ("learner_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "lessons" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "order" bigint,
    "title" text NOT NULL,
    "description" text,
    "content" text,
    "chapter_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_chapters_lessons" FOREIGN KEY ("chapter_id") REFERENCES "chapters"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "lesson_learners" (
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "lesson_id" bigint,
    "learner_id" varchar(64),
    "learned" boolean,
    PRIMARY KEY ("lesson_id", "learner_id"),
    CONSTRAINT "fk_lesson_learners_lesson" FOREIGN KEY ("lesson_id") REFERENCES "lessons"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_lesson_learners_learner" FOREIGN KEY ("learner_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "course_learners" (
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "course_id" bigint,
    "learner_id" text,
    "leaning_status" text DEFAULT 'learning',
    "rate" decimal,
    "checkout_id" text,
    PRIMARY KEY ("course_id", "learner_id"),
    CONSTRAINT "fk_course_learners_course" FOREIGN KEY ("course_id") REFERENCES "courses"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_course_learners_learner" FOREIGN KEY ("learner_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "questions" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "content" text,
    "description" text,
    "duration" bigint,
    "deleted_at" timestamptz,
    "test_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_tests_questions" FOREIGN KEY ("test_id") REFERENCES "tests"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "question_answers" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "question_id" bigint,
    "learner_id" varchar(64),
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_question_answers_question" FOREIGN KEY ("question_id") REFERENCES "questions"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_question_answers_learner" FOREIGN KEY ("learner_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "categories" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "uni_categories_name" UNIQUE ("name")
);

CREATE TABLE "course_categories" (
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "course_id" bigint,
    "category_id" bigint,
    PRIMARY KEY ("course_id", "category_id"),
    CONSTRAINT "fk_course_categories_course" FOREIGN KEY ("course_id") REFERENCES "courses"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_course_categories_category" FOREIGN KEY ("category_id") REFERENCES "categories"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "objectives" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "content" text,
    "course_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_courses_objectives" FOREIGN KEY ("course_id") REFERENCES "courses"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "requirements" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "content" text,
    "course_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_courses_requirements" FOREIGN KEY ("course_id") REFERENCES "courses"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "options" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "content" text,
    "is_correct" boolean,
    "deleted_at" timestamptz,
    "question_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_questions_options" FOREIGN KEY ("question_id") REFERENCES "questions"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "option_selections" (
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "option_id" bigint,
    "question_answer_id" bigint,
    PRIMARY KEY ("option_id", "question_answer_id"),
    CONSTRAINT "fk_option_selections_question_answer" FOREIGN KEY ("question_answer_id") REFERENCES "question_answers"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "fk_option_selections_option" FOREIGN KEY ("option_id") REFERENCES "options"("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "files" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "key" text NOT NULL,
    "url" text NOT NULL,
    "course_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_courses_files" FOREIGN KEY ("course_id") REFERENCES "courses"("id") ON DELETE CASCADE ON UPDATE CASCADE
);
```

### Index Creation

```sql
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_authors_deleted_at" ON "authors" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_courses_deleted_at" ON "courses" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_chapters_deleted_at" ON "chapters" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_tests_deleted_at" ON "tests" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_test_results_deleted_at" ON "test_results" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_lessons_deleted_at" ON "lessons" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_lesson_learners_deleted_at" ON "lesson_learners" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_course_learners_deleted_at" ON "course_learners" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_questions_deleted_at" ON "questions" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_question_answers_deleted_at" ON "question_answers" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_categories_deleted_at" ON "categories" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_course_categories_deleted_at" ON "course_categories" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_objectives_deleted_at" ON "objectives" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_requirements_deleted_at" ON "requirements" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_options_deleted_at" ON "options" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_option_selections_deleted_at" ON "option_selections" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_files_deleted_at" ON "files" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_comments_deleted_at" ON "comments" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_notifications_deleted_at" ON "notifications" ("deleted_at");
```

## Example SQL Queries

### Users
- **Create user with email and password (password is hashed with bcrypt):**
```sql
INSERT INTO "users" ("id", "email", "password", "full_name", "email_verified") 
VALUES ('7308f1cf-2ede-476d-869a-7be31d03afbf', 'ouchene@estin.dz', '$2a$10$FX1XpMqd/GgJ6RWbugxINu.RR/Zfk2b7sZG1UIwnHM7MOGdsoCdla', 'ouchene mohamed', false);
```

- **Upgrade the user to an author:**
```sql
INSERT INTO "authors" ("id", "deleted_at", "bio", "user_id") 
VALUES ('325fba00-f411-43f8-b3c8-0ed2b57f040f', NULL, NULL, '9f69b120-4451-4f14-9705-dd438d882143');
```

- **Downgrade from author to normal user (logically delete the author profile):**
```sql
UPDATE "authors" 
SET "deleted_at" = '2025-01-19 15:35:46.663' 
WHERE "id" = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "deleted_at" IS NULL;
```

### Courses

- **Create course:**

```sql
-- 1 create the course descriptive components:
INSERT INTO "files" ("deleted_at","url","height","width","thumbnail_url","image_kit_id","user_id","author_id","lesson_id","video_course_id","image_course_id") VALUES (NULL,'https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/videos/file_xlkchxPX1',360,640,'',NULL,NULL,NULL,NULL,37,NULL) ON CONFLICT ("id") DO UPDATE SET "video_course_id"="excluded"."video_course_id" RETURNING "id";

INSERT INTO "files" ("deleted_at","url","height","width","thumbnail_url","image_kit_id","user_id","author_id","lesson_id","video_course_id","image_course_id") VALUES (NULL,'https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/images/file_4iESjwQwJ',980,1848,'https://ik.imagekit.io/cdejmhtxd/tr:n-ik_ml_thumbnail/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/images/file_4iESjwQwJ',NULL,NULL,NULL,NULL,NULL,37) ON CONFLICT ("id") DO UPDATE SET "image_course_id"="excluded"."image_course_id" RETURNING "id";

INSERT INTO "courses" ("created_at","updated_at","deleted_at","title","description","price","payment_price_id","payment_product_id","language","level","duration","is_completed","author_id") VALUES ('2025-01-16 20:59:47.027','2025-01-16 20:59:47.027',NULL,'course_title_example','course_description_example',2000,'01jhrb6d9902zr4dad4a1anfhw','01jhrb6bve7jktcrck6d064msh','en','advanced',5,false,'72125ca0-94f3-42e5-ac3a-0d797ec9078b') RETURNING "id";

-- 2 create the course chapters:
SELECT count(*) FROM "courses" WHERE (id = '37' and author_id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b') AND "courses"."deleted_at" IS NULL;

INSERT INTO "chapters" ("created_at","updated_at","deleted_at","title","description","course_id") VALUES ('2025-01-16 21:09:05.61','2025-01-16 21:09:05.61',NULL,'chapter_title','chapter_description_example',37) RETURNING "id";

-- 3 create the chapter lessons:

-- 3.1 create the lesson with video content:
INSERT INTO "files" ("deleted_at","url","height","width","thumbnail_url","image_kit_id","user_id","author_id","lesson_id","video_course_id","image_course_id") VALUES (NULL,'https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/chapters/27/videos/file_yiTW-ddD4P',0,0,'',NULL,NULL,NULL,47,NULL,NULL) ON CONFLICT ("id") DO UPDATE SET "lesson_id"="excluded"."lesson_id" RETURNING "id";

INSERT INTO "lessons" ("created_at","updated_at","deleted_at","title","description","content","chapter_id") VALUES ('2025-01-16 21:12:41.066','2025-01-16 21:12:41.066',NULL,'course_title_example','course_description_example',NULL,27) RETURNING "id";

-- 3.2 create the lesson with rich text content:
INSERT INTO "lessons" ("created_at","updated_at","deleted_at","title","description","content","chapter_id") VALUES ('2025-01-16 21:16:18.687','2025-01-16 21:16:18.687',NULL,'lesson_title','lesson_description_example','{"child":{"style":{"font_size":200,"font_weight":"100"},"text":"lesson 1"},"style":{"font_size":100,"font_weight":"bold"},"text":"welcome to this lesson"}',27) RETURNING "id";

-- 3.3 create test for the chapter:
INSERT INTO "tests" ("created_at","updated_at","max_chances","chapter_id") VALUES ('2025-01-19 14:33:38.688','2025-01-19 14:33:38.688',1,27) RETURNING "id";

-- 3.3.1 create the test questions one by one:
INSERT INTO "options" ("content","is_correct","deleted_at","question_id") VALUES ('me',false,NULL,3),('him',true,NULL,3) ON CONFLICT ("id") DO UPDATE SET "question_id"="excluded"."question_id" RETURNING "id";

INSERT INTO "questions" ("created_at","updated_at","content","description","duration","deleted_at","test_id") VALUES ('2025-01-19 14:39:14.142','2025-01-19 14:39:14.142','who is your father ?','',30,NULL,3) RETURNING "id";
```

- **Get course by id:**

```sql
-- 1 ensure that the course requester is a learner or the author of this course:
SELECT * FROM "users" WHERE "users"."id" = '92d856d0-a0ee-4220-81b4-66e8adc073fc';

SELECT * FROM "authors" WHERE "authors"."id" = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "authors"."deleted_at" IS NULL;

SELECT lessons.id, lessons.title, lessons.description, lessons.chapter_id, CASE WHEN files.id IS NOT NULL THEN TRUE ELSE FALSE END AS is_video, CASE WHEN lesson_learners.lesson_id IS NOT NULL AND lesson_learners.learned = TRUE THEN TRUE ELSE FALSE END AS learned FROM "lessons" LEFT JOIN files ON lessons.id = files.lesson_id LEFT JOIN lesson_learners ON lesson_learners.lesson_id = lessons.id AND lesson_learners.learner_id = '92d856d0-a0ee-4220-81b4-66e8adc073fc' WHERE "lessons"."chapter_id" IN (24,25) AND "lessons"."deleted_at" IS NULL;

-- 2 get the questions count in the course tests:
SELECT tests.*, COUNT(questions.id) AS questions_count, CASE WHEN test_results.test_id IS NOT NULL AND test_results.has_succeed = TRUE THEN TRUE ELSE FALSE END AS has_succeed FROM "tests" LEFT JOIN test_results ON test_results.test_id = tests.id AND test_results.learner_id = '92d856d0-a0ee-4220-81b4-66e8adc073fc' LEFT JOIN questions ON questions.test_id = tests.id WHERE "tests"."chapter_id" IN (24,25) GROUP BY tests.id, test_results.test_id, test_results.has_succeed;

-- 3 get the course associations and calculate its rate:
SELECT * FROM "chapters" WHERE "chapters"."course_id" = 25 AND "chapters"."deleted_at" IS NULL;

SELECT * FROM "files" WHERE "files"."image_course_id" = 25 AND "files"."deleted_at" IS NULL;

SELECT * FROM "files" WHERE "files"."video_course_id" = 25 AND "files"."deleted_at" IS NULL;

SELECT courses.*, COALESCE(AVG(course_learners.rate), 0) AS rate FROM "courses" LEFT JOIN course_learners ON course_learners.course_id = courses.id WHERE id = '25' AND author_id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "courses"."deleted_at" IS NULL GROUP BY "courses"."id" ORDER BY "courses"."id" LIMIT 1;
```

- **Courses search:**

```sql
SELECT * FROM "files" WHERE "files"."image_course_id" = 18 AND "files"."deleted_at" IS NULL;

SELECT courses.*, COALESCE(AVG(course_learners.rate), 0) AS rate, SUM(CASE WHEN course_learners.rate IS NOT NULL THEN 1 ELSE 0 END) AS raters_count FROM "courses" JOIN course_categories ON course_categories.course_id = courses.id JOIN categories ON course_categories.category_id = categories.id LEFT JOIN course_learners ON course_learners.course_id = courses.id WHERE is_completed = true AND LOWER(title) LIKE '%go%' AND price = 0 AND language = 'en' AND level = 'advanced' AND duration >= 5 AND duration <= 150 AND categories.name IN ('dev') AND "courses"."deleted_at" IS NULL GROUP BY "courses"."id" ORDER BY rate DESC, price DESC, created_at DESC, duration DESC LIMIT 10 OFFSET 0;
```

### Subscription

#### Start A Paid Course

```sql
-- 1. Check if the user has already subscribed:
SELECT count(*) > 0 
FROM "course_learners" 
WHERE (learner_id = '443f3042-0356-450c-bc51-606cb1b0c207' AND course_id = 28) 
AND "course_learners"."deleted_at" IS NULL;

-- 2. Retrieve course details:
SELECT * 
FROM "courses" 
WHERE id = 28 
AND "courses"."deleted_at" IS NULL 
ORDER BY "courses"."id" 
LIMIT 1;

-- 3. Increment the author's balance:
UPDATE "authors" 
SET "balance" = balance + 1800 
WHERE id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' 
AND "authors"."deleted_at" IS NULL;

-- 4. Create the course session:
INSERT INTO "course_learners" ("created_at", "updated_at", "deleted_at", "course_id", "learner_id", "leaning_status", "rate", "checkout_id") 
VALUES ('2025-01-19 13:16:46.304', '2025-01-19 13:16:46.304', NULL, 28, '443f3042-0356-450c-bc51-606cb1b0c207', 'learning', NULL, '01jhzba113cbd3gd5jvgx17m75');
```