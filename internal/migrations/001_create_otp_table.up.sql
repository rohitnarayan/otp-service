CREATE TABLE otp (
       id SERIAL PRIMARY KEY,
       otp INT NOT NULL,
       user_id VARCHAR(255) NOT NULL,
       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP NOT NULL,
       status VARCHAR(255) NOT NULL
);

-- Create an index on the otp column
CREATE INDEX idx_otp ON otp (otp);

-- Create an index on the user_id column
CREATE INDEX idx_user_id ON otp (user_id);
