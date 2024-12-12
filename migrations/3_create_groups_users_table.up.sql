CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- groups_users table
CREATE TABLE IF NOT EXISTS groups_users (
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    permission INT NOT NULL,
    invited_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_unique_group_user UNIQUE (group_id, user_id)
);

--  indexes
CREATE INDEX uid_user_group ON groups_users (group_id, user_id);
