-- +migrate Up
-- ベースユーザー
create table if not exists base_users (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "email" varchar(255) not null,
  "name" varchar(255) not null
);

-- 管理企業ユーザー
create table if not exists admin_users (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "base_user_id" bigint references base_users (id) on delete cascade on update cascade
);

-- 会社
create table if not exists companies (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "name" varchar(255) not null,
  "tel" varchar(255)
);

-- 事業所
create table if not exists offices (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "contract_status" int not null,
  "name" varchar(255) not null,
  "tel" varchar(255)
);

-- 契約企業ユーザー
create table if not exists users (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "base_user_id" bigint references base_users (id) on delete cascade on update cascade,
  "role" varchar(255) not null,
  "office_id" bigint references offices (id) on delete cascade on update cascade
);

-- 取引会社
create table if not exists partner_companies (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "name" varchar(255) not null,
  "category" varchar(255) not null,
  "industry" varchar(255) not null
);

-- 取引会社の人の連絡先
create table if not exists personnel (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "partner_company_id" bigint references partner_companies (id) on delete cascade on update cascade,
  "name" varchar(255) not null,
  "tel" varchar(255),
  "email" varchar(255) not null
);

-- 期間
create table if not exists terms (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "started_on" date not null,
  "ended_on" date not null
);

-- 特別休日期間
create table if not exists except_terms (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "started_on" date not null,
  "ended_on" date not null
);

-- 画像リスト
create table if not exists image_lists (
  "id" bigserial primary key,
  "created_at" timestamp not null
);

-- 画像
create table if not exists images (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "name" varchar(2047) not null,
  "file_path" varchar(2047) not null
);

-- 現場メール配信
create table if not exists construction_site_subscriptions (
  "personnel_item_id" bigint references personnel (id) on delete cascade on update cascade,
  "user_id" bigint references users (id) on delete cascade on update cascade
);

-- 現場/案件
create table if not exists construction_sites (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "name" varchar(255) not null,
  "status" int not null,
  "construction_type" int not null,
  "address" varchar(1023) not null,
  "contracted_at" timestamp not null,
  "started_on" date not null,
  "ended_on" date not null,
  "holiday_type" timestamp not null
);

-- ユーザーと現場の中間
create table if not exists user_construction_sites (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "user_id" bigint references users (id) on delete cascade on update cascade,
  "construction_site_id" bigint references construction_sites (id) on delete cascade on update cascade
);

-- 現場ごとの工程表更新の通知先 取引会社の人の連絡先と現場の中間
create table if not exists construction_site_subscriptions (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "construction_site_id" bigint references construction_sites (id) on delete cascade on update cascade,
  "personnel_item_id" bigint references personnel (id) on delete cascade on update cascade
);

-- 工事
create table if not exists constructions (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "name" varchar(255) not null,
  "user_id" bigint references users (id) on delete cascade on update cascade
);

-- 小工事
create table if not exists sub_constructions (
  "id" bigserial primary key,
  "created_at" timestamp not null,
  "updated_at" timestamp not null,
  "name" varchar(255) not null,
  "personnel_item_id" bigint references personnel (id) on delete cascade on update cascade,
  "image_list_id" bigint references image_lists (id) on delete cascade on update cascade,
  "term_id" bigint references terms (id) on delete cascade on update cascade
);

-- +migrate Down
drop table if exists companies;
drop table if exists user_construction_sites;
drop table if exists construction_site_subscriptions;
drop table if exists construction_site_subscriptions;
drop table if exists construction_site_subscriptions;
drop table if exists construction_sites;
drop table if exists constructions;
drop table if exists sub_constructions;
drop table if exists terms;
drop table if exists except_terms;
drop table if exists image_lists;
drop table if exists images;
drop table if exists image_information;
drop table if exists admin_users;
drop table if exists personnel;
drop table if exists partner_companies;
drop table if exists users;
drop table if exists base_users;
drop table if exists offices;
