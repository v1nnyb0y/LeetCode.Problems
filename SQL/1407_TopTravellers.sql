# Write your MySQL query statement below
SELECT name, COALESCE(SUM(distance),0) as travelled_distance
FROM Users
LEFT JOIN Rides on Users.id = Rides.user_id
GROUP BY name
ORDER BY travelled_distance DESC, name
