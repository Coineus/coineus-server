CREATE TABLE users (
    hash_id   VARCHAR(500) NOT NULL,
    username    VARCHAR(100) NOT NULL,
    created_at     DATE NOT NULL,
    password_hash  VARCHAR(500) NOT NULL,
    email  VARCHAR(300) NOT NULL,
   CONSTRAINT users PRIMARY KEY (hash_id)
);

CREATE TABLE wallets (
    hash_id   VARCHAR(500) NOT NULL,
    user_id   VARCHAR(500) NOT NULL,
    name   VARCHAR(100) NOT NULL,
   CONSTRAINT wallets PRIMARY KEY (hash_id)
);

CREATE TABLE wallet_operations (
    hash_id   VARCHAR(500) NOT NULL,
    wallet_id  VARCHAR(500) NOT NULL,
    operation_id VARCHAR(500) NOT NULL,
   CONSTRAINT wallet_operations PRIMARY KEY (hash_id)
);

CREATE TABLE archived_operations (
    hash_id   VARCHAR(500) NOT NULL,
    user_id  VARCHAR(500) NOT NULL,
    archived_at DATE NOT NULL,
    coin_symbol VARCHAR(100) NOT NULL,
    coin_amount DOUBLE PRECISION NOT NULL,
    buy_cost DOUBLE PRECISION NOT NULL,
    sell_cost DOUBLE PRECISION NOT NULL,
    CONSTRAINT archived_operations PRIMARY KEY (hash_id)
);


CREATE TABLE recent_operations (
    hash_id   VARCHAR(500) NOT NULL,
    user_id  VARCHAR(500) NOT NULL,
    created_at DATE NOT NULL,
    coin_symbol VARCHAR(100) NOT NULL,
    coin_amount DOUBLE PRECISION NOT NULL,
    buy_cost DOUBLE PRECISION NOT NULL,
    CONSTRAINT recent_operations PRIMARY KEY (hash_id)
);

