INSERT INTO users (id, email, password, created_at, updated_at, deleted_at) VALUES
('48f43073-0016-44ee-acd4-549ec499fc1b', 'rayendratimotius@gmail.com', '$2a$10$G/2D8xxEQMis8ZfqTOGd3uCVV2lZVJ4HOF8CkHtA4kSb.ChBSJuvq', NOW(), NOW(), NULL)
ON CONFLICT (id) 
DO UPDATE SET 
    email = EXCLUDED.email,
    password = EXCLUDED.password,
    updated_at = NOW();