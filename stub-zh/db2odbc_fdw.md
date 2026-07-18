## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/stanislawbartkowski/db2odbc_fdw/blob/fefa0d2181e4089fa39f0743c473d49dcf7a5e3f/README.md)
- [1.0 版本 SQL 对象](https://github.com/stanislawbartkowski/db2odbc_fdw/blob/fefa0d2181e4089fa39f0743c473d49dcf7a5e3f/db2odbc_fdw--1.0.sql)
- [FDW 实现](https://github.com/stanislawbartkowski/db2odbc_fdw/blob/fefa0d2181e4089fa39f0743c473d49dcf7a5e3f/db2odbc_fdw.c)

`db2odbc_fdw` 是 PostgreSQL 12 时代通过操作系统 ODBC 数据源读取 DB2 的包装器。外部表映射到用户提供的远端 `sql_query`；服务器选项包括 ODBC `dsn` 和可选的连接缓存设置。

```sql
CREATE EXTENSION db2odbc_fdw;
CREATE SERVER db2_server
  FOREIGN DATA WRAPPER db2odbc_fdw OPTIONS (dsn 'REPORTING');
CREATE USER MAPPING FOR report_role SERVER db2_server
  OPTIONS (username 'reporter', password 'secret');
CREATE FOREIGN TABLE db2_customer (id integer, name text)
  SERVER db2_server OPTIONS (sql_query 'SELECT ID, NAME FROM CUSTOMER');
```

不要在 SQL 中保留示例密码，应使用真正的秘密流程，限制用户映射与 ODBC 配置访问，也绝不能把 FDW 或服务器授予 `PUBLIC`。`sql_query` 定义远端 SQL，必须视为可信管理员输入。

上游建议高级用户改用功能完整的 DB2 FDW，且只记录 PostgreSQL 12。应验证 ODBC 驱动 ABI、编码、数值与时间转换、空值、远端错误、连接重试/缓存、取消、超时、凭据轮换、事务边界和谓词行为。除非精确源码与测试证明，否则不要假定它支持写入或分布式事务语义。
