## 用法

来源：

- [固定提交的上游 README](https://github.com/durch/google-bigtable-postgres-fdw/blob/8337808ab544a202755f38f506cde8def8bfcabb/README.md)
- [固定提交的扩展 control 文件](https://github.com/durch/google-bigtable-postgres-fdw/blob/8337808ab544a202755f38f506cde8def8bfcabb/bigtable.control)

`bigtable` 是面向 Google Cloud Bigtable 的 Rust 外部数据包装器。已复核实现把远端行表示为一个 PostgreSQL `json` 值，并通过外部服务器、外部表和用户映射选项配置云实例、项目、表及服务账号凭据路径。

```sql
CREATE EXTENSION bigtable;

CREATE SERVER bigtable_srv
  FOREIGN DATA WRAPPER bigtable
  OPTIONS (instance 'instance_id', project 'project_id');

CREATE FOREIGN TABLE bt_rows (bt json)
  SERVER bigtable_srv
  OPTIONS (name 'table_name');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER bigtable_srv
  OPTIONS (credentials_path '/secure/service-account.json');

SELECT bt->>'rowKey' FROM bt_rows LIMIT 10;
```

固定版本的 README 记录了 `SELECT`、`SELECT LIMIT` 和 `INSERT`。其路线图仍未完成通用 `WHERE`、`OFFSET`、`DELETE` 与原生 `UPDATE`；文档把插入已有键描述为更新方式。应把谓词行为视为实现细节，并验证过滤是在远端下推还是在本地执行。

已复核 control 版本为 `0.1.0`，上游环境仍面向较旧的 PostgreSQL 9.6/Rust 1.16 时代。凭据文件应放在数据库不可随意读取的位置，收紧其文件权限与用户映射，并在使用前验证 Bigtable schema/column-family 假设，以及当前 API、TLS、重试、配额和事务语义。
