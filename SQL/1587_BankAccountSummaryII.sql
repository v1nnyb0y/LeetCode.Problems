SELECT name, SUM(amount) as balance
FROM Users
INNER JOIN Transactions on Transactions.account = Users.account
GROUP BY name
HAVING SUM(amount) > 10000
