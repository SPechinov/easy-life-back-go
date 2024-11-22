CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- groups_users table
CREATE TABLE IF NOT EXISTS public.notes_info (
    note_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    description TEXT,

    CONSTRAINT fk_note_id FOREIGN KEY (note_id) REFERENCES public.notes(id)
);

--  indexes
CREATE INDEX uid_note_id ON public.notes_info(note_id);
