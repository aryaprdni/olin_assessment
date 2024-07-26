SELECT users.name, SUM(orders.amount) as total_spent
FROM users
JOIN orders on users.id = orders.user_id
where orders.created_at >= '2022-01-01'
GROUP BY users.name;