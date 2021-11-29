CREATE TABLE employees
(
    id BIGSERIAL PRIMARY KEY,
    name varchar(100)
);

CREATE TABLE salaries
(
    id BIGSERIAL PRIMARY KEY,
    employee_id int,
    name varchar(100),
    CONSTRAINT salary_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employees (id)
);