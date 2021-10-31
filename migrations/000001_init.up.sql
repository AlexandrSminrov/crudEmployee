CREATE TABLE IF NOT EXISTS employees (
    id                  serial4,
    first_name       VARCHAR (255) NOT NULL,
    last_name        VARCHAR (255) NOT NULL,
    middle_name      VARCHAR (255) NOT NULL,
    date_of_birth   DATE NOT NULL,
    address         VARCHAR (3000) NOT NULL,
    department     VARCHAR (255) NOT NULL,
    about_me         VARCHAR(3000),
    phone         VARCHAR(20) NOT NULL,
    email           VARCHAR(320) NOT NULL
    )
-- INSERT INTO public.emploees (firstname, lastname, middlename, date_of_birth,
--                              addres, department, about_me, phone, email)
--     VALUES ('Иванов', 'Иван', 'Иванович', '10.11.1980', 'Москва', 'HR',
--             'qwe', '9000000000', 'exaple@ex.ex') RETURNING id;
--
-- INSERT INTO public.emploees (firstname, lastname, middlename, date_of_birth,
--                              addres, department, about_me, phone, email)
--     VALUES ('Иванов', 'Иван', 'Иванович', '10.11.1999', 'Москва', 'разраб',
--