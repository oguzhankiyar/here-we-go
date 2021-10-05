CREATE TABLE products (
    id              VARCHAR     PRIMARY KEY,
    name            VARCHAR     NOT NULL,
    price           DECIMAL     NOT NULL,
    status          INT         NOT NULL,
    is_deleted      BOOLEAN     NOT NULL,
    created_at      TIMESTAMP   NOT NULL,
    updated_at      TIMESTAMP   NULL,
    deleted_at      TIMESTAMP   NULL
);