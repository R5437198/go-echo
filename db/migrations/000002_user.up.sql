CREATE TABLE "user" (
    "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(10) NOT NULL,
    last_name VARCHAR(10) NOT NULL,
    birthday DATE NOT NULL,
    sex_id uuid NOT NULL,
    FOREIGN KEY (sex_id) REFERENCES sex_type(id)
);

INSERT INTO "user" ("id", first_name, last_name, birthday, sex_id) VALUES ('8cd3c5a9-a2a9-4a57-bdb5-1ee5ef3d7d3e', '坂本', '太郎', '2000-05-14', '87b42076-4fc2-49ca-9712-f2bac55d1a02');
INSERT INTO "user" ("id", first_name, last_name, birthday, sex_id) VALUES ('cde97584-a8f3-4ba5-86e7-9d4f67ed86d2', '坂本', '雪', '2000-05-13', 'b19d942e-06d1-4bc3-993a-50da505ff942');