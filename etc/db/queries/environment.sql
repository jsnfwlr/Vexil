-- name: ListEnvironments :many
SELECT name FROM environment WHERE deleted_at IS NULL;