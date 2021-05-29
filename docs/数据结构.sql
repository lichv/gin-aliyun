
CREATE SEQUENCE if not exists "public"."aliyun_config_id_seq" 
INCREMENT 1
MINVALUE  10000
MAXVALUE 9999999
START 10000
CACHE 10000;

ALTER SEQUENCE "public"."aliyun_config_id_seq" OWNER TO "postgres";

DROP TABLE IF EXISTS "public"."aliyun_config";
CREATE TABLE "public"."aliyun_config" (
  "id" int4 NOT NULL DEFAULT nextval('aliyun_config_id_seq'::regclass),
  "code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "group" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "company" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "access_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "access_key" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "region" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "rolesessionrole" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "rolearn" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "tags" json NOT NULL DEFAULT '{}'::json,
  "created_at" timestamp ,
  "updated_at" timestamp ,
  "deleted_at" timestamp ,
  "flag" bool NOT NULL DEFAULT false,
  "state" bool NOT NULL DEFAULT false
)
;

ALTER TABLE "public"."aliyun_config" 
  OWNER TO "postgres";

COMMENT ON COLUMN "public"."aliyun_config"."code" IS '系统自增编号';

COMMENT ON COLUMN "public"."aliyun_config"."code" IS '阿里云平台编号';

COMMENT ON COLUMN "public"."aliyun_config"."group" IS '分组';

COMMENT ON COLUMN "public"."aliyun_config"."company" IS '公司';

COMMENT ON COLUMN "public"."aliyun_config"."access_id" IS 'access_id';

COMMENT ON COLUMN "public"."aliyun_config"."access_key" IS 'access_key';

COMMENT ON COLUMN "public"."aliyun_config"."region" IS '区域';

COMMENT ON COLUMN "public"."aliyun_config"."rolesessionrole" IS '认证回调地址';

COMMENT ON COLUMN "public"."aliyun_config"."rolearn" IS '通知地址';

COMMENT ON COLUMN "public"."aliyun_config"."tags" IS '标签';

COMMENT ON COLUMN "public"."aliyun_config"."created_at" IS '创建时间';

COMMENT ON COLUMN "public"."aliyun_config"."updated_at" IS '修改时间';

COMMENT ON COLUMN "public"."aliyun_config"."deleted_at" IS '删除时间';

COMMENT ON COLUMN "public"."aliyun_config"."flag" IS '标记';

COMMENT ON COLUMN "public"."aliyun_config"."state" IS '状态';

COMMENT ON TABLE "public"."aliyun_config" IS '阿里云配置';




ALTER TABLE "public"."aliyun_config" DROP CONSTRAINT IF EXISTS "aliyun_config_pkey";
ALTER TABLE "public"."aliyun_config" ADD CONSTRAINT "aliyun_config_pkey" PRIMARY KEY ("id");


DROP INDEX IF EXISTS "aliyun_config_code";
CREATE INDEX "aliyun_config_code" ON "public"."aliyun_config" ("code");


INSERT INTO "public"."aliyun_config"("code", "group", "company", "access_id", "access_key", "region", "rolesessionrole", "rolearn", "tags", "flag", "state") VALUES ('dg54df3g', 'easytc', 'easytc', 'LTAI5tKnra54xozuA3KktFur', 'VyIHrtVQZxXeiuuBWUW2oG34qe87dk', 'cn-hangzhou', 'osser', 'acs:ram::1378573870105843:role/osser', '{}', 't', 't');
