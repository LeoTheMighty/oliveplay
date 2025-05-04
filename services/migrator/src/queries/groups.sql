-- -------------------------------------------------------------
-- Get all groups whose home_location is within :radius_meters of
-- the supplied point (:longitude, :latitude).  Works with sqlc.
-- -------------------------------------------------------------
-- name: GetGroupsWithinRadius :many
SELECT 
    g.id,
    g.name,
    g.image,
    g.description,
    ST_X(g.location) AS longitude,
    ST_Y(g.location) AS latitude,
    g.created_at,
    g.updated_at
FROM groups AS g
WHERE g.location IS NOT NULL
  AND ST_DWithin(
        g.location::geography,
        ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
        $3  -- radius in metres
  );

-- -------------------------------------------------------------
-- Fetch all distinct activities that belong to groups found by
-- the same radius filter.  Again uses positional params for sqlc.
-- -------------------------------------------------------------
-- name: GetActivitiesWithinRadius :many
SELECT DISTINCT
    a.id,
    a.name,
    a.image,
    a.created_at,
    a.updated_at
FROM activities        AS a
JOIN group_activities  AS ga ON ga.activity_id = a.id
JOIN groups            AS g  ON g.id = ga.group_id
WHERE g.location IS NOT NULL
  AND ST_DWithin(
        g.location::geography,
        ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
        $3
  );

-- name: CreateGroup :one
INSERT INTO groups (
    name,
    image,
    description,
    location
) VALUES (
    $1, -- name
    $2, -- image (nullable)
    $3, -- description (nullable)
    -- Convert lat/lon to PostGIS Point with SRID 4326 (WGS84)
    ST_SetSRID(ST_MakePoint(sqlc.arg(longitude)::float, sqlc.arg(latitude)::float), 4326)
)
RETURNING 
    id,
    name,
    image,
    description,
    -- Extract coordinates for convenience
    ST_X(location) AS longitude,
    ST_Y(location) AS latitude,
    created_at,
    updated_at;
