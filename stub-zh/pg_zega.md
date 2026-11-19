## 用法

来源：

- [官方资料](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/README.md)
- [官方资料](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/postgres/pg_zega.control)
- [官方资料](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/LICENSE)
- [官方资料](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/postgres/pg_zega--0.3.sql)

pg_zega 0.3 将支持的表及 DDL 同步到外部 Zega 图数据库，要求 PostgreSQL 17、libcurl 及超级用户管理。PostgreSQL 凭据与 Zega bearer token 是不同的秘密。跟踪数据前应配置 zega.url 和受保护的 zega.password，重载并重新连接。

### 跟踪与导入

```sql
CREATE EXTENSION pg_zega;
CREATE TABLE users(user_id integer PRIMARY KEY, name text NOT NULL);
SELECT zega.track('users');
INSERT INTO users VALUES (10, 'Nia');
SELECT zega.import('users');
```

跟踪会添加行触发器，但不复制已有行。节点表要求单列主键；字段支持整数、浮点、布尔及文本类类型。ze_ 前缀的边表要求恰好两个列不相交的外键约束，并以这些列组成主键。不支持的类型包括 numeric、timestamp、uuid 和 json，会被拒绝。

### 投递与维护

默认 precommit 模式等待 Zega，拒绝会中止 PostgreSQL 事务；这不构成通用的分布式事务保证。Journal 模式先本地提交，再借助 Zega 收据异步投递，要求预加载并重启、逻辑 WAL、投递工作进程，以及在启用模式前创建复制槽。需要监测 journal 延迟与保留 WAL。

zega.import 会阻止所有被跟踪表的写入，并可能替换其范围内的图节点与边。应在事务中不包含其他行变更的情况下运行；导入后的引擎 ID 会变化。DDL 事件触发器会拒绝与现有图值不兼容的修改。Precommit 模式无需预加载，journal 模式必须采用官方指南中的额外配置。
