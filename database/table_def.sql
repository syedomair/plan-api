
drop table public.users CASCADE;
CREATE TABLE public.users (
        id                  uuid PRIMARY KEY,
        first_name          varchar(100) NOT NULL,
        last_name           varchar(100) NOT NULL,
        email               varchar(100) NOT NULL,
        password            varchar(100) NOT NULL,
        verified            smallint NULL DEFAULT 0,
        is_admin            smallint NULL DEFAULT 0,
        created_at          timestamp,
        updated_at          timestamp
);
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f6eebf36-fd66-4a71-a32f-8ea36d4617ef', 'FirstAdmin1', 'LastAdmin1', 'admin@gmail.com', 'MTIzNDU2', 1, '2018-09-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a72-a32f-8ea36d4617ef', 'FirstAdmin2', 'lastAdmin2', 'admin2@gmail.com', 'MTIzNDU2', 1, '2019-01-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a73-a32f-8ea36d4617ef', 'FirstAdmin3', 'lastAdmin3', 'admin3@gmail.com', 'MTIzNDU2', 1, '2019-01-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a74-a32f-8ea36d4617ef', 'FirstAdmin4', 'lastAdmin4', 'admin4@gmail.com', 'MTIzNDU2', 1, '2018-06-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a75-a32f-8ea36d4617ef', 'FirstAdmin5', 'lastAdmin5', 'admin5@gmail.com', 'MTIzNDU2', 1, '2018-07-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a76-a32f-8ea36d4617ef', 'FirstAdmin6', 'lastAdmin6', 'admin6@gmail.com', 'MTIzNDU2', 1, '2018-08-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a77-a32f-8ea36d4617ef', 'FirstAdmin7', 'lastAdmin7', 'admin7@gmail.com', 'MTIzNDU2', 1, '2018-09-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a78-a32f-8ea36d4617ef', 'FirstAdmin8', 'lastAdmin8', 'admin8@gmail.com', 'MTIzNDU2', 1, '2018-10-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a79-a32f-8ea36d4617ef', 'FirstAdmin9', 'lastAdmin9', 'admin9@gmail.com', 'MTIzNDU2', 1, '2018-10-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a71-a32f-8ea46d4617ef', 'FirstAdmin10', 'lastAdmin10', 'admin10@gmail.com', 'MTIzNDU2', 1, '2018-11-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a71-a32f-8ea56d4617ef', 'FirstAdmin11', 'lastAdmin11', 'admin11@gmail.com', 'MTIzNDU2', 1, '2018-11-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a71-a32f-8ea66d4617ef', 'FirstAdmin12', 'lastAdmin12', 'admin12@gmail.com', 'MTIzNDU2', 1, '2018-11-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a71-a32f-8ea76d4617ef', 'FirstAdmin13', 'lastAdmin13', 'admin13@gmail.com', 'MTIzNDU2', 1, '2018-11-14');
INSERT INTO public.users (id, first_name, last_name, email, password, verified, created_at) VALUES ('f7eebf36-fd66-4a71-a32f-8ea86d4617ef', 'FirstAdmin14', 'lastAdmin14', 'admin14@gmail.com', 'MTIzNDU2', 1, '2018-12-14');

drop table public.users_logins CASCADE;
CREATE TABLE public.users_logins(
        user_id             uuid PRIMARY KEY,
        token               varchar(500) NOT NULL,
        created_at          timestamp
);


drop table public.plans CASCADE;
CREATE TABLE public.plans(
        id                  uuid PRIMARY KEY,
        title               varchar(100) NULL,
        status              smallint NULL DEFAULT 0,
        validity            smallint NULL DEFAULT 0,
        cost                smallint NULL DEFAULT 0,
        created_at          timestamp,
        updated_at          timestamp
);

INSERT INTO public.plans (id, title, status, validity, cost) VALUES ('d107aa5c-9995-47b2-b34a-203ad655b621', 'Monthly Plan', 1, 30, 9999);
INSERT INTO public.plans (id, title, status, validity, cost) VALUES ('c9de5200-dbad-44b8-b5fc-ab1381730de7', 'Weekly Plan', 1, 7, 4999);
INSERT INTO public.plans (id, title, status, validity, cost) VALUES ('7be49965-fc69-4d15-ae53-aebcd7367402', 'Daily Plan', 1, 1, 1999);


drop table public.plan_messages CASCADE;
CREATE TABLE public.plan_messages (
        id                  uuid PRIMARY KEY,
        plan_id             uuid references public.plans(id) ON DELETE CASCADE,
        message             varchar(2000) NULL,
        action              varchar(20) NULL,
        created_at          timestamp,
        updated_at          timestamp
);

INSERT INTO public.plan_messages (id, plan_id, message, action) VALUES ('aeb9844a-b3fb-4df8-b684-7c4cf1f3ce32', 'd107aa5c-9995-47b2-b34a-203ad655b621','Message to notify that plan: #PLAN# cost has been updated. Here is the new cost: #COST#', 'COST_UPDATE');
INSERT INTO public.plan_messages (id, plan_id, message, action) VALUES ('faada022-40ff-4647-8ef9-542315570d61', 'd107aa5c-9995-47b2-b34a-203ad655b621','Message to notify that plan: #PLAN# validity has been updated. Here is the new validity: #VALIDITY#', 'VALIDITY_UPDATE');

INSERT INTO public.plan_messages (id, plan_id, message, action) VALUES ('64132a23-1c45-4344-8a8e-4633845a4ace', 'c9de5200-dbad-44b8-b5fc-ab1381730de7','Message to notify that plan: #PLAN# cost has been updated. Here is the new cost: #COST#', 'COST_UPDATE');
INSERT INTO public.plan_messages (id, plan_id, message, action) VALUES ('7eae9c25-a9f6-4bda-a7ef-a2dd9585a5fc', 'c9de5200-dbad-44b8-b5fc-ab1381730de7','Message to notify that plan: #PLAN# validity has been updated. Here is the new validity: #VALIDITY#', 'VALIDITY_UPDATE');

INSERT INTO public.plan_messages (id, plan_id, message, action) VALUES ('108f470a-93c3-41e1-a52b-20a3d344f9d6', '7be49965-fc69-4d15-ae53-aebcd7367402','Message to notify that plan: #PLAN# cost has been updated. Here is the new cost: #COST#', 'COST_UPDATE');
INSERT INTO public.plan_messages (id, plan_id, message, action) VALUES ('3afb692a-26bc-4c78-8f45-554ecf80fbf4', '7be49965-fc69-4d15-ae53-aebcd7367402','Message to notify that plan: #PLAN# validity has been updated. Here is the new validity: #VALIDITY#', 'VALIDITY_UPDATE');
