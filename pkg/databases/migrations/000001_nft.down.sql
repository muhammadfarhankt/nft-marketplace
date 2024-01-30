BEGIN;

DROP TRIGGER IF EXISTS nfts_update_at ON nfts;
DROP TRIGGER IF EXISTS users_update_at ON users;
DROP TRIGGER IF EXISTS categories_update_at ON categories;
DROP TRIGGER IF EXISTS products_categories_update_at ON products_categories;
DROP TRIGGER IF EXISTS images_update_at ON images;

DROP FUNCTION IF EXISTS update_updated_at_column;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS nfts CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS products_categories CASCADE;
DROP TABLE IF EXISTS images CASCADE;

DROP SEQUENCE IF EXISTS nfts_id_seq;
DROP SEQUENCE IF EXISTS users_id_seq;
DROP SEQUENCE IF EXISTS categories_id_seq;
DROP SEQUENCE IF EXISTS products_categories_id_seq;
DROP SEQUENCE IF EXISTS images_id_seq;

COMMIT;