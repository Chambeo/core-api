CREATE TABLE public.Users (
                       id varchar(250) primary key,
                       email varchar(250) not null,
                       password varchar(250) not null,
                       first_name varchar(250) not null,
                       last_name varchar(250) not null,
                       phone_code varchar(10),
                       phone_number varchar(25),
                       address_street varchar(250) not null,
                       address_number varchar(10),
                       address_details varchar(150),
                       zip_code varchar(25),
                       province varchar(200),
                       city varchar(250),
                       role varchar(15),
                       account_verified boolean default false,
                       agree_terms_and_conditions boolean default false,
                       profile_completed boolean default false,
                       created_at timestamp without time zone NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at timestamp without time zone NULL,
                       deleted_at timestamp without time zone NULL
);



--ALTER TABLE user_lists ADD CONSTRAINT user_lists_list_id_fk FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE RESTRICT ON UPDATE RESTRICT;
--ALTER TABLE user_lists ADD CONSTRAINT user_lists_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT;


--TODO cambiar las tablas por las correspondientes de la forma mas noSql posible