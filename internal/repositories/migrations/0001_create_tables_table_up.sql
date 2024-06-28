CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255),
    customer_email VARCHAR(255) UNIQUE,
    customer_address TEXT 
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) ,
    category VARCHAR(255)
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES customers(customer_id),
    product_id INT NOT NULL REFERENCES products(product_id),
    date_of_sale DATE,
    quantity_sold INT,
    unit_price DECIMAL(10, 2),
    discount DECIMAL(5, 2),
    shipping_cost DECIMAL(10, 2),
    payment_method VARCHAR(255),
    region VARCHAR(255)
);
