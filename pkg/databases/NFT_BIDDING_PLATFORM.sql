CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "role_id" int,
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
  "wallet_amount" "decimal(10, 2)" DEFAULT 0,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "roles" (
  "id" int PRIMARY KEY,
  "title" varchar(20)
);

CREATE TABLE "oauth" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "access_token" varchar(12),
  "refresh_token" varchar(12),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "nfts" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "description" text,
  "image_url" varchar(255) NOT NULL,
  "author_id" int,
  "owner_id" int,
  "category" int,
  "listing_type" varchar(7) DEFAULT 'fixed',
  "price" "decimal(10, 2)" NOT NULL,
  "floor_bid" "decimal(10, 2)",
  "end_time" timestamp,
  "status" varchar(20) DEFAULT 'available',
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "products_categories" (
  "id" serial PRIMARY KEY,
  "nft_id" varchar,
  "category_id" varchar
);

CREATE TABLE "images" (
  "id" varchar PRIMARY KEY,
  "filename" varchar,
  "url" varchar,
  "nft_id" varchar
);

CREATE TABLE "transactions" (
  "transaction_id" serial PRIMARY KEY,
  "buyer_id" int,
  "id" int,
  "gateway_id" varchar(20) DEFAULT 'Razorpay',
  "transaction_amount" "decimal(10, 2)" NOT NULL,
  "transaction_status" varchar(20) DEFAULT 'completed',
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "carts" (
  "user_id" int,
  "nft_id" int,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "wishlists" (
  "user_id" int,
  "nft_id" int,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "bids" (
  "bid_id" serial PRIMARY KEY,
  "user_id" int,
  "nft_id" int,
  "bid_amount" "decimal(10, 2)" NOT NULL,
  "bid_status" varchar(20) DEFAULT 'active',
  "bid_expiry" timestamp NOT NULL,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "won_bids" (
  "bid_id" int,
  "id" int,
  "nft_id" int,
  "bid_amount" "decimal(10, 2)" NOT NULL,
  "bid_timestamp" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "lost_bids" (
  "bid_id" int,
  "user_id" int,
  "nft_id" int,
  "bid_amount" "decimal(10, 2)" NOT NULL,
  "bid_timestamp" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "wallet_transactions" (
  "transaction_id" serial PRIMARY KEY,
  "id" int,
  "transaction_type" varchar(20) NOT NULL,
  "transaction_amount" "decimal(10, 2)" NOT NULL,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "messages" (
  "message_id" serial PRIMARY KEY,
  "sender_id" int,
  "receiver_id" int,
  "content" text NOT NULL,
  "read_status" boolean DEFAULT false,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "blogs" (
  "blog_id" serial PRIMARY KEY,
  "id" int,
  "title" varchar(255) NOT NULL,
  "content" text,
  "image_url" varchar(255),
  "category_id" int,
  "status" varchar(20) DEFAULT 'draft',
  "published_at" timestamp,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "blog_categories" (
  "category_id" serial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "update_at" timestamp,
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

ALTER TABLE "bids" ADD FOREIGN KEY ("bid_status") REFERENCES "bids" ("bid_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("transaction_status") REFERENCES "transactions" ("created_at");
