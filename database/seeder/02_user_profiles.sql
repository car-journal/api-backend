INSERT INTO user_profiles (user_id, gender, first_name, last_name, date_of_birth, picture_url, created_at, updated_at, deleted_at) VALUES
('48f43073-0016-44ee-acd4-549ec499fc1b', true, 'Rayendra', 'Sabandar', '1997-09-20 12:51:30.448 +0700', NULL, NOW(), NOW(), NULL)
ON CONFLICT (id) 
DO UPDATE SET 
    gender = EXCLUDED.gender,
    first_name = EXCLUDED.first_name,
    last_name = EXCLUDED.last_name,
    date_of_birth = EXCLUDED.date_of_birth,
    picture_url = EXCLUDED.picture_url,
    updated_at = NOW();
