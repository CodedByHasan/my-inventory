CREATE TABLE IF NOT EXISTS products (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    quantity INT DEFAULT NULL,
    price FLOAT(10,7) DEFAULT NULL 
);

INSERT INTO products (name, quantity, price) VALUES
    ('chair', 100, 200.0),
    ('desk', 20, 500.0);
