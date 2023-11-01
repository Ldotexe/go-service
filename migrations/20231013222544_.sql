-- +goose Up
-- +goose StatementBegin
create table students(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL,
    points int not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table students;
-- +goose StatementEnd

