BEGIN;

-- Insert Roles
INSERT INTO roles (title) VALUES
  ('customer'),
  ('admin');

-- Insert Users
INSERT INTO users (id, username, email, password, role_id, first_name, last_name, is_verified, created_at, update_at)
VALUES
  (1, 'user', 'user@example.com', '$2a$10$8KzaNdKIMyOkASCH4QvSKuEMIY7Jc3vcHDuSJvXLii1rvBNgz60a6', 1, 'USER FIRST', 'USER LAST', true, now(), now()),
  (2, 'admin', 'admin1@example.com', '$2a$10$3qqNPE.TJpNGYCohjTgw9.v1z0ckovx95AmiEtUXcixGAgfW7.wCi', 2, 'ADMIN FIRST', 'ADMIN LAST', true, now(), now());

-- Insert Categories
INSERT INTO categories (name, created_at, update_at)
VALUES
  ('Digital Art', now(), now()),
  ('Photography', now(), now()),
  ('Music', now(), now()),
  ('Collectibles', now(), now()),
  ('Sports', now(), now()),
  ('Virtual Real Estate', now(), now());

-- Insert NFTs
INSERT INTO nfts (title, description, price, image_url, author_id, owner_id, category, listing_type, floor_bid, end_time, status, created_at, update_at)
VALUES
  ('NFT1', 'Description for NFT1', 100.00, 'nft1_image_url.jpg', 1, 1, 1, 'fixed', NULL, NULL, 'available', now(), now()),
  ('NFT2', 'Description for NFT2', 150.00, 'nft2_image_url.jpg', 1, 1, 2, 'fixed', NULL, NULL, 'available', now(), now()),
  ('NFT3', 'Description for NFT3', 200.00, 'nft3_image_url.jpg', 1, 1, 3, 'fixed', NULL, NULL, 'available', now(), now()),
  ('NFT4', 'Description for NFT4', 120.00, 'nft4_image_url.jpg', 1, 1, 4, 'auction', NULL, '2024-03-01 12:00:00', 'available', now(), now()),
  ('NFT5', 'Description for NFT5', 180.00, 'nft5_image_url.jpg', 1, 1, 5, 'fixed', NULL, NULL, 'available', now(), now()),
  ('NFT6', 'Description for NFT6', 250.00, 'nft6_image_url.jpg', 1, 1, 6, 'fixed', NULL, NULL, 'available', now(), now());

-- Insert Products Categories
INSERT INTO products_categories (nft_id, category_id)
VALUES
  ('N000001', 1),
  ('N000002', 2),
  ('N000003', 3),
  ('N000004', 4),
  ('N000005', 5),
  ('N000006', 6);

-- Insert Images
INSERT INTO images (id, filename, url, nft_id, created_at, update_at)
VALUES
  ('1a11ae5b-e905-49ed-9b3b-e3977a59afd8', 'nft1_image.jpg', 'https://img.freepik.com/free-psd/3d-nft-icon-nft-memes_629802-24.jpg', 'N000001', now(), now()),
  ('41646f8a-64ae-4ccf-a9c8-a8debe4cd16c', 'nft2_image.jpg', 'https://img.freepik.com/free-psd/nft-cryptocurrency-3d-illustration_1419-2742.jpg', 'N000002', now(), now()),
  ('e70ada61-e118-40bc-972d-0980d8302e77', 'nft3_image.jpg', 'https://img.freepik.com/free-vector/hand-drawn-nft-style-ape-illustration_23-2149622021.jpg', 'N000003', now(), now()),
  ('f0e0b6a9-9b9a-4b9e-9e9a-9b9a9b9a9b9a', 'nft4_image.jpg', 'https://img.freepik.com/premium-photo/item-3d-render-icon-illustration_726846-3292.jpg', 'N000004', now(), now()),
  ('c37cabeb-3f8e-4f75-98d0-fd08c0aaf70f', 'nft5_image.jpg', 'https://img.freepik.com/premium-photo/artist-3d-render-icon-illustration_726846-3249.jpg', 'N000005', now(), now()),
  ('94f24a64-496e-4700-a946-de45f76b0df8', 'nft6_image.jpg', 'https://img.freepik.com/free-vector/isometric-nft-illustration_23-2148950467.jpg', 'N000006', now(), now());

COMMIT;