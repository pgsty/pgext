


## 用法

来源：[README](https://github.com/orioledb/orioledb)、[beta16 release](https://github.com/orioledb/orioledb/releases/tag/beta16)、[patched PostgreSQL tree](https://github.com/orioledb/postgres)

OrioleDB 是 PostgreSQL 的新型存储引擎，为数据库容量、能力和性能提供现代化方案。它使用基于撤销日志的 MVCC、写时复制检查点和行级 WAL，消除了膨胀问题和 VACUUM 的需求。

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

所有标准 PostgreSQL 操作均可用于 OrioleDB 表：

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
-- 或
CREATE DATABASE mydb LOCALE_PROVIDER icu ICU_LOCALE 'en' TEMPLATE template0;
```

### 主要优势

- **无膨胀**：基于撤销日志的 MVCC 意味着旧元组版本不会膨胀主存储
- **无需 VACUUM**：页面合并和撤销日志自动回收空间
- **无回卷问题**：原生 64 位事务标识符
- **无锁页面读取**：内存页面直接链接到存储页面
- **行级 WAL**：紧凑的预写日志，适合并行回放

### 限制

- 公测阶段 -- 建议用于测试，不建议用于生产
- 需要来自 [orioledb/postgres](https://github.com/orioledb/postgres) 的补丁版 PostgreSQL 构建
- 仅支持 ICU、C 和 POSIX 排序规则

### 版本说明

OrioleDB 1.8-beta16 将扩展 SQL version 提升到 `1.8`，patched PostgreSQL builds 基于 16.13、17.9 和 18.4，并增加 PostgreSQL 18 支持。新的用户可见接口包括用于 SERIALIZABLE 支持的 `orioledb.serializable`，以及用于 `pg_amcheck` 集成的 `verify_orioledb(regclass, boolean)`。该版本还包含 recovery、replication、index-scan、vacuum 和 DDL correctness 修复。
