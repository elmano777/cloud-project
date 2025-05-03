-- +goose Up
-- +goose StatementBegin
CREATE TABLE TABLA1 ( 
    codigo VARCHAR(50) PRIMARY KEY, 
    nombre VARCHAR(100), 
    horario VARCHAR(50), 
    ciclo VARCHAR(20) 
);

CREATE TABLE TABLA2 ( 
    id_estudiante INT, 
    codigo_curso VARCHAR(50), 
    FOREIGN KEY (codigo_curso) REFERENCES TABLA1(codigo) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
