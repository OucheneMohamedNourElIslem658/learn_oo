-- Create user with email and password (password is hashed with bcrypt):

INSERT INTO "users" ("id","email","password","full_name","email_verified") VALUES ('7308f1cf-2ede-476d-869a-7be31d03afbf','ouchene@estin.dz','$2a$10$FX1XpMqd/GgJ6RWbugxINu.RR/Zfk2b7sZG1UIwnHM7MOGdsoCdla','ouchene mohamed',false)

-- upgrade the user to an author:

INSERT INTO "authors" ("id","deleted_at","bio","balance","user_id") VALUES ('325fba00-f411-43f8-b3c8-0ed2b57f040f',NULL,NULL,0,'9f69b120-4451-4f14-9705-dd438d882143')

-- downgrade from author to normal user (logicaly delete the author profile):

UPDATE "authors" SET "deleted_at"='2025-01-19 15:35:46.663' WHERE id = '72125ca0-94f3-42e5-ac3a-0d797ec9078b' AND "authors"."deleted_at" IS NULL