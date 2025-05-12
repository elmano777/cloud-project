-- 1
SELECT *
FROM "proyecto-cloud"."ejerce" A
INNER JOIN "proyecto-cloud"."notadata" B
ON A.curso_id = B.codigo_curso
LIMIT 10;

-- 2
SELECT *
FROM "proyecto-cloud"."ejerce" A
INNER JOIN "proyecto-cloud"."cursosdata" B
ON A.curso_id = B.codigo
LIMIT 10
;

-- 3
SELECT DISTINCT B.profesor_id
FROM "proyecto-cloud"."estudiantecursosdata" A 
INNER JOIN "proyecto-cloud"."ejerce" B
ON B.curso_id = A.curso_codigo
LIMIT 1

-- 4
