-- 1
SELECT * FROM transactions WHERE user_id = 1
UNION
SELECT * FROM transactions WHERE user_id = 2;

-- 2
SELECT SUM(total_price) AS total_price FROM transactions WHERE user_id = 1;

-- 3
SELECT SUM(t.total_price) AS total_price, td.qty
FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
JOIN product_types pt ON p.product_type_id = pt.id
WHERE pt.id = 2
GROUP BY td.qty;

-- 4
SELECT products.*, product_types.name AS product_type_name FROM products
JOIN product_types ON products.product_type_id = product_types.id;

-- 5
SELECT t.*, u.nama, p.name
FROM transactions t
JOIN users u ON t.user_id = u.id
JOIN payment_methods pm ON t.payment_method_id = pm.id
JOIN transaction_details td ON td.transaction_id = t.id
JOIN products p ON td.product_id = p.id
JOIN product_types pt ON p.product_type_id = pt.id;
-- 6
delimiter //
create trigger after_transaction_delete
after delete on transactions
for each ROW
begin
  DELETE FROM transaction_details
  WHERE  transaction_id = OLD.id;
end;
//
delimiter;
-- 7
DELIMITER //
CREATE TRIGGER update_total_qty
AFTER DELETE ON transaction_details FOR EACH ROW
BEGIN
    DECLARE deleted_qty INT;
    SELECT OLD.qty INTO deleted_qty;
    UPDATE transactions
    SET total_qty = total_qty - deleted_qty
    WHERE id = OLD.transaction_id;
END;
//
DELIMITER ;

-- 8
SELECT * FROM products WHERE id NOT IN (SELECT DISTINCT product_id FROM transaction_details);