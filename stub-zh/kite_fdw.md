## 用法

来源：

- [Official README](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/README.md)
- [Extension control file](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/kite_fdw.control)
- [Installation SQL](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/kite_fdw--1.0.sql)
- [Installation SQL](https://github.com/vderic/postgres-kite/blob/ef6334886b99ea83cf8df253da5992732599f551/kite_fdw--1.0--1.1.sql)

`kite_fdw` 通过 PostgreSQL 外部表查询兼容的远端 Kite 服务。已核验的源码较旧，不能据此确认与当前 PostgreSQL 版本兼容。

### 基本用法

配置 Kite 端点、远端数据库及逐用户凭据。外部表选项描述远端表或文件模式及其数据格式。

```sql
CREATE EXTENSION kite_fdw;
CREATE SERVER kite_server FOREIGN DATA WRAPPER kite_fdw
  OPTIONS (host '127.0.0.1:7878', dbname 'pgsql', fragcnt '4');
CREATE USER MAPPING FOR CURRENT_USER SERVER kite_server
  OPTIONS (username 'app_user', password 'replace-with-password');
CREATE FOREIGN TABLE warehouse (warehouse_id int, warehouse_name text)
  SERVER kite_server
  OPTIONS (schema_name 'public', table_name 'warehouse*', fmt 'csv', csv_header 'false');
SELECT * FROM warehouse;
```

### 选项与要求

`host` 接受逗号分隔的端点，`dbname` 为必填项，`fragcnt` 控制查询分片。`fetch_size` 默认为 100，可设置在服务器或表级。`fmt` 支持 CSV 或 Parquet，并提供 CSV 分隔符、引号、转义、表头及空值字符串选项。control 版本为 `1.1`，由基础 SQL 和升级脚本构成。未找到独立的许可证声明或受支持大版本矩阵；使用前应审查源码及外部 Kite 依赖。
