## 用法

来源：

- [Official README](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/README.md)
- [Extension control file](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/pg_xdbc_fdw.control)
- [Installation SQL](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/pg_xdbc_fdw--0.1.sql)
- [Official SQL example](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/test/test_fdw_create.sql)
- [Official README](https://github.com/polydbms/pg_xdbc_fdw/blob/fac1fc6b5275d327843a3efb4d519f9afe6a161b/docker/README.md)

`pg_xdbc_fdw` 是连接 PostgreSQL 与 XDBC 客户端/服务端数据传输系统的研究性连接器。上游提供 PostgreSQL 13 容器拓扑与实验配置，并非通用 ODBC 兼容层。

### 基本用法

查询前需安装 XDBC 客户端依赖并运行对应服务端。本地 JSON 模式必须与声明列及远端数据一致。

```sql
CREATE EXTENSION pg_xdbc_fdw;
CREATE SERVER xdbcserver FOREIGN DATA WRAPPER pg_xdbc_fdw;
CREATE FOREIGN TABLE transfer_data (id integer, value text)
  SERVER xdbcserver
  OPTIONS (schema_file_path '/srv/xdbc/schema.json', server_host 'xdbcserver', table 'transfer_data');
SELECT * FROM transfer_data;
```

### 使用边界

扩展安装 `pg_xdbc_fdw_handler()` 及其外部数据包装器。`schema_file_path` 从 PostgreSQL 主机读取，`server_host` 和 `table` 选择 XDBC 数据源。匹配的模式文件和可访问的服务必须预先存在，并非由这些 SQL 创建。上游测试还使用缓冲区设置，应按实际 XDBC 构建配置。项目文档未确立通用的模式导入、写入、安全或跨系统事务保证。
