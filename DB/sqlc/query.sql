-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY name;

-- name: CreateEmployee :one
INSERT INTO employees (
    name
) VALUES (
             $1
         )
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = $1;

-- name: ListSalary :many
SELECT * FROM salaries
ORDER BY id;

-- name: CreateSalary :one
INSERT INTO salaries (
    employee_id, name
) VALUES (
             $1, $2
         )
RETURNING *;

-- name: DeleteSalary :exec
DELETE FROM salaries
WHERE id = $1;