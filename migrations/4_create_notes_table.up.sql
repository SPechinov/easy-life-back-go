CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- groups_users table
CREATE TABLE IF NOT EXISTS notes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(50),
    group_id UUID NOT NULL,
    user_creator_id UUID NOT NULL,
    user_updater_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id),
    CONSTRAINT fk_user_creator FOREIGN KEY (user_creator_id) REFERENCES users(id),
    CONSTRAINT fk_user_updater FOREIGN KEY (user_updater_id) REFERENCES users(id)
);

--  indexes
CREATE INDEX uid_user_id ON notes (id);
