CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    telegram_username VARCHAR(255) NOT NULL,
    telegram_chat_id INT,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    grade INT NOT NULL,
    class_letter VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    referral VARCHAR(255) NOT NULL,
    by_referral VARCHAR(255) DEFAULT 'none'
);

