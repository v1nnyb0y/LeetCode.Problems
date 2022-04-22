SELECT product_id, product_name
FROM Product
WHERE product_id not in (
    SELECT product_id
    FROM Sales
    WHERE sale_date NOT BETWEEN '2019-01-01' and '2019-04-01'
)
