CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    telegram_username VARCHAR(255) NOT NULL,
    telegram_chat_id INT,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    grade INT NOT NULL,
    class_letter VARCHAR(5) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    referral VARCHAR(255) NOT NULL,
    by_referral VARCHAR(255) DEFAULT 'none',
    balance INT NOT NULL DEFAULT 0
);

CREATE TABLE achievements (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    node_id INT NOT NULL,
    coins INT NOT NULL,
    text VARCHAR(255) NOT NULL,
    document_url VARCHAR(255) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('pending', 'approved', 'rejected')) DEFAULT 'pending',
    reject_comment VARCHAR(255) NOT NULL DEFAULT 'none'
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    photo_url VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE products_in_cart (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    product_amount INT NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status VARCHAR(20) CHECK (status IN ('pending', 'rejected', 'ready', 'issued')) DEFAULT 'pending'
);

CREATE TABLE ordered_products (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    product_amount INT NOT NULL,
    price INT NOT NULL
);

ALTER TABLE achievements ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE products_in_cart ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE products_in_cart ADD FOREIGN KEY (product_id) REFERENCES products(id);
ALTER TABLE orders ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE ordered_products ADD FOREIGN KEY (product_id) REFERENCES products(id);
ALTER TABLE ordered_products ADD FOREIGN KEY (order_id) REFERENCES orders(id);