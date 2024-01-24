-- Create the posts table
CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       user_id UUID REFERENCES "user"(id),
                       content TEXT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- Add other post-related columns as needed
);

-- Create the images table
CREATE TABLE images (
                        id SERIAL PRIMARY KEY,
                        post_id INT REFERENCES "posts"(id),
                        image_url VARCHAR(255) NOT NULL,
                        filename VARCHAR(255) NOT NULL
    -- Add other image-related columns as needed
);