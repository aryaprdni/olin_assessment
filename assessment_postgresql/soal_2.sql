SELECT users.name, orders.amount, orders.created_at
FROM dblink(
  'dbname=first user=postgres password=root',
  'SELECT id, name FROM users'
) AS users(id INT, name TEXT)
JOIN dblink(
  'dbname=second user=postgres password=root',
  'SELECT user_id, amount, created_at FROM orders'
) AS orders(user_id INT, amount NUMERIC, created_at TIMESTAMP)
ON users.id = orders.user_id;
