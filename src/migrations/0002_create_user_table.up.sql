CREATE TYPE auth_providers AS ENUM ('google', 'apple');

CREATE TABLE IF NOT EXISTS public.user (
	pk: SERIAL PRIMARY KEY,
	user_id: INT4
	provider: AUTH_PROVIDERS,
	UNIQUE (user_id, provider)
);


CREATE TABLE IF NOT EXISTS public.image (
	user: REFERENCES user.pk
	url: TEXT NOT NULL;
)
-- https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g