CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE access_roles AS ENUM ('admin', 'user');

CREATE TABLE users (
	id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL,
  email TEXT NOT NULL,
	role access_roles NOT NULL DEFAULT 'user',
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMP DEFAULT NULL,
	PRIMARY KEY (id)
);