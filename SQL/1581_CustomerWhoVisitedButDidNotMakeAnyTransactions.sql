# Write your MySQL query statement below
SELECT V.customer_id, count(V.customer_id) AS count_no_trans FROM VISITS V
LEFT JOIN Transactions T ON T.visit_id = V.visit_ID
WHERE T.transaction_id IS NULL 
GROUP BY V.customer_id
