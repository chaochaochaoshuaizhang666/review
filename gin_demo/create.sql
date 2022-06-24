CREATE SEQUENCE review_id_seq START 1;
create table review(
    id int4 PRIMARY KEY DEFAULT nextval('review_id_seq'::regclass),
    title varchar,
    create_time TIMESTAMP,
    content varchar
)