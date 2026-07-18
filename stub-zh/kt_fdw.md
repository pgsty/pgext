## 用法

来源：

- [已复核 commit 的归档上游 README](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/README.md)
- [FDW SQL 定义](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/sql/kt_fdw.sql)
- [扩展 control 文件](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/kt_fdw.control)

`kt_fdw` 将 Kyoto Tycoon 键值服务器暴露为 PostgreSQL 双列外部表。它针对文本键和值支持 `SELECT`、`INSERT`、`UPDATE` 与 `DELETE`。

```sql
CREATE EXTENSION kt_fdw;

CREATE SERVER kt_server
  FOREIGN DATA WRAPPER kt_fdw
  OPTIONS (host '127.0.0.1', port '1978', timeout '-1');

CREATE USER MAPPING FOR PUBLIC SERVER kt_server;
CREATE FOREIGN TABLE cache_entry (key text, value text)
  SERVER kt_server;

SELECT * FROM cache_entry WHERE key = 'example';
```

事务模式要求 Kyoto Tycoon 在构建时启用 Lua，并由 `ktserver` 加载随附的 `transactions.lua`。构建时可移除 `USE_TRANSACTIONS` 来禁用该集成。

### 注意事项

该仓库现位于 Cloudflare 的 archive 组织中，最后一次源码变更可追溯到 2013 年；上游也没有说明它与当前 PostgreSQL 或 Kyoto Tycoon 版本的兼容性。远端服务不处于 PostgreSQL 的持久性与访问控制边界内；将其用于重要数据前，必须测试故障、超时和事务行为。
