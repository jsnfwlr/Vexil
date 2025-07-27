INSERT into environment (name) VALUES ('dev');

INSERT into feature_flag (name, default_value, value_type)
VALUES ('FIRST_STRING', '30cm red three-strand nylon', 'string'),
       ('SECOND_STRING', '30cm green six-strand cotton', 'string'),
       ('THIRD_STRING', '30cm blue five-strand wool', 'string');

INSERT into feature_flag_value (feature_flag_id, environment_id, value)
VALUES (1, 1, '25cm red three-strand nylon'),
       (2, 1, '25cm green six-strand cotton'),
       (3, 1, '25cm blue five-strand wool');

