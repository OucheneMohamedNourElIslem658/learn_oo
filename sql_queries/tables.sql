-- TABLE CREATION
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
    "level" text DEFAULT 'bigener',
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


-- INDEX CREATION FOR deleted_at COLUMN

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