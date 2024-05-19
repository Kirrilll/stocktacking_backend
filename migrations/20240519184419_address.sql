-- +goose Up
-- +goose StatementBegin
alter table branches drop column address;
alter table branches add column address text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table branches drop column address;
alter table branches add column address integer;
-- +goose StatementEnd
