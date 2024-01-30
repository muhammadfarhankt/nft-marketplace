BEGIN;

TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE nfts CASCADE;
TRUNCATE TABLE categories CASCADE;
TRUNCATE TABLE products_categories CASCADE;
TRUNCATE TABLE images CASCADE;

SELECT SETVAL ((SELECT pg_get_serial_sequence('"roles"', 'id')), 1, false);
SELECT SETVAL ((SELECT pg_get_serial_sequence('"categories"', 'id')), 1, false);

COMMIT;