CREATE TABLE sex_type (
    "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    value VARCHAR(10) NOT NULL,
    "order" VARCHAR(2) NOT NULL
);

INSERT INTO sex_type ("id", value, "order") VALUES('87b42076-4fc2-49ca-9712-f2bac55d1a02', '男', '1');
INSERT INTO sex_type ("id", value, "order") VALUES('b19d942e-06d1-4bc3-993a-50da505ff942', '女', '2');