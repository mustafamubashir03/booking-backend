-- +goose Up
CREATE TABLE IF NOT EXISTS permissions(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO permissions(name,description,resource,action) VALUES
('user:read','Read user data','user','read'),
('user:create','Create user data','user','create'),
('user:delete','Delete user data','user','delete'),
('user:update','Update user data','user','update'),
('role:read','Read role data','role','read'),
('role:create','Create role data','role','create'),
('role:delete','Delete role data','role','delete'),
('role:update','Update role data','role','update'),
('permission:create','Create permission data','permission','create'),
('permission:delete','Delete permission data','permission','delete'),
('permission:update','Update permission data','permission','update');

-- +goose Down
DROP TABLE IF EXISTS permissions;
