## 用法

来源：

- [Official README](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/README.md)
- [Extension control file](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_core_fdw/pgspider_core_fdw.control)
- [Installation SQL](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_core_fdw/pgspider_core_fdw--1.0.sql)
- [Build configuration](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_core_fdw/Makefile)

`pgspider_core_fdw` 在 PGSpider 数据库分支中协调子外部表查询。它不同于可独立安装的 `pgspider_ext`，需要 PGSpider 内核改动及其 `pgspider_keepalive` 支持库。

### 基本用法

在匹配的 PGSpider 服务器上安装包装器及子 FDW。以下示例假设远端数据库已存在列结构匹配的 t1 表，连接参数需替换为实际值。

```sql
CREATE EXTENSION pgspider_core_fdw;
CREATE EXTENSION postgres_fdw;
CREATE SERVER parent FOREIGN DATA WRAPPER pgspider_core_fdw;
CREATE SERVER postgres_svr FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '127.0.0.1', port '5432', dbname 'postgres');
CREATE USER MAPPING FOR CURRENT_USER SERVER postgres_svr
  OPTIONS (user 'app_user', password 'replace-me');
CREATE FOREIGN TABLE t1(i integer, t text, __spd_url text) SERVER parent;
CREATE FOREIGN TABLE t1__postgres_svr__0(i integer, t text)
  SERVER postgres_svr OPTIONS (table_name 't1');
SELECT * FROM t1;
```

### 映射与对象

父表包含标识数据来源的 `__spd_url` 列。子表名称遵循文档中父表名、服务器名及后缀的约定；可通过另一个子 FDW 将第二种数据源映射为相同的逻辑结构。模块安装处理器、校验器、外部数据包装器，以及返回整数版本标识的 `pgspider_core_fdw_version()`。

### 运行边界

应为每个数据源配置 FDW、服务器权限与用户映射。INSERT 路由到适当子节点，UPDATE/DELETE 行为取决于所有参与的 FDW。上游不支持 `RETURNING`、`WITH CHECK OPTION`、`ON CONFLICT` 及通过外部分区执行修改，COPY 也有限制。须规划远端故障与事务处理，不能据此假设跨系统原子提交。该模块未声明预加载要求；SQL/control 版本 `1.0` 表示模块版本，不代表原生 PostgreSQL 版本。
