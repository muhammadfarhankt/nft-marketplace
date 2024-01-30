BEGIN;

-- SET TIMEZONE
SET TIME ZONE 'Asia/Kolkata';

-- CREATE EXTENSIONS
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--users_id eg START WITH -> U000001
--nfts_id -> N000001
--CREATE SEQUENCES
CREATE SEQUENCE users_id_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE nfts_id_seq START WITH 1 INCREMENT BY 1;

--AUTO UPDATE
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.update_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';

-- CREATE TABLES
CREATE TABLE "users" (
  "id" varchar(7) PRIMARY KEY DEFAULT CONCAT('U', LPAD(nextval('users_id_seq')::text, 6, '0')),
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role_id" int NOT NULL,
  "first_name" varchar(100),
  "last_name" varchar(100),
  "is_verified" boolean DEFAULT false,
  "verification_token" varchar(255),
  "image_url" varchar(255),
  "bio" text,
  "banner" varchar(255),
  "instagram" varchar(255),
  "twitter" varchar(255),
  "site" varchar(255),
  "wallet_amount" numeric(10, 2) DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "roles" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL UNIQUE
);

CREATE TABLE "oauth" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" varchar NOT NULL,
  "access_token" varchar(12) NOT NULL,
  "refresh_token" varchar(12) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "nfts" (
  "id" varchar(7) PRIMARY KEY DEFAULT CONCAT('N', LPAD(nextval('nfts_id_seq')::text, 6, '0')),
  "title" varchar(255) NOT NULL,
  "description" varchar NOT NULL,
  "price" FLOAT NOT NULL,
  "image_url" varchar(255) NOT NULL,
  "author_id" varchar(7)  NOT NULL,
  "owner_id" varchar(7) NOT NULL,
  "category" int,
  "listing_type" varchar(7) DEFAULT 'fixed',
  "floor_bid" numeric(10, 2),
  "end_time" timestamp,
  "status" varchar(20) DEFAULT 'available',
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "products_categories" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "nft_id"  varchar(7) NOT NULL,
  "category_id" int NOT NULL
);

CREATE TABLE "images" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "filename" varchar NOT NULL,
  "url" varchar NOT NULL,
  "nft_id" varchar(7) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "transactions" (
  "transaction_id" serial PRIMARY KEY,
  "buyer_id" varchar(7) NOT NULL,
  "id" varchar(7) NOT NULL,
  "gateway_id" varchar(20) DEFAULT 'Razorpay',
  "transaction_amount" numeric(10, 2) NOT NULL,
  "transaction_status" varchar(20) DEFAULT 'completed',
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "carts" (
  "user_id" varchar(7) NOT NULL,
  "nft_id" varchar(7) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "wishlists" (
  "user_id" varchar(7) NOT NULL,
  "nft_id" varchar(7) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "bids" (
  "bid_id" serial PRIMARY KEY,
  "user_id" varchar(7) NOT NULL,
  "nft_id" varchar(7) NOT NULL,
  "bid_amount" numeric(10, 2) NOT NULL,
  "bid_status" varchar(20) DEFAULT 'active',
  "bid_expiry" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "won_bids" (
  "bid_id" int,
  "id" varchar(7) NOT NULL,
  "nft_id" varchar(7) NOT NULL,
  "bid_amount" numeric(10, 2) NOT NULL,
  "bid_timestamp" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "lost_bids" (
  "bid_id" int,
  "user_id" varchar(7) NOT NULL,
  "nft_id" varchar(7) NOT NULL,
  "bid_amount" numeric(10, 2) NOT NULL,
  "bid_timestamp" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "wallet_transactions" (
  "transaction_id" serial PRIMARY KEY,
  "id" varchar(7) NOT NULL,
  "transaction_type" varchar(20) NOT NULL,
  "transaction_amount" numeric(10, 2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "messages" (
  "message_id" serial PRIMARY KEY,
  "sender_id" varchar(7) NOT NULL,
  "receiver_id" varchar(7) NOT NULL,
  "content" text NOT NULL,
  "read_status" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "blogs" (
  "blog_id" serial PRIMARY KEY,
  "id" varchar(7) NOT NULL,
  "title" varchar(255) NOT NULL,
  "content" text,
  "image_url" varchar(255),
  "category_id" int,
  "status" varchar(20) DEFAULT 'draft',
  "published_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "blog_categories" (
  "category_id" serial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "update_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE UNIQUE INDEX ON "carts" ("user_id", "nft_id");
CREATE UNIQUE INDEX ON "wishlists" ("user_id", "nft_id");
CREATE UNIQUE INDEX ON "won_bids" ("id", "nft_id");
CREATE UNIQUE INDEX ON "lost_bids" ("user_id", "nft_id");
CREATE INDEX "user_blogs" ON "blogs" ("id", "created_at");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "nfts" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");
ALTER TABLE "nfts" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
ALTER TABLE "nfts" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("id") REFERENCES "nfts" ("id");
ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "carts" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "wishlists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "wishlists" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "bids" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "bids" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "won_bids" ADD FOREIGN KEY ("bid_id") REFERENCES "bids" ("bid_id");
ALTER TABLE "won_bids" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "won_bids" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "lost_bids" ADD FOREIGN KEY ("bid_id") REFERENCES "bids" ("bid_id");
ALTER TABLE "lost_bids" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "lost_bids" ADD FOREIGN KEY ("nft_id") REFERENCES "nfts" ("id");
ALTER TABLE "wallet_transactions" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "messages" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");
ALTER TABLE "messages" ADD FOREIGN KEY ("receiver_id") REFERENCES "users" ("id");
ALTER TABLE "blogs" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "blogs" ADD FOREIGN KEY ("category_id") REFERENCES "blog_categories" ("category_id");

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_nfts_updated_at BEFORE UPDATE ON "nfts" FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_categories_updated_at BEFORE UPDATE ON "categories" FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_products_categories_updated_at BEFORE UPDATE ON "products_categories" FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_images_updated_at BEFORE UPDATE ON "images" FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

COMMIT;