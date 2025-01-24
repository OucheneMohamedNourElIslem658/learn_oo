-- Create course:

-- 1 create the course descriptive componetents:
INSERT INTO "files" ("deleted_at","url","height","width","thumbnail_url","image_kit_id","user_id","author_id","lesson_id","video_course_id","image_course_id") VALUES (NULL,'https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/videos/file_xlkchxPX1',360,640,'',NULL,NULL,NULL,NULL,37,NULL) ON CONFLICT ("id") DO UPDATE SET "video_course_id"="excluded"."video_course_id" RETURNING "id"

INSERT INTO "files" ("deleted_at","url","height","width","thumbnail_url","image_kit_id","user_id","author_id","lesson_id","video_course_id","image_course_id") VALUES (NULL,'https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/images/file_4iESjwQwJ',980,1848,'https://ik.imagekit.io/cdejmhtxd/tr:n-ik_ml_thumbnail/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/images/file_4iESjwQwJ',NULL,NULL,NULL,NULL,NULL,37) ON CONFLICT ("id") DO UPDATE SET "image_course_id"="excluded"."image_course_id" RETURNING "id"

INSERT INTO "courses" ("created_at","updated_at","deleted_at","title","description","price","payment_price_id","payment_product_id","language","level","duration","is_completed","author_id") VALUES ('2025-01-16 20:59:47.027','2025-01-16 20:59:47.027',NULL,'course_title_example','course_description_example',2000,'01jhrb6d9902zr4dad4a1anfhw','01jhrb6bve7jktcrck6d064msh','en','advanced',5,false,'72125ca0-94f3-42e5-ac3a-0d797ec9078b') RETURNING "id"

-- 2 create the course chapters:

SELECT count(*) FROM "courses" WHERE (id = '37' and author_id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b') AND "courses"."deleted_at" IS NULL

INSERT INTO "chapters" ("created_at","updated_at","deleted_at","title","description","course_id") VALUES ('2025-01-16 21:09:05.61','2025-01-16 21:09:05.61',NULL,'chapter_title','chapter_description_example',37) RETURNING "id"

-- 3 create the chapter lessons:

-- 3.1 create the lesson with video content:

INSERT INTO "files" ("deleted_at","url","height","width","thumbnail_url","image_kit_id","user_id","author_id","lesson_id","video_course_id","image_course_id") VALUES (NULL,'https://ik.imagekit.io/cdejmhtxd/learn_oo/authors/72125ca0-94f3-42e5-ac3a-0d797ec9078b/courses/chapters/27/videos/file_yiTW-ddD4P',0,0,'',NULL,NULL,NULL,47,NULL,NULL) ON CONFLICT ("id") DO UPDATE SET "lesson_id"="excluded"."lesson_id" RETURNING "id"

INSERT INTO "lessons" ("created_at","updated_at","deleted_at","title","description","content","chapter_id") VALUES ('2025-01-16 21:12:41.066','2025-01-16 21:12:41.066',NULL,'course_title_example','course_description_example',NULL,27) RETURNING "id"

-- 3.2 create the lesson with rich text content:

INSERT INTO "lessons" ("created_at","updated_at","deleted_at","title","description","content","chapter_id") VALUES ('2025-01-16 21:16:18.687','2025-01-16 21:16:18.687',NULL,'lesson_title','lesson_description_example','{"child":{"style":{"font_size":200,"font_weight":"100"},"text":"lesson 1"},"style":{"font_size":100,"font_weight":"bold"},"text":"welcome to this lesson"}',27) RETURNING "id"

-- 3.3 create test for the chapter:

INSERT INTO "tests" ("created_at","updated_at","max_chances","chapter_id") VALUES ('2025-01-19 14:33:38.688','2025-01-19 14:33:38.688',1,27) RETURNING "id" 

-- 3.3.1 create the test questions one by one:

INSERT INTO "options" ("content","is_correct","deleted_at","question_id") VALUES ('me',false,NULL,3),('him',true,NULL,3) ON CONFLICT ("id") DO UPDATE SET "question_id"="excluded"."question_id" RETURNING "id"

INSERT INTO "questions" ("created_at","updated_at","content","description","duration","deleted_at","test_id") VALUES ('2025-01-19 14:39:14.142','2025-01-19 14:39:14.142','who is your father ?','',30,NULL,3) RETURNING "id"


-- --------------------------------------------------------------------------------------------------------------------------------


-- Get course by id:

-- 1 ensure that the course requrester is a learner or the author of this course :

SELECT * FROM "users" WHERE "users"."id" = '92d856d0-a0ee-4220-81b4-66e8adc073fc'

SELECT * FROM "authors" WHERE "authors"."id" = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "authors"."deleted_at" IS NULL

SELECT lessons.id, lessons.title, lessons.description, lessons.chapter_id, CASE WHEN files.id IS NOT NULL THEN TRUE ELSE FALSE END AS is_video, CASE WHEN lesson_learners.lesson_id IS NOT NULL AND lesson_learners.learned = TRUE THEN TRUE ELSE FALSE END AS learned FROM "lessons" LEFT JOIN files ON lessons.id = files.lesson_id LEFT JOIN lesson_learners ON lesson_learners.lesson_id = lessons.id AND lesson_learners.learner_id = '92d856d0-a0ee-4220-81b4-66e8adc073fc' WHERE "lessons"."chapter_id" IN (24,25) AND "lessons"."deleted_at" IS NULL

-- 2 get the questons count is the course tests:

SELECT tests.*, COUNT(questions.id) AS questions_count, CASE WHEN test_results.test_id IS NOT NULL AND test_results.has_succeed = TRUE THEN TRUE ELSE FALSE END AS has_succeed FROM "tests" LEFT JOIN test_results ON test_results.test_id = tests.id AND test_results.learner_id = '92d856d0-a0ee-4220-81b4-66e8adc073fc' LEFT JOIN questions ON questions.test_id = tests.id WHERE "tests"."chapter_id" IN (24,25) GROUP BY tests.id, test_results.test_id, test_results.has_succeed

-- 3 get the course associations and calculate its rate:

SELECT * FROM "chapters" WHERE "chapters"."course_id" = 25 AND "chapters"."deleted_at" IS NULL

SELECT * FROM "files" WHERE "files"."image_course_id" = 25 AND "files"."deleted_at" IS NULL

SELECT * FROM "files" WHERE "files"."video_course_id" = 25 AND "files"."deleted_at" IS NULL

SELECT courses.*, COALESCE(AVG(course_learners.rate), 0) AS rate FROM "courses" LEFT JOIN course_learners ON course_learners.course_id = courses.id WHERE id = '25' AND author_id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "courses"."deleted_at" IS NULL GROUP BY "courses"."id" ORDER BY "courses"."id" LIMIT 1


-- --------------------------------------------------------------------------------------------------------------------------------

-- courses search: smart search with pagination and muli filters with sorting by rate and ... to get top courses:

SELECT * FROM "files" WHERE "files"."image_course_id" = 18 AND "files"."deleted_at" IS NULL

SELECT courses.*,
    COALESCE(AVG(course_learners.rate), 0) AS rate,
    SUM(CASE WHEN course_learners.rate IS NOT NULL THEN 1 ELSE 0 END) AS raters_count FROM "courses" 
    JOIN course_categories ON course_categories.course_id = courses.id 
    JOIN categories ON course_categories.category_id = categories.id LEFT 
    JOIN course_learners ON course_learners.course_id = courses.id 
    WHERE is_completed = true AND LOWER(title) LIKE '%go%' AND price = 0 AND language = 'en' AND level = 'advanced' AND duration >= 5 AND duration <= 150 AND categories.name IN ('dev') AND "courses"."deleted_at" IS NULL 
    GROUP BY "courses"."id" 
    ORDER BY rate DESC, price DESC, created_at DESC, duration DESC 
    LIMIT 10
    OFFSET 0