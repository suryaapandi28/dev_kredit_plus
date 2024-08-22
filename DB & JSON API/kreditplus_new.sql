/*
 Navicat Premium Data Transfer

 Source Server         : kreditplus
 Source Server Type    : PostgreSQL
 Source Server Version : 160003 (160003)
 Source Host           : localhost:5432
 Source Catalog        : kreditplus
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160003 (160003)
 File Encoding         : 65001

 Date: 22/08/2024 16:06:35
*/


-- ----------------------------
-- Table structure for Transaksi_headers
-- ----------------------------
DROP TABLE IF EXISTS "public"."Transaksi_headers";
CREATE TABLE "public"."Transaksi_headers" (
  "Kode_transaksi" uuid NOT NULL,
  "Kode_konsumen" uuid,
  "Nomor_kontrak" uuid,
  "OTR" varchar(255) COLLATE "pg_catalog"."default",
  "Admin_fee" int8,
  "Jumlah_cicilan" int8,
  "Jumlah_bunga" int8,
  "Nama_aset" varchar(255) COLLATE "pg_catalog"."default",
  "Create_at" timestamptz(6),
  "Update_at" timestamptz(6),
  "Deleted_at" timestamptz(6)
)
;

-- ----------------------------
-- Records of Transaksi_headers
-- ----------------------------

-- ----------------------------
-- Table structure for accounts
-- ----------------------------
DROP TABLE IF EXISTS "public"."accounts";
CREATE TABLE "public"."accounts" (
  "kode_account" uuid NOT NULL,
  "kode_konsumen" uuid,
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "status_account" varchar(100) COLLATE "pg_catalog"."default",
  "verification_status" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "verification_code" varchar(6) COLLATE "pg_catalog"."default" NOT NULL,
  "verification_expired" timestamptz(6) NOT NULL,
  "jwt_token" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "jwt_token_expired_at" timestamptz(6) NOT NULL,
  "device_id" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;

-- ----------------------------
-- Records of accounts
-- ----------------------------
INSERT INTO "public"."accounts" VALUES ('94aaf628-2306-4ad7-89c3-9aa6a5ec51ff', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$.kU3Vi9DBqaZt6/QIdQHhuuEt4VZfgF5MDqhGHb7LzP3hspQjTWNu', 'No_Active', 'No_Verify', '106402', '2024-08-21 23:13:30.674319+00', '', '0001-01-01 00:00:00+00', '', '2024-08-20 16:13:30.674319+00', '2024-08-20 16:13:30.674319+00', NULL);
INSERT INTO "public"."accounts" VALUES ('6a4891a5-6786-49a3-9145-fe284ca56dff', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$Hfohf6Kv86ehS/Gu8jPPd.yWlJYwACUkbQeDsRXyE5VUgAyq3o.GS', 'No_Active', 'No_Verify', '938186', '2024-08-21 23:27:57.259525+00', '', '0001-01-01 00:00:00+00', '', '2024-08-20 16:27:57.278881+00', '2024-08-20 16:27:57.278881+00', NULL);
INSERT INTO "public"."accounts" VALUES ('dcb06973-45ba-4e5b-85fc-6dddc067d00c', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$6CdzqfderAvBYX9WxPwnPucZEmkELcGqBRc9uLdfiI.NKKqjYX4me', 'No_Active', 'No_Verify', '590052', '2024-08-21 23:31:25.046854+00', '', '0001-01-01 00:00:00+00', '', '2024-08-20 16:31:25.080122+00', '2024-08-20 16:31:25.080122+00', NULL);
INSERT INTO "public"."accounts" VALUES ('5ca3d64c-72d4-4f41-bd9c-dd86ca269ffd', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$dd96MM7kTS9HAGJdVRZH5OENqkwh5HewlD1mhVjvkZsT.W6HPduLu', 'No_Active', 'No_Verify', '473110', '2024-08-21 23:36:48.874023+00', '', '0001-01-01 00:00:00+00', '', '2024-08-20 16:36:48.87508+00', '2024-08-20 16:36:48.87508+00', NULL);
INSERT INTO "public"."accounts" VALUES ('56f3000e-9f1c-4162-ae12-dad221b9ef15', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$MTVT14hqzVdJdOqQR/DTVecCcFDbmyXjPhNN9MA8HJs7/1Y5zP.Lu', 'No_Active', 'No_Verify', '577449', '2024-08-21 23:39:06.553738+00', '', '0001-01-01 00:00:00+00', '', '2024-08-20 16:39:06.554724+00', '2024-08-20 16:39:06.554724+00', NULL);
INSERT INTO "public"."accounts" VALUES ('be3088e0-ec0a-4c69-9615-8bdc48fe52d5', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$Pd3.YDxcFdQ7J4NX4LPVuOR6mk6ZH4sQLAKZcGEaEMjBrKPBfkgbm', 'No_Active', 'No_Verify', '077410', '2024-08-21 23:53:02.658618+00', '', '0001-01-01 00:00:00+00', '', '2024-08-20 16:53:02.686504+00', '2024-08-20 16:53:02.686504+00', NULL);
INSERT INTO "public"."accounts" VALUES ('ce36f80a-e70b-45a1-b3ef-d5f9b5b5d644', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$t2Eytb.zfiEKxEo3XeBGy.j5Z2WSVja7x3qNdi9Vm0mY.iMRCahWW', 'No_Active', 'No_Verify', '123507', '2024-08-21 13:29:21.558787+00', '', '0001-01-01 00:00:00+00', '', '2024-08-21 06:24:21.582557+00', '2024-08-21 06:24:21.582557+00', NULL);
INSERT INTO "public"."accounts" VALUES ('8154af37-15d6-4031-a470-e6fd9b3438d4', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$rF7BjhR7dhIcEmqs3E1ZE.i8Y47D9MWz7LvC8xmFKFiQZkost.eju', 'No_Active', 'No_Verify', '034040', '2024-08-21 13:33:05.029235+00', '', '0001-01-01 00:00:00+00', 'sdskdsjdsd', '2024-08-21 06:28:05.061986+00', '2024-08-21 06:28:05.061986+00', NULL);
INSERT INTO "public"."accounts" VALUES ('b2dcf668-45c6-48f7-894d-a51390bb04da', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$6dVlD9eeIm8xXReVndcJ4OVPgqxTkcD8vASTtxYxHOSqK8izueIpy', 'Active', 'Verify', ' ', '0001-01-01 00:00:00+00', '', '0001-01-01 00:00:00+00', 'sdskdsjdsd', '2024-08-21 17:40:00.207803+00', '2024-08-21 17:40:47.022372+00', NULL);
INSERT INTO "public"."accounts" VALUES ('05ac6c17-ec2f-477b-858c-5551a9e2b8fa', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$nezLCo4x1SHw/r5LezwKGerYkHkvcyhSQ4UgHsHM1TXkjuD.tu.iG', 'Active', 'Verify', ' ', '0001-01-01 00:00:00+00', '', '0001-01-01 00:00:00+00', 'sdskdsjdsd', '2024-08-21 17:07:15.216043+00', '2024-08-21 19:36:11.300999+00', NULL);
INSERT INTO "public"."accounts" VALUES ('9697182f-2c10-4f30-b2c2-de548e5390f6', 'e8669f42-df77-43de-a6fe-903f3852fdb6', 'surya.apandi28@gmail.com', '$2a$10$tZG721XANKjs4awMpEyw4um6VkzgUMs9cWw833GJdu.HnuaV3o0gK', 'Active', 'Verify', ' ', '0001-01-01 00:00:00+00', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijk2OTcxODJmLTJjMTAtNGYzMC1iMmMyLWRlNTQ4ZTUzOTBmNiIsImVtYWlsIjoic3VyeWEuYXBhbmRpMjhAZ21haWwuY29tIiwicm9sZSI6InVzZXIiLCJpc3MiOiJrcmVkaXRwbHVzIiwiZXhwIjoxNzI0MzU3NjQ0fQ.mMoO4QLPdPklmI75Y8syqipxmKpdOZ5U-AWgma7BYXw', '2024-08-22 20:14:04.119177+00', 'sdskdsjdsd', '2024-08-21 19:39:13.475893+00', '2024-08-21 20:14:04.170632+00', NULL);

-- ----------------------------
-- Table structure for konsumens
-- ----------------------------
DROP TABLE IF EXISTS "public"."konsumens";
CREATE TABLE "public"."konsumens" (
  "kode_konsumen" uuid NOT NULL,
  "nik" text COLLATE "pg_catalog"."default",
  "full_name" varchar(255) COLLATE "pg_catalog"."default",
  "legal_name" varchar(100) COLLATE "pg_catalog"."default",
  "tempat_lahir" varchar(255) COLLATE "pg_catalog"."default",
  "tanggal_lahir" date,
  "gaji" int8,
  "foto_ktp" text COLLATE "pg_catalog"."default",
  "foto_selfie" text COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "kredit_score" varchar(1) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of konsumens
-- ----------------------------
INSERT INTO "public"."konsumens" VALUES ('9cccc10e-e058-4e45-aedf-e5d1c3c7cbf3', 'ODk0ODM5ODM5NjM5ODI=', 'Muhammad Surya Apandi', 'Surya', 'Palembang', '2024-01-19', 5000000, 'MDJiZWY1MDUtMDdlZC00YzA3LTkwODEtMWYzZTQ5YTM2MzA2LnBuZw==', 'MDJiZWY1MDUtMDdlZC00YzA3LTkwODEtMWYzZTQ5YTM2MzA2LmpwZWc=', '2024-08-22 05:26:05.248744+00', '2024-08-22 05:26:05.248744+00', NULL, 'A');
INSERT INTO "public"."konsumens" VALUES ('e8669f42-df77-43de-a6fe-903f3852fdb6', 'NzM0ODczMzkzODQ=', 'dadada', 'dasdsada', 'dsasdasdsada', '2024-01-19', 22222, 'MzVlZTU5OGMtZmQzYS00NTg4LWE2N2YtNjQ5OGYwYzNkMTQ1LnBuZw==', 'MzVlZTU5OGMtZmQzYS00NTg4LWE2N2YtNjQ5OGYwYzNkMTQ1LmpwZWc=', '2024-08-20 09:37:08.360356+00', '2024-08-20 09:37:08.360356+00', NULL, 'A');
INSERT INTO "public"."konsumens" VALUES ('f64818e3-ef80-48e1-bc32-9f9f4baf34fe', 'ODk0ODM5ODM5OA==', 'Muhammad Surya Apandi', 'Surya', 'Palembang', '2024-01-19', 5000000, 'NTI5MzQzMWUtZmM4ZS00NDFmLWFiNjItM2EzNTE1NGI4NWE0LnBuZw==', 'NTI5MzQzMWUtZmM4ZS00NDFmLWFiNjItM2EzNTE1NGI4NWE0LmpwZWc=', '2024-08-20 09:58:15.231997+00', '2024-08-20 09:58:15.231997+00', NULL, 'A');
INSERT INTO "public"."konsumens" VALUES ('c234b027-f1ff-43de-96cb-4892f6d1c3bb', 'ODk0ODM5ODM5ODEyMw==', 'Muhammad Surya Apandi', 'Surya', 'Palembang', '2024-01-19', 5000000, 'ZjExOGQyMGUtMDk4NC00M2I2LThmNGQtNzJkMDkyYjlhNDI1LnBuZw==', 'ZjExOGQyMGUtMDk4NC00M2I2LThmNGQtNzJkMDkyYjlhNDI1LmpwZWc=', '2024-08-20 10:09:19.608707+00', '2024-08-20 10:09:19.608707+00', NULL, 'A');
INSERT INTO "public"."konsumens" VALUES ('483fd762-bd0f-4983-8e05-d08af280df53', 'ODk0ODM5ODM5ODEyMzE=', 'Muhammad Surya Apandi', 'Surya', 'Palembang', '2024-01-19', 5000000, 'ZjQ0MzlhOTQtZjI5Ni00ZGRmLWE1ZmMtZTk3OTkyYzg5ODgzLnBuZw==', 'ZjQ0MzlhOTQtZjI5Ni00ZGRmLWE1ZmMtZTk3OTkyYzg5ODgzLmpwZWc=', '2024-08-21 20:23:33.002535+00', '2024-08-21 20:23:33.002535+00', NULL, 'A');
INSERT INTO "public"."konsumens" VALUES ('03a060e7-b97d-40ad-b79a-745ff01720c7', 'ODk0ODM5ODM5', 'Muhammad Surya Apandi', 'Surya', 'Palembang', '2024-01-19', 5000000, 'Y2M0MjVkOTAtNjIxMi00MGI3LWE0ZTUtYjcyNzk2MDYwYzg2LnBuZw==', 'Y2M0MjVkOTAtNjIxMi00MGI3LWE0ZTUtYjcyNzk2MDYwYzg2LmpwZWc=', '2024-08-21 20:40:33.792092+00', '2024-08-21 20:40:33.792092+00', NULL, 'A');
INSERT INTO "public"."konsumens" VALUES ('0524cc0b-b699-4cfe-9251-a9af7c1f45d7', 'ODk0ODM5ODM5NjM5ODIyNDU=', 'Muhammad Surya Apandi', 'Surya', 'Palembang', '2024-01-19', 5000000, 'OWJmZDAzNzAtZTJlYi00MGRjLTg2NDQtMDJhNDU0MGMzYzIwLnBuZw==', 'OWJmZDAzNzAtZTJlYi00MGRjLTg2NDQtMDJhNDU0MGMzYzIwLmpwZWc=', '2024-08-22 06:46:29.616681+00', '2024-08-22 06:46:29.616681+00', NULL, 'A');

-- ----------------------------
-- Table structure for tenor_limits
-- ----------------------------
DROP TABLE IF EXISTS "public"."tenor_limits";
CREATE TABLE "public"."tenor_limits" (
  "kode_tenor" uuid NOT NULL,
  "kode_konsumen" uuid,
  "nama_konsumen" varchar(255) COLLATE "pg_catalog"."default",
  "tenor_pertama" int8,
  "tenor_kedua" int8,
  "tenor_ketiga" int8,
  "tenor_keenam" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;

-- ----------------------------
-- Records of tenor_limits
-- ----------------------------
INSERT INTO "public"."tenor_limits" VALUES ('33e62161-f786-4ad4-9522-495810866350', '0524cc0b-b699-4cfe-9251-a9af7c1f45d7', 'Muhammad Surya Apandi', 2500000, 3000000, 3750000, 5000000, '2024-08-22 08:06:03.458537+00', '2024-08-22 08:06:03.458537+00', NULL);

-- ----------------------------
-- Primary Key structure for table Transaksi_headers
-- ----------------------------
ALTER TABLE "public"."Transaksi_headers" ADD CONSTRAINT "Transaksi_header_pkey" PRIMARY KEY ("Kode_transaksi");

-- ----------------------------
-- Primary Key structure for table accounts
-- ----------------------------
ALTER TABLE "public"."accounts" ADD CONSTRAINT "accounts_pkey" PRIMARY KEY ("kode_account");

-- ----------------------------
-- Primary Key structure for table konsumens
-- ----------------------------
ALTER TABLE "public"."konsumens" ADD CONSTRAINT "Konsumen_pkey" PRIMARY KEY ("kode_konsumen");

-- ----------------------------
-- Primary Key structure for table tenor_limits
-- ----------------------------
ALTER TABLE "public"."tenor_limits" ADD CONSTRAINT "Tenor_Limit_pkey" PRIMARY KEY ("kode_tenor");

-- ----------------------------
-- Foreign Keys structure for table Transaksi_headers
-- ----------------------------
ALTER TABLE "public"."Transaksi_headers" ADD CONSTRAINT "Transaksi_headers_Kode_konsumen_fkey" FOREIGN KEY ("Kode_konsumen") REFERENCES "public"."konsumens" ("kode_konsumen") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table accounts
-- ----------------------------
ALTER TABLE "public"."accounts" ADD CONSTRAINT "accounts_kode_konsumen_fkey" FOREIGN KEY ("kode_konsumen") REFERENCES "public"."konsumens" ("kode_konsumen") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table tenor_limits
-- ----------------------------
ALTER TABLE "public"."tenor_limits" ADD CONSTRAINT "tenor_limits_kode_konsumen_fkey" FOREIGN KEY ("kode_konsumen") REFERENCES "public"."konsumens" ("kode_konsumen") ON DELETE NO ACTION ON UPDATE NO ACTION;
