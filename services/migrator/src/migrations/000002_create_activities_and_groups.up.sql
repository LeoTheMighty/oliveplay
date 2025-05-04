-- Enable PostGIS (needed for radius queries on lat/lon points)
CREATE EXTENSION IF NOT EXISTS postgis;

-- ------------------------------------------------------------------
-- Helper function – updates the `updated_at` column automatically.
-- We define it once with OR REPLACE so other migrations can reuse it.
-- ------------------------------------------------------------------
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ------------------------------------------------------------------
-- Activities
-- ------------------------------------------------------------------
CREATE TABLE activities (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    image       TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_activities
BEFORE UPDATE ON activities
FOR EACH ROW EXECUTE FUNCTION trigger_set_timestamp();

-- ------------------------------------------------------------------
-- Groups
-- home_location stored as geometry(Point,4326).  We'll cast to
-- geography on-the-fly in distance queries so callers can still
-- work in metres.
-- ------------------------------------------------------------------
CREATE TABLE groups (
    id            SERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    image         TEXT,
    description   TEXT,
    location geometry(Point,4326),    -- nullable
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Spatial index greatly speeds up the ST_DWithin radius query
CREATE INDEX idx_groups_location_gist ON groups USING GIST (location);

CREATE TRIGGER set_timestamp_groups
BEFORE UPDATE ON groups
FOR EACH ROW EXECUTE FUNCTION trigger_set_timestamp();

-- ------------------------------------------------------------------
-- Join table – Many-to-Many (Group × Activity)
-- ------------------------------------------------------------------
CREATE TABLE group_activities (
    group_id     INTEGER NOT NULL REFERENCES groups(id)     ON DELETE CASCADE,
    activity_id  INTEGER NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, activity_id)
);