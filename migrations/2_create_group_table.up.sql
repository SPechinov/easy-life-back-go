CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- groups table
CREATE TABLE IF NOT EXISTS public.groups
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    is_payed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

--  indexes
CREATE UNIQUE INDEX IF NOT EXISTS unique_id_index_group_id ON public.groups (id);

--  triggers
DROP TRIGGER IF EXISTS update_groups_updated_at ON public.groups;
CREATE OR REPLACE FUNCTION update_updated_at_column()
	RETURNS TRIGGER AS $$
BEGIN
		NEW.updated_at = NOW();
RETURN NEW;
END;
	$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_groups_updated_at
    BEFORE UPDATE ON public.groups
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
