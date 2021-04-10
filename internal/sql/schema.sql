CREATE TABLE main (
    id serial PRIMARY KEY
    ip_id bigint FOREIGN KEY
    query_type_id bigint FOREIGN KEY
    query_url_id bigint FOREIGN KEY
    session_ids_id bigint FOREIGN KEY
)

CREATE TABLE ip_address(
    id serial PRIMARY KEY
    value string
)

CREATE TABLE query_type(
    id serial PRIMARY KEY
    value string
)

CREATE TABLE query_url(
    id serial PRIMARY KEY
    value string
)

CREATE TABLE session_ids(
    id serial PRIMARY KEY
    value string
)