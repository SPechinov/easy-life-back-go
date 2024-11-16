CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- users_groups table
CREATE TABLE IF NOT EXISTS public.users_groups (
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    permission INT NOT NULL,
    invited_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES public.groups(id),
    CONSTRAINT fk_admin FOREIGN KEY (user_id) REFERENCES public.users(id),
    CONSTRAINT unique_user_group UNIQUE (group_id, user_id)
);

--  indexes
CREATE INDEX unique_id_index_user_group ON users_groups (group_id, user_id);
