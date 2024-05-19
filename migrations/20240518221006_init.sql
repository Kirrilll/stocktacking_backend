-- +goose Up
-- +goose StatementBegin
CREATE TABLE configurations (
                                id int NOT NULL,
                                code text NOT NULL,
                                name text NOT NULL,
                                PRIMARY KEY (id)
);

CREATE TABLE organizations (
                               id int NOT NULL,
                               name text NOT NULL,
                               phone text NOT NULL,
                               inn text NOT NULL,
                               PRIMARY KEY (id)
);

CREATE TABLE branches (
                          id int NOT NULL,
                          name varchar(255) NOT NULL,
                          organization_id integer NOT NULL,
                          address integer NOT NULL,
                          lat text NOT NULL,
                          lon text NOT NULL,
                          PRIMARY KEY (id),
                          CONSTRAINT branches_organization_id_organizations_id_foreign FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE TABLE categories (
                            id int NOT NULL,
                            name varchar(255) NOT NULL,
                            conf_id integer DEFAULT NULL,
                            PRIMARY KEY (id),
                            CONSTRAINT categories_conf_id_configurations_id_foreign FOREIGN KEY (conf_id) REFERENCES configurations (id)
);

CREATE TABLE options (
                         id int NOT NULL,
                         name varchar(255) NOT NULL,
                         code varchar(255) NOT NULL,
                         type text NOT NULL,
                         value text NOT NULL,
                         conf_id int NOT NULL,
                         PRIMARY KEY (id),
                         CONSTRAINT options_conf_id_configurations_id_foreign FOREIGN KEY (conf_id) REFERENCES configurations (id)
);

CREATE TABLE users (
                       id int NOT NULL,
                       name text NOT NULL,
                       role text NOT NULL,
                       phone text NOT NULL,
                       organization_id int NOT NULL,
                       PRIMARY KEY (id),
                       CONSTRAINT users_organization_id_organizations_id_foreign FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE TABLE storages (
                          id int NOT NULL,
                          name text NOT NULL,
                          parent_storage_id integer DEFAULT NULL,
                          branch_id int NOT NULL,
                          PRIMARY KEY (id),
                          CONSTRAINT storage_branch_id_branches_id_foreign FOREIGN KEY (branch_id) REFERENCES branches (id),
                          CONSTRAINT storage_storage_id_storage_id_foreign FOREIGN KEY (parent_storage_id) REFERENCES storages (id)
);

CREATE TABLE items (
                       id int NOT NULL,
                       name varchar(255) NOT NULL,
                       description text DEFAULT NULL,
                       photo_link text DEFAULT NULL,
                       storage_id int DEFAULT NULL,
                       user_id int DEFAULT NULL,
                       comment text DEFAULT NULL,
                       status varchar(255) NOT NULL,
                       category_id int DEFAULT NULL,
                       branch_id int DEFAULT NULL,
                       PRIMARY KEY (id),
                       CONSTRAINT items_storage_id_storage_id_foreign FOREIGN KEY (storage_id) REFERENCES storages (id),
                       CONSTRAINT items_user_id_users_id_foreign FOREIGN KEY (user_id) REFERENCES users (id),
                       CONSTRAINT items_category_id_categories_id_foreign FOREIGN KEY (category_id) REFERENCES categories (id),
                       CONSTRAINT items_branch_id_branches_id_foreign FOREIGN KEY (branch_id) REFERENCES branches (id)
);

CREATE TABLE reports (
                         id int NOT NULL,
                         created_at date NOT NULL,
                         item_id int NOT NULL,
                         closed_at date DEFAULT NULL,
                         satus_start text NOT NULL,
                         comment_start text NOT NULL,
                         user_start_id int NOT NULL,
                         status_end text DEFAULT NULL,
                         comment_end text DEFAULT NULL,
                         user_end_id int DEFAULT NULL,
                         storage_id int DEFAULT NULL,
                         PRIMARY KEY (id),
                         CONSTRAINT status_records_item_id_items_id_foreign FOREIGN KEY (item_id) REFERENCES items (id),
                         CONSTRAINT reports_user_end_id_users_id_foreign FOREIGN KEY (user_end_id) REFERENCES users (id),
                         CONSTRAINT reports_user_start_id_users_id_foreign FOREIGN KEY (user_start_id) REFERENCES users (id),
                         CONSTRAINT reports_storage_id_storages_id_foreign FOREIGN KEY (storage_id) REFERENCES storages (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE branches DROP CONSTRAINT branches_organization_id_organizations_id_foreign;
ALTER TABLE categories DROP CONSTRAINT categories_conf_id_configurations_id_foreign;
ALTER TABLE options DROP CONSTRAINT options_conf_id_configurations_id_foreign;
ALTER TABLE users DROP CONSTRAINT users_organization_id_organizations_id_foreign;
ALTER TABLE storages DROP CONSTRAINT storage_branch_id_branches_id_foreign;
ALTER TABLE storages DROP CONSTRAINT storage_storage_id_storage_id_foreign;
ALTER TABLE items DROP CONSTRAINT items_storage_id_storage_id_foreign;
ALTER TABLE items DROP CONSTRAINT items_user_id_users_id_foreign;
ALTER TABLE items DROP CONSTRAINT items_category_id_categories_id_foreign;
ALTER TABLE items DROP CONSTRAINT items_branch_id_branches_id_foreign;
ALTER TABLE reports DROP CONSTRAINT status_records_item_id_items_id_foreign;
ALTER TABLE reports DROP CONSTRAINT reports_user_end_id_users_id_foreign;
ALTER TABLE reports DROP CONSTRAINT reports_user_start_id_users_id_foreign;
ALTER TABLE reports DROP CONSTRAINT reports_storage_id_storages_id_foreign;
DROP TABLE configurations;
DROP TABLE organizations;
DROP TABLE branches;
DROP TABLE categories;
DROP TABLE options;
DROP TABLE users;
DROP TABLE storages;
DROP TABLE items;
DROP TABLE reports;
-- +goose StatementEnd
