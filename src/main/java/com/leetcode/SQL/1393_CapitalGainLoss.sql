# Write your MySQL query statement below
SELECT stock_name, SUM(CASE WHEN operation='Sell' then price else price*-1 end) as capital_gain_loss
FROM Stocks
GROUP BY stock_name
