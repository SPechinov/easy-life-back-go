package users

const createTableSQL = `
	CREATE TABLE IF NOT EXISTS users (
    	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    	email VARCHAR(255) NOT NULL,
    	password VARCHAR(255) NOT NULL,
    	first_name VARCHAR(255) NOT NULL,
    	last_name VARCHAR(255),
    	created_at TIMESTAMPTZ DEFAULT NOW(),
    	updated_at TIMESTAMPTZ DEFAULT NOW()
	);`
