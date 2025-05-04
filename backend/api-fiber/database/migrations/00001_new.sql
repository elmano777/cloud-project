-- +goose Up
CREATE TABLE cursos (
    codigo INT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    horario VARCHAR(50) NOT NULL,
    ciclo VARCHAR(20) NOT NULL, -- Example: 2023-1, 2023-2
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE estudiante_cursos (
    estudiante_id INT NOT NULL,
    curso_codigo INT NOT NULL,
    inscrito_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (estudiante_id, curso_codigo),
    FOREIGN KEY (curso_codigo) REFERENCES cursos (codigo) ON DELETE CASCADE
);

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS estudiante_cursos;

DROP TABLE IF EXISTS cursos;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd
