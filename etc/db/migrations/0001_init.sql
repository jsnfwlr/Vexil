CREATE TYPE FLAG_TYPE AS ENUM (
    'boolean',
    'string',
    'integer',
    'json',
    'string_array',
    'integer_array'
);

CREATE TABLE IF NOT EXISTS feature_flag (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    default_value TEXT NOT NULL DEFAULT 'false', -- when a new environment is created, this value will be used for the flag in that environment
    value_type FLAG_TYPE NOT NULL default 'boolean',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS environment (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS feature_flag_value (
    feature_flag_id INTEGER NOT NULL REFERENCES feature_flag(id),
    environment_id INTEGER NOT NULL REFERENCES environment(id),
    value TEXT NOT NULL DEFAULT 'false',
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT feature_flag_environment_unique
    UNIQUE (feature_flag_id, environment_id)
);