## 用法

来源：

- [官方 README](https://github.com/masaofujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/README.md)
- [1.0 扩展 SQL](https://github.com/masaofujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/pg_fallback_utf8_to_euc_jp--1.0.sql)

`pg_fallback_utf8_to_euc_jp` 提供另一种 UTF-8 到 EUC_JP 的编码转换。它适用于需要用回退转换器替换 PostgreSQL 内置默认转换器的数据库；仅安装扩展并不会改变 PostgreSQL 选择的转换器。

### 核心流程

在所有可能发生该转换的数据库中创建扩展，再使用管理会话切换默认转换标志。

```sql
CREATE EXTENSION pg_fallback_utf8_to_euc_jp;

BEGIN;
UPDATE pg_conversion
SET condefault = false
WHERE conname = 'utf8_to_euc_jp';

UPDATE pg_conversion
SET condefault = true
WHERE conname = 'pg_fallback_utf8_to_euc_jp';
COMMIT;
```

断开连接并建立新会话后，再依赖新的默认转换。

### 对象与恢复边界

- `pg_fallback_utf8_to_euc_jp`：扩展为 UTF-8 到 EUC_JP 创建的转换对象。

修改 `pg_conversion` 是直接操作系统目录，需要具备相应管理权限。必须确保该编码对始终只有一个默认转换，并在删除扩展前反向切换这两个标志。

上游明确指出，`pg_dump` 与 `pg_dumpall` 不会保留这些目录标志变更。恢复转储后，应在所需数据库中创建扩展并重新执行默认选择事务。该扩展可重定位，不需要预加载或重启服务器。
