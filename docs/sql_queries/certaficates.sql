-- Start A paid course:

-- 1 check if the user has already :
SELECT count(*) > 0 FROM "course_learners" WHERE (learner_id = '443f3042-0356-450c-bc51-606cb1b0c207' AND course_id = 28) AND "course_learners"."deleted_at" IS NULL

SELECT * FROM "courses" WHERE id = 28 AND "courses"."deleted_at" IS NULL ORDER BY "courses"."id" LIMIT 1

-- 2 increment the author balance:
UPDATE "authors" SET "balance"=balance + 1800 WHERE id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "authors"."deleted_at" IS NULL

-- 3 create the course session:
INSERT INTO "course_learners" ("created_at","updated_at","deleted_at","course_id","learner_id","leaning_status","rate","checkout_id") VALUES ('2025-01-19 13:16:46.304','2025-01-19 13:16:46.304',NULL,28,'443f3042-0356-450c-bc51-606cb1b0c207','learning',NULL,'01jhzba113cbd3gd5jvgx17m75')