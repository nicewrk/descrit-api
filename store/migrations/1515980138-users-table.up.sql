CREATE EXTENSION "uuid-ossp";

CREATE TABLE users (
	uid UUID PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	username TEXT UNIQUE NOT NULL,
	is_verified BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE FUNCTION func_generate_uid()
RETURNS TRIGGER
LANGUAGE plpgsql
AS
$$
	BEGIN
		NEW.uid := uuid_generate_v4();
		RETURN NEW;
	END;
$$;

CREATE TRIGGER trigger_generate_uid
	BEFORE INSERT ON users
	FOR EACH ROW
	EXECUTE PROCEDURE func_generate_uid();

CREATE FUNCTION func_updated_at()
RETURNS TRIGGER
LANGUAGE plpgsql
AS
$$
	BEGIN
		NEW.updated_at := NOW();
		RETURN NEW;
	END;
$$;

CREATE TRIGGER trigger_updated_at
	BEFORE UPDATE ON users
	FOR EACH ROW
	EXECUTE PROCEDURE func_updated_at();

