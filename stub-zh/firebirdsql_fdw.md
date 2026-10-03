## 用法

来源：

- [Official README](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/README.md)
- [Extension control file](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/firebirdsql_fdw.control)
- [Implementation (lib.rs)](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/src/lib.rs)
- [Cargo package metadata](https://github.com/nakagami/firebirdsql_fdw/blob/c223920d389c3a247d89388472f73d65f5660a0c/Cargo.toml)

`firebirdsql_fdw` 是 Rust/pgrx 实现的 Firebird 包装器。创建扩展后，实际 SQL FDW 名称为 `firebird_fdw`，因此可能与已经拥有该包装器名称的另一扩展冲突。

### 基本用法

安装与 PostgreSQL 13–18 中实际版本匹配的构建。远端凭据保存在服务器选项中，应限制目录访问和服务器定义权限。

```sql
CREATE EXTENSION firebirdsql_fdw;
CREATE SERVER my_firebird FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (host 'localhost', port '3050', db_name '/firebird/data/mydb.fdb',
           username 'app_user', password 'replace-with-password');
CREATE FOREIGN TABLE employees (id integer, name text, hired date)
  SERVER my_firebird OPTIONS (table 'EMPLOYEES', rowid_column 'ID');
SELECT * FROM employees WHERE id = 1;
```

### 对象与限制

`IMPORT FOREIGN SCHEMA firebird` 可发现远端表。读取支持谓词、排序及行数限制下推，也提供插入、更新、删除；更新和删除须通过 `rowid_column` 识别行。客户端使用 `firebirust`，未声明其他 SQL 扩展依赖。control 版本 `0.1.0` 不可迁移，设置了 `superuser=false`，但创建基于 C 函数的 FDW 仍需相应 PostgreSQL 权限。不能假设跨系统事务具有原子性。
