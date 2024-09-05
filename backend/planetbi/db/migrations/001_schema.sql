-- Write your migrate up statements here
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS
	secrets (
		id TEXT PRIMARY KEY NOT NULL,
		VALUE TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);

CREATE TABLE IF NOT EXISTS
	users (
		id TEXT PRIMARY KEY NOT NULL,
		NAME TEXT NOT NULL,
		email CITEXT NOT NULL,
		email_verified BOOLEAN NOT NULL DEFAULT FALSE,
		email_verification_token TEXT NOT NULL DEFAULT ''::TEXT,
		email_verification_token_expiry TIMESTAMPTZ,
		PASSWORD TEXT NOT NULL,
		password_reset_token TEXT NOT NULL DEFAULT ''::TEXT,
		password_reset_correlation_token TEXT NOT NULL DEFAULT ''::TEXT,
		password_reset_token_expiry TIMESTAMPTZ,
		bio TEXT NOT NULL DEFAULT ''::TEXT,
		ROLE TEXT NOT NULL DEFAULT 'user'::TEXT,
		recent_login_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		CONSTRAINT users_unique_email UNIQUE (email)
	);

---- create above / drop below ----
DROP TABLE IF EXISTS secrets;

DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS citext;