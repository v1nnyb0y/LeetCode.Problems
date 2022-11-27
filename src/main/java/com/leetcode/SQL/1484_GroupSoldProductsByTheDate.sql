/* Write your T-SQL query statement below */
SELECT sell_date,COUNT(product) AS num_sold, string_agg( product,',') WITHIN GROUP ( ORDER BY product) as products
FROM
(SELECT DISTINCT(product),sell_date FROM Activities
group by sell_date , product) as dateAndDistinctProduct
group by sell_date
