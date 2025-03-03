BEGIN;
ALTER TABLE links
ADD COLUMN title varchar,
ADD COLUMN og_title varchar,
ADD COLUMN og_description varchar,
ADD COLUMN og_site_name varchar,
ADD COLUMN og_image varchar,
ADD COLUMN og_image_options json,
ADD COLUMN og_optional json;
COMMIT;