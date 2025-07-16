-- name: GetFlagsByEnvironmentName :many
SELECT sqlc.embed(feature_flag), 
       sqlc.embed(environment),
       sqlc.embed(feature_flag_value)
FROM feature_flag
JOIN feature_flag_value ON feature_flag.id = feature_flag_value.feature_flag_id
JOIN environment ON environment.id = feature_flag_value.environment_id
WHERE feature_flag.deleted_at IS NULL
  AND environment.deleted_at IS NULL
  AND feature_flag_value.deleted_at IS NULL
  AND environment.name = $1;

-- name: GetFlags :many
SELECT sqlc.embed(feature_flag), 
       sqlc.embed(environment),
       sqlc.embed(feature_flag_value)
FROM feature_flag
JOIN feature_flag_value ON feature_flag.id = feature_flag_value.feature_flag_id
JOIN environment ON environment.id = feature_flag_value.environment_id
WHERE feature_flag.deleted_at IS NULL
  AND environment.deleted_at IS NULL
  AND feature_flag_value.deleted_at IS NULL;