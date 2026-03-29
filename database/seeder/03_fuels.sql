INSERT INTO fuels (id, brand, "type", "name", price, created_at, updated_at, deleted_at) VALUES
('503166d1-5b84-40a8-9460-ba3283259694', 'pertamina', 'petrol', 'pertalite', 10000, NOW(), NOW(), NULL),
('bc1d7d87-c426-4174-8336-1c05e80f44c7', 'pertamina', 'petrol', 'pertamax', 0, NOW(), NOW(), NULL),
('614ed908-3b4e-4c76-8c3f-c66e92002e7a', 'pertamina', 'petrol', 'pertamax turbo', 0, NOW(), NOW(), NULL),
('eeaccb27-fc62-4fa0-be21-ff9bef2ee5de', 'pertamina', 'diesel', 'biosolar', 6800, NOW(), NOW(), NULL),
('8dd843d8-3057-447a-9bd5-71292b3865d0', 'pertamina', 'diesel', 'dexlite', 0, NOW(), NOW(), NULL),
('906a7174-2207-493c-9707-9c5e47d3d0ad', 'pertamina', 'diesel', 'dex', 0, NOW(), NOW(), NULL)
ON CONFLICT (id)
DO UPDATE SET 
    brand = EXCLUDED.brand,
    "type" = EXCLUDED."type",
    "name" = EXCLUDED."name",
    price = EXCLUDED.price,
    updated_at = NOW();
