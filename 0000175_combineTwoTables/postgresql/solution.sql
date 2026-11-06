SELECT
    p.firstname AS "firstName",
    p.lastname AS "lastName",
    a.city,
    a.state
FROM person AS p
LEFT JOIN address AS a
    ON p.personid = a.personid;
