## 用法

来源：

- [官方扩展控制文件](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/check_orapg.control)
- [官方扩展 SQL](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/check_orapg--3.0.sql)
- [官方 README](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/README.md)

`check_orapg` 3.0 用于辅助验证 Oracle 到 PostgreSQL 或 Oracle 到 EDB 的迁移。它通过 `oracle_fdw` 复制选定的 Oracle 目录信息，统计两侧的对象、权限、属性、注释与表行数，并输出两份文本报告供外部比较。它报告迁移覆盖情况，但不会转换模式，也不能证明应用行为完全等价。

### 核心流程

控制文件要求 `oracle_fdw`，并把安装限定为超级用户；上游只记录了 PostgreSQL 11–13 和 EDB Postgres 11–13 的测试。创建 `check_orapg` 前，需要安装 Oracle 客户端库并配置好 `oracle_fdw`。

```sql
CREATE EXTENSION oracle_fdw;
CREATE EXTENSION check_orapg;

SELECT check_orapg.create_oracle_server(
    'ora_migration', '192.0.2.10', 1521, 'ORCL'
);
SELECT check_orapg.create_user(
    'ora_migration', 'migration_reader', 'oracle_password'
);

SELECT check_orapg.create_oracle_tables(
    'ora_migration', NULL, '''HR'',''SALES'''
);

SELECT *
FROM check_orapg.schemas_objects_postgres('''HR'',''SALES''') AS (
    schema_name text, tables int, ordinary_views int,
    materialized_views int, triggers int, sequences int,
    indexes int, functions int, procedures int, constraints int,
    table_comments int, column_comments int
);
```

导入 Oracle 目录后，可用 `check_orapg.update_all_oracle_tables` 刷新数据，用 `check_orapg.register_postgres_validation` 和 `check_orapg.register_oracle_validation` 登记两侧快照，再比较 `check_orapg.postgres_file` 与 `check_orapg.oracle_file` 返回的行。`generate_postgres_file` 和 `generate_oracle_file` 变体会把报告写入数据库服务器文件系统。

### 重要对象

- 服务器与映射辅助函数：`create_oracle_server`、`update_server`、`show_servers`、`create_user`、`update_user_password` 和 `show_users`。
- Oracle 目录刷新：`create_oracle_tables`、`update_all_oracle_tables`、`update_oracle_table` 和 `update_oracle_tables_rows`。
- 比较函数：`cluster_postgres`、`cluster_oracle`、权限报告、模式对象报告与行数报告。
- 持久化结果：31 张导入的 Oracle 目录表，以及模式 `check_orapg` 中的 `postgres_validation` 和 `oracle_validation`。

### 安全与准确性边界

`create_user` 创建的是面向 `PUBLIC` 的用户映射，而不是只面向调用者；密码作为函数文本传入并保存在外部服务器元数据中。该 SQL 还会用标识符和凭据拼接动态 DDL，缺少稳健引用。所有管理函数都应只授权给严格受控的管理员角色，并使用只读 Oracle 账号。

文件生成函数使用服务器端 `COPY`，只能写入 PostgreSQL 操作系统账号有权限的位置；输出路径是服务器路径。由于 PostgreSQL 不实现 Oracle 包或同义词，而且对象定义、数据值、约束与行为并未逐项比较，统计差异可能是合理的。生成报告只能作为迁移证据，不能当成自动通过或失败的证明。
