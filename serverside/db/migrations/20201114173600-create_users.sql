-- +migrate Up
-- 企業
-- ベース法人情報
create table if not exists corporations (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "name" varchar(255) not null,
  "tel" varchar(16)
);

-- 顧客企業
create table if not exists client_companies (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "corporation_id" bigint references corporations (id),
  "name" varchar(255) not null,
  "status" integer not null
);

-- 契約会社
create table if not exists contract_companies (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "corporation_id" bigint references corporations (id),
  "name" varchar(255) not null,
  "status" integer not null
);

-- 協力会社
create table if not exists coop_companies (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "corporation_id" bigint references corporations (id)
);

-- 事業所
-- ベース事業所情報
create table if not exists offices (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "corporation_id" bigint references corporations (id)
);

-- 顧客事業所
create table if not exists client_offices (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "office_id" bigint references offices (id)
);

-- 契約企業事業所
create table if not exists contract_offices (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "office_id" bigint references offices (id)
);

-- 協力会社事業所
create table if not exists coop_offices (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "office_id" bigint references offices (id)
);

-- ユーザー
-- ベースユーザー
create table if not exists users (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "office_id" bigint references offices (id)
);

-- 契約企業ユーザー
create table if not exists contract_users (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "user_id" bigint references users (id)
);

-- 担当者
-- ベース担当者情報
create table if not exists in_charge_peaple (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "office_id" bigint references offices (id)
);

-- 顧客担当者
create table if not exists client_in_charge_peaple (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "in_charge_person_id" bigint references in_charge_peaple (id)
);

-- 協力会社担当者
create table if not exists coop_in_charge_peaple (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "in_charge_person_id" bigint references in_charge_peaple (id)
);

-- +migrate Down
drop table if exists contract_users;

drop table if exists users;

drop table if exists client_in_charge_peaple;

drop table if exists coop_in_charge_peaple;

drop table if exists in_charge_peaple;

drop table if exists contract_offices;

drop table if exists coop_offices;

drop table if exists client_offices;

drop table if exists offices;

drop table if exists client_companies;

drop table if exists contract_companies;

drop table if exists coop_companies;

drop table if exists corporations;
