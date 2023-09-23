CREATE TABLE otp (
       id SERIAL PRIMARY KEY,
       otp VARCHAR(10) NOT NULL,
       user_id VARCHAR(255) NOT NULL,
       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP NOT NULL,
       status VARCHAR(10) NOT NULL
);

CREATE INDEX idx_user_id_status ON otp (user_id, status);

CREATE INDEX idx_user_id_otp_status ON otp (user_id, otp, status);
