## 用法

来源：

- [1.10.2 README](https://github.com/orioledb/orioledb/blob/v1.10.2/README.md)
- [1.10.2 release](https://github.com/orioledb/orioledb/releases/tag/v1.10.2)
- [Control file](https://github.com/orioledb/orioledb/blob/v1.10.2/orioledb.control)
- [Patched PostgreSQL](https://github.com/orioledb/postgres)
- [Patched server versions](https://github.com/orioledb/orioledb/blob/v1.10.2/.pgtags)

`orioledb` 是一种替代表存储引擎，采用基于撤销日志的 MVCC、写时复制检查点和行级 WAL。它要求专门打补丁的 PostgreSQL 服务器，目前仍处于公测阶段。

### 配置

在 `postgresql.conf` 中添加（需要重启）：

```ini
shared_preload_libraries = 'orioledb.so'
```

然后启用扩展：

```sql
CREATE EXTENSION orioledb;
```

### 创建表

使用 `USING orioledb` 子句创建采用 OrioleDB 存储引擎的表：

```sql
CREATE TABLE my_table (
    id serial PRIMARY KEY,
    name text,
    value numeric
) USING orioledb;
```

可在上游兼容性限制内对 OrioleDB 表使用普通 DML：

```sql
INSERT INTO my_table (name, value) VALUES ('test', 42);
SELECT * FROM my_table WHERE id = 1;
UPDATE my_table SET value = 100 WHERE id = 1;
DELETE FROM my_table WHERE id = 1;
```

### 排序规则要求

OrioleDB 表仅支持 **ICU**、**C** 和 **POSIX** 排序规则。为避免在每个文本字段上指定 COLLATE，请使用适当的默认值创建数据库：

```sql
CREATE DATABASE mydb LOCALE 'C' TEMPLATE template0;
-- OR
CREATE DATABASE mydb LOCALE_PROVIDER icu ICU_LOCALE 'en' TEMPLATE template0;
```

### 兼容性与 1.10.2 维护说明

本版 `.pgtags` 将 PostgreSQL 16、17、18 分别固定到 `patches16_50`、`patches17_23`、`patches18_4`。普通 PostgreSQL 安装不能满足要求。排序规则仅支持 ICU、C 和 POSIX。撤销日志与页面合并降低了原生表的清理需求，其他 PostgreSQL 表仍需要正常维护。

1.10.2 保持 SQL 版本 `1.10`，不改变存储或 WAL 格式。它修复了 `ORIOLEDB_VERSION` 报告的运行时版本、并发情况下的行版本处理、并发创建索引时捕获更新的问题，以及范围/数组/复合键导致的后台写入进程崩溃、诊断函数的 NULL 处理、表达式索引的内存泄漏和备库表空间移动。

所有节点应使用匹配的服务器补丁与库文件，在备份后遵循对应版本的升级流程。不要指定不存在的 SQL 版本 1.10.2。公测限制仍然适用，上游不建议将其普遍用于生产环境。
