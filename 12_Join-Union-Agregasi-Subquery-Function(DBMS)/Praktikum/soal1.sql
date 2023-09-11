-- INSERT
-- Insert 5 operators pada tabel operators:
INSERT INTO operator (name, created_at) VALUES
    ('Operator 1', NOW()),
    ('Operator 2', NOW()),
    ('Operator 3', NOW()),
    ('Operator 4', NOW()),
    ('Operator 5', NOW());

-- Insert 3 product types:
INSERT INTO product_types (name, created_at) VALUES
    ('baju', NOW()),
    ('celana', NOW()),
    ('sepatu', NOW());

-- Insert 2 product dengan product type id = 1, dan operators id = 3:
INSERT INTO product (product_type_id, operator_id, name, status, created_at) VALUES
    (1, 3, 'baju merah', 1, NOW()),
    (1, 3, 'baju hijau', 1, NOW());

-- Insert 3 product dengan product type id = 2, dan operators id = 1:
INSERT INTO product (product_type_id, operator_id, name, status, created_at) VALUES
    (2, 1, 'celana merah', 1, NOW()),
    (2, 1, 'celana putih', 1, NOW()),
    (2, 1, 'celana biru', 1, NOW());

-- Insert 3 product dengan product type id = 3, dan operators id = 4:
INSERT INTO product (product_type_id, operator_id, name, status, created_at) VALUES
    (3, 4, 'tas merah', 1, NOW()),
    (3, 4, 'tas putih', 1, NOW()),
    (3, 4, 'tas biru', 1, NOW());

-- Insert product description pada setiap product:
INSERT INTO product_descriptions (description, created_at) VALUES
    ('baju berwarna merah', NOW()),
    ('baju berwarna hijau', NOW()),
    ('baju berwarna putih', NOW()),
    ('celana berwarna merah', NOW()),
    ('celana berwarna putih', NOW()),
    ('celana berwarna biru', NOW()),
    ('tas berwarna merah', NOW()),
    ('tas berwarna putih', NOW()),
    ('tas berwarna biru', NOW());

-- Insert 3 payment methods:
INSERT INTO payment_methods (name, status, created_at) VALUES
    ('gopay', 1, NOW()),
    ('dana', 1, NOW()),
    ('bank', 1, NOW());

-- Insert 5 user pada tabel user:
INSERT INTO users (nama, alamat, status, dob, gender, created_at) VALUES
    ('tono', 'bogor', 1, '2000-01-01', 'M', NOW()),
    ('andini', 'bandung', 1, '1995-05-15', 'F', NOW()),
    ('topan', 'jakarta', 1, '1988-09-20', 'M', NOW()),
    ('sri', 'tangerang', 1, '2002-03-10', 'F', NOW()),
    ('budi', 'depok', 1, '1990-11-30', 'M', NOW());

-- Insert 3 transaksi di masing-masing user:
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price, created_at) VALUES
    -- Transaksi untuk User 1
    (1, 1, 'Completed', 3, 150.00, NOW()),
    (1, 2, 'Pending', 3, 150000.00, NOW()),
    (1, 3, 'Completed', 3, 150000.00, NOW()),

    -- Transaksi untuk User 2
    (2, 1, 'Pending', 3, 150000.00, NOW()),
    (2, 2, 'Pending', 3, 150000.00, NOW()),
    (2, 3, 'Pending', 3, 150000.00, NOW()),

    -- Transaksi untuk User 3
    (3, 1, 'Pending', 3, 150000.00, NOW()),
    (3, 2, 'Pending', 3, 150000.00, NOW()),
    (3, 3, 'Completed', 3, 150000.00, NOW());

-- Insert 3 produk di masing-masing transaksi:
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price, created_at) VALUES
    -- Produk untuk Transaksi User 1, Transaksi 1
    (1, 1, 'Completed', 1, 50000.00, NOW()),
    (1, 2, 'Completed', 1, 50000.00, NOW()),
    (1, 3, 'Completed', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 1, Transaksi 2
    (2, 1, 'Pending', 1, 50000.00, NOW()),
    (2, 2, 'Pending', 1, 50000.00, NOW()),
    (2, 3, 'Pending', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 1, Transaksi 3
    (3, 1, 'Pending', 1, 50000.00, NOW()),
    (3, 2, 'Completed', 1, 50000.00, NOW()),
    (3, 3, 'Completed', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 2, Transaksi 1
    (4, 4, 'Completed', 1, 100.00, NOW()),
    (4, 5, 'Pending', 1, 50.00, NOW()),
    (4, 6, 'Pending', 1, 150.00, NOW()),

    -- Produk untuk Transaksi User 2, Transaksi 2
    (5, 1, 'Completed', 1, 50000.00, NOW()),
    (5, 2, 'Completed', 1, 50000.00, NOW()),
    (5, 3, 'Completed', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 2, Transaksi 3
    (6, 4, 'Pending', 1, 50000.00, NOW()),
    (6, 5, 'Pending', 1, 50000.00, NOW()),
    (6, 6, 'Pending', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 3, Transaksi 1
    (7, 7, 'Pending', 1, 50000.00, NOW()),
    (7, 8, 'Pending', 1, 50000.00, NOW()),
    (7, 9, 'Completed', 1, 50000.00, NOW()),

-- Produk untuk Transaksi User 3, Transaksi 2
    (8, 7, 'Completed', 1, 50000.00, NOW()),
    (8, 8, 'Completed', 1, 50000.00, NOW()),
    (8, 9, 'Pending', 1, 50000.00, NOW()),

-- Produk untuk Transaksi User 3, Transaksi 3
    (9, 1, 'Completed', 1, 50000.00, NOW()),
    (9, 2, 'Completed', 1, 50000.00, NOW()),
    (9, 3, 'Completed', 1, 50000.00, NOW());
SELECT * FROM transaction_details;

-- SELECT

-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT id FROM users WHERE gender = 'M';

-- Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id = 3;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM users WHERE created_at >= date_sub(NOW(), INTERVAL 7 DAY) AND nama LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) FROM users WHERE gender = 'F';

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users order by nama ASC;

-- Tampilkan 5 data pada data product
select * from products LIMIT 5;

-- UPDATE 
-- Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE products set name ='product dummy' WHERE id =1;
select * from products;

-- Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;
select * from products;

-- DELETE
-- Delete data pada tabel product dengan id = 1.
DELETE FROM products WHERE id=1;
select * from products;

-- Delete pada pada tabel product dengan product type id 1.
DELETE FROM products WHERE product_type_id = 1;
select * from products;