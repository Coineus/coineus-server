CREATE TABLE IF NOT EXISTS users (
    hash_id   VARCHAR(500) NOT NULL PRIMARY KEY,
    username    VARCHAR(100) NOT NULL,
    created_at     DATE NOT NULL,
    password_hash  VARCHAR(500) NOT NULL,
    email  VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS wallets (
    hash_id   VARCHAR(500) NOT NULL PRIMARY KEY,
    user_id   VARCHAR(500) NOT NULL,
    name   VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS wallet_operations (
    hash_id   VARCHAR(500) NOT NULL PRIMARY KEY,
    wallet_id  VARCHAR(500) NOT NULL,
    operation_id VARCHAR(500) NOT NULL
);

CREATE TABLE IF NOT EXISTS archived_operations (
    hash_id   VARCHAR(500) NOT NULL PRIMARY KEY,
    user_id  VARCHAR(500) NOT NULL,
    archived_at DATE NOT NULL,
    coin_symbol VARCHAR(100) NOT NULL,
    coin_amount DOUBLE PRECISION NOT NULL,
    buy_cost DOUBLE PRECISION NOT NULL,
    sell_cost DOUBLE PRECISION NOT NULL
);


CREATE TABLE IF NOT EXISTS recent_operations (
    hash_id   VARCHAR(500) NOT NULL PRIMARY KEY,
    user_id  VARCHAR(500) NOT NULL,
    created_at DATE NOT NULL,
    coin_symbol VARCHAR(100) NOT NULL,
    coin_amount DOUBLE PRECISION NOT NULL,
    buy_cost DOUBLE PRECISION NOT NULL
);

