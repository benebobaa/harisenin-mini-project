CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
                        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                        username VARCHAR(255) UNIQUE NOT NULL,
                        password VARCHAR(255) NOT NULL
);
