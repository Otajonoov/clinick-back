CREATE TABLE IF NOT EXISTS "doctors" (
    "id" SERIAL PRIMARY KEY,
    "fullname" VARCHAR NOT NULL,
    "type" VARCHAR NOT NULL,
    "about" TEXT NOT NULL,
    "img_url" TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "services" (
    "id" SERIAL PRIMARY KEY,
    "servicename" VARCHAR NOT NULL,
    "about" TEXT NOT NULL,
    "img_url" TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "customers" (
    "id" SERIAL PRIMARY KEY,
    "fullname" VARCHAR NOT NULL,
    "stars" VARCHAR NOT NULL,
    "about" TEXT NOT NULL,
    "img_url" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR NOT NULL,
    "password" TEXT NOT NULL
);