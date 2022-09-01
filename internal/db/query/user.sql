-- name: CreateUser :one
INSERT INTO users (
  id,
  first_name,
  last_name,
  nickname,
  hashed_password,
  email,
  country,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users
WHERE country LIKE $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateUser :exec
UPDATE users
SET first_name = $1,
    last_name = $2,
    nickname = $3,
    email = $4,
    country = $5
WHERE id = $6;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
