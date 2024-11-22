CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--  table
CREATE TABLE IF NOT EXISTS public.users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(60) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    password BYTEA NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

--  indexes
CREATE UNIQUE INDEX IF NOT EXISTS unique_id_index_users_id ON public.users(id);
CREATE UNIQUE INDEX IF NOT EXISTS unique_id_index_users_email ON public.users(email);
CREATE UNIQUE INDEX IF NOT EXISTS unique_id_index_users_phone ON public.users(phone);

--  triggers
DROP TRIGGER IF EXISTS update_users_updated_at ON public.users;
CREATE OR REPLACE FUNCTION update_updated_at_column()
	RETURNS TRIGGER AS $$
BEGIN
		NEW.updated_at = NOW();
RETURN NEW;
END;
	$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON public.users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();