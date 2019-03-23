
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
INSERT INTO public.users (id, first_name, last_name, email, password, verified) VALUES ('f6eebf36-fd66-4a71-a32f-8ea36d4617ef', 'FirstAdmin', 'LastAdmin', 'admin@gmail.com', 'MTIzNDU2', 1);
INSERT INTO public.users (id, first_name, last_name, email, password, verified) VALUES ('f7eebf36-fd66-4a71-a32f-8ea36d4617ef', 'firstApplicant', 'lastApplicant', 'applicant@gmail.com', 'MTIzNDU2', 1);
INSERT INTO public.users (id, first_name, last_name, email, password, verified) VALUES ('8ef90caf-69a8-4247-945f-978094901c8c', 'firstClerk', 'lastClerk', 'clerk@gmail.com', 'MTIzNDU2', 1);
INSERT INTO public.users (id, first_name, last_name, email, password, verified) VALUES ('ba8ea792-328a-4215-be33-69623bbeda57', 'FirstManager', 'LastManager', 'manager@gmail.com', 'MTIzNDU2', 1);
INSERT INTO public.users (id, first_name, last_name, email, password, verified) VALUES ('24fc0c41-75e6-43aa-9781-b66a0d7ae0dd', 'firstApplicant1', 'lastApplicant1', 'applicant1@gmail.com', 'MTIzNDU2', 1);

drop table public.users_logins CASCADE;
CREATE TABLE public.users_logins(
        user_id             uuid PRIMARY KEY,
        token               varchar(500) NOT NULL,
        created_at          timestamp
);


drop table public.plan CASCADE;
CREATE TABLE public.plan(
        id                  uuid PRIMARY KEY,
        title               varchar(100) NULL,
        status              smallint NULL DEFAULT 0,
        validity            smallint NULL DEFAULT 0,
        cost                smallint NULL DEFAULT 0,
        created_at          timestamp,
        updated_at          timestamp
);

INSERT INTO public.plan (id, title, status, validity, cost) VALUES ('d107aa5c-9995-47b2-b34a-203ad655b621', 'Monthly Plan', 1, 30, 9999);
INSERT INTO public.plan (id, title, status, validity, cost) VALUES ('c9de5200-dbad-44b8-b5fc-ab1381730de7', 'Weekly Plan', 1, 7, 4999);
INSERT INTO public.plan (id, title, status, validity, cost) VALUES ('7be49965-fc69-4d15-ae53-aebcd7367402', 'Daily Plan', 1, 1, 1999);


drop table public.plan_messages CASCADE;
CREATE TABLE public.plan_messages (
        id                  uuid PRIMARY KEY,
        plan_id             uuid references public.plan(id),
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
