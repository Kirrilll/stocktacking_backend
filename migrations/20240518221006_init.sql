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
                            parent_category_id int DEFAULT NULL,
                            conf_id integer DEFAULT NULL,
                            PRIMARY KEY (id),
                            CONSTRAINT products_category_parent_category_id_products_foreign FOREIGN KEY (parent_category_id) REFERENCES categories (id),
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
                          capacity int NOT NULL,
                          address text NOT NULL,
                          frame integer NOT NULL,
                          PRIMARY KEY (id),
                          CONSTRAINT storage_branch_id_branches_id_foreign FOREIGN KEY (branch_id) REFERENCES branches (id),
                          CONSTRAINT storage_storage_id_storage_id_foreign FOREIGN KEY (parent_storage_id) REFERENCES storages (id)
);

CREATE TABLE items (
                       id int NOT NULL,
                       name varchar(255) NOT NULL,
                       description text DEFAULT NULL,
                       photo_link text DEFAULT NULL,
                       category text DEFAULT NULL,
                       storage_id int DEFAULT NULL,
                       user_id int DEFAULT NULL,
                       comment text DEFAULT NULL,
                       status varchar(255) NOT NULL,
                       branches_id int NOT NULL,
                       qr_img text DEFAULT NULL,
                       PRIMARY KEY (id),
                       CONSTRAINT items_storage_id_storage_id_foreign FOREIGN KEY (storage_id) REFERENCES storages (id),
                       CONSTRAINT items_user_id_users_id_foreign FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE items_categories (
                                  category_id int NOT NULL,
                                  item_id int NOT NULL,
                                  PRIMARY KEY (category_id, item_id),
                                  CONSTRAINT items_categories_category_id_categories_id_foreign FOREIGN KEY (category_id) REFERENCES categories (id),
                                  CONSTRAINT items_categories_item_id_items_id_foreign FOREIGN KEY (item_id) REFERENCES items (id)
);

CREATE TABLE reports (
                         id int NOT NULL,
                         text text NOT NULL,
                         created_at date NOT NULL,
                         new_status varchar(255) NOT NULL,
                         item_id int NOT NULL,
                         user_id int NOT NULL,
                         closed_at date DEFAULT NULL,
                         PRIMARY KEY (id),
                         CONSTRAINT status_records_item_id_items_id_foreign FOREIGN KEY (item_id) REFERENCES items (id),
                         CONSTRAINT status_records_user_id_users_id_foreign FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE actions (
                         id int NOT NULL,
                         name text NOT NULL,
                         description text DEFAULT NULL,
                         report_id int NOT NULL,
                         date date NOT NULL,
                         PRIMARY KEY (id),
                         CONSTRAINT actions_report_id_reports_id_foreign FOREIGN KEY (report_id) REFERENCES reports (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE branches DROP CONSTRAINT branches_organization_id_organizations_id_foreign;
ALTER TABLE categories DROP CONSTRAINT products_category_parent_category_id_products_foreign;
ALTER TABLE categories DROP CONSTRAINT categories_conf_id_configurations_id_foreign;
ALTER TABLE options DROP CONSTRAINT options_conf_id_configurations_id_foreign;
ALTER TABLE users DROP CONSTRAINT users_organization_id_organizations_id_foreign;
ALTER TABLE storages DROP CONSTRAINT storage_branch_id_branches_id_foreign;
ALTER TABLE storages DROP CONSTRAINT storage_storage_id_storage_id_foreign;
ALTER TABLE items DROP CONSTRAINT items_storage_id_storage_id_foreign;
ALTER TABLE items DROP CONSTRAINT items_user_id_users_id_foreign;
ALTER TABLE items_categories DROP CONSTRAINT items_categories_category_id_categories_id_foreign;
ALTER TABLE items_categories DROP CONSTRAINT items_categories_item_id_items_id_foreign;
ALTER TABLE reports DROP CONSTRAINT status_records_item_id_items_id_foreign;
ALTER TABLE reports DROP CONSTRAINT status_records_user_id_users_id_foreign;
ALTER TABLE actions DROP CONSTRAINT actions_report_id_reports_id_foreign;
DROP TABLE configurations;
DROP TABLE organizations;
DROP TABLE branches;
DROP TABLE categories;
DROP TABLE options;
DROP TABLE users;
DROP TABLE storages;
DROP TABLE items;
DROP TABLE items_categories;
DROP TABLE reports;
DROP TABLE actions;
-- +goose StatementEnd
