-- Drop the index on the "otp" column
DROP INDEX IF EXISTS idx_otp;

-- Drop the index on the "user_id" column
DROP INDEX IF EXISTS idx_user_id;

-- Drop the "otp" table if it exists
DROP TABLE IF EXISTS otp;
