CREATE TYPE project_type AS ENUM (
    'КГ',
    'ШТ'
);

CREATE TABLE projects (
  id serial primary key,
  name varchar(50),
  cost float,
  dimension varchar(255),
  type project_type
);