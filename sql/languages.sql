BEGIN;
CREATE TABLE languages (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    
    -- format:  <country>[-<language>]
    id VARCHAR
    
    -- ISO 3166 Country Code
    country VARCHAR NOT NULL

    -- ISO 639 Language Code
    language VARCHAR 
);
CREATE UNIQUE INDEX language_pk_idx ON languages (id);
CREATE INDEX language_language_idx ON languages (country);
COMMIT;
