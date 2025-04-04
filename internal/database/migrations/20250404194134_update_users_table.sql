-- +goose Up
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN username;
ALTER TABLE users DROP COLUMN phone;
ALTER TABLE users DROP COLUMN gender;
ALTER TABLE users DROP COLUMN birthday;
ALTER TABLE users DROP COLUMN location;
ALTER TABLE users DROP COLUMN religion;
ALTER TABLE users DROP COLUMN department;
ALTER TABLE users DROP COLUMN stream;
ALTER TABLE users DROP COLUMN degree;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users 
 ADD COLUMN username varchar(255) NOT NULL,
 ADD COLUMN phone varchar(255) NOT NULL,
 ADD COLUMN gender varchar(255) NOT NULL,
 ADD COLUMN birthday date NOT NULL,
 ADD COLUMN location varchar(255) NOT NULL,
 ADD COLUMN religion varchar(255) NOT NULL,
 ADD COLUMN department varchar(255) NOT NULL,
 ADD COLUMN stream varchar(255) NOT NULL,
 ADD COLUMN degree varchar(255) NOT NULL,
-- +goose StatementEnd
