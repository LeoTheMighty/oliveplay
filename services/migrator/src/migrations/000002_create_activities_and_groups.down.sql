-- Drop join table first because of FKs
DROP TABLE IF EXISTS group_activities;

-- Drop the main tables
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS activities;

-- Spatial index is dropped automatically with `groups` but keep this
DROP INDEX IF EXISTS idx_groups_home_location_gist;

-- Triggers are removed automatically when their table is dropped.
-- We purposely leave the `trigger_set_timestamp` function and the
-- PostGIS extension in place because other migrations may rely on them.