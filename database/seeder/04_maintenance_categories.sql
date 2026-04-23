INSERT INTO maintenance_categories (id, name, description, created_at, updated_at, deleted_at) VALUES
('96a5d0cc-56ab-4116-bccb-ec49c486c9f3', 'Powertrain', 'Focuses on the engine and transmission—the most critical and expensive components.', NOW(), NOW(), NULL),
('40995a28-781f-4705-ae90-f5df0bd2244c', 'Chassis & Handling', 'Includes everything under the car that manages movement, stopping, and safety.', NOW(), NOW(), NULL),
('adb4cc0b-5f22-4260-8686-a4dbc9f479cb', 'Support & Fluids', 'Consumable liquids and cooling components that prevent mechanical failure', NOW(), NOW(), NULL),
('8568897d-d3f2-465f-abee-09aa9656de49', 'Electrical & Tech', 'Manages the power, lighting, and modern conveniences of the vehicle.', NOW(), NOW(), NULL),
('4ce1ef90-3549-4c66-a927-0e29f695510e', 'Body & Exterior', 'Focuses on environmental protection and the physical look of the car.', NOW(), NOW(), NULL),
('964b5ddb-be72-4792-95ad-93403cc32565', 'Interior & Comfort', 'Everything inside the car that affects your driving experience.', NOW(), NOW(), NULL),
('cffa5633-fc81-4cb3-b1ae-9719f01024d4', 'Service', 'The place you went to to get your vehicle serviced.', NOW(), NOW(), NULL)
ON CONFLICT (id)
DO UPDATE SET 
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    updated_at = NOW();