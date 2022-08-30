CREATE DATABASE review_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE reviews (
	review_id serial4 NOT NULL,
	review_hash_id UUID NOT NULL DEFAULT uuid_generate_v1(),
	review_item_id VARCHAR (38) NOT NULL,
	CONSTRAINT reviews_pkey PRIMARY KEY (review_id),
	UNIQUE(review_hash_id, review_item_id)
);

CREATE TABLE review_comments (
	review_comments_id serial4 NOT NULL,
	review_id serial4 NOT NULL,
	review_comment VARCHAR (200) NOT NULL,
	review_customer_id VARCHAR (58) NOT NULL,
	review_order_id VARCHAR (100),
	review_datatime timestamp DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT review_comments_pkey PRIMARY KEY (review_comments_id)
	CONSTRAINT fk_reviews FOREIGN key (review_id) REFERENCES reviews (review_id)
);