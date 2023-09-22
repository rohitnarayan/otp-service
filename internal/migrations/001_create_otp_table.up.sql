CREATE TABLE otp (
       id SERIAL PRIMARY KEY,
       otp VARCHAR(10) NOT NULL,
       user_id VARCHAR(255) NOT NULL,
       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP NOT NULL,
       status VARCHAR(10) NOT NULL
);

-- Create an index on the otp column
-- CREATE INDEX idx_otp ON otp (otp);

-- Create an index on the user_id column
-- CREATE INDEX idx_user_id ON otp (user_id);

-- Create an index on the status column
CREATE INDEX idx_user_id_status ON otp (user_id, status);

CREATE INDEX idx_user_id_otp_status ON otp (user_id, otp, status);
