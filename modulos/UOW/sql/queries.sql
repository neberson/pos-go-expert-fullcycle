-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES (?, ?);

-- name: CreateCourse :exec
INSERT INTO courses (id, category_id, name) VALUES (?, ?, ?);