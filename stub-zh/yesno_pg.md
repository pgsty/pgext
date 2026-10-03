## 用法

来源：

- [Integration guide](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/docs/integrations.md)
- [Extension control file](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/yesno-pg/yesno_pg.control)
- [Installation SQL](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/yesno-pg/sql/yesno_pg--0.1.0.sql)
- [Official README](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/yesno-pg/README.md)
- [Operational guide](https://github.com/moriyoshi/yesnodb/blob/6e7ac29bc45256ec3d4e8d15c1a992ca71af19f6/docs/operations.md)

`yesno_pg` 是 yesnodb 序号集合的实验性 PostgreSQL 集成，提供外部数据包装器、索引访问方法及单列表访问方法。PostgreSQL 与 yesnodb 之间存在必须了解的持久性边界。

### 基本用法

安装针对 PostgreSQL 17 或 18 的 ABI 匹配上游产物，再由超级用户创建扩展。每个外部表将一个远端键映射为序号值集合。

```sql
CREATE EXTENSION yesno_pg;
SELECT yesno_pg_version();
CREATE SERVER yesno FOREIGN DATA WRAPPER yesno_fdw
  OPTIONS (endpoint 'https://yesno.internal:50051');
CREATE FOREIGN TABLE rust_docs (ordinal bigint NOT NULL)
  SERVER yesno OPTIONS (key '42');
SELECT count(*) FROM rust_docs;
```

### 访问方法

`yesno` 提供等值倒排列表和位图索引扫描；`yesno_table` 存储单列 `bigint` 集合。服务器的字典选项支持通过 `IMPORT FOREIGN SCHEMA` 按名称导入键。应检查执行计划，确认访问方法确实被选用。

```sql
CREATE TABLE selected_ordinals (ordinal bigint) USING yesno_table;
INSERT INTO selected_ordinals VALUES (1), (5), (9);
SELECT * FROM selected_ordinals ORDER BY ordinal;
```

### 隔离与备份

两个引擎不共享 WAL 或提交时钟。远端写入无法与 PostgreSQL 堆表写入原子提交，必须单独备份 yesnodb。对 yesno 表，`REPEATABLE READ` 固定一个版本，`READ COMMITTED` 使用语句级快照，但跨两个引擎的事务可能看到不同提交时刻。表访问方法不支持 `UPDATE`、`SELECT FOR UPDATE`、`ON CONFLICT`、`CLUSTER` 和 `TABLESAMPLE`。构建产物来自上游 Bazel/Docker 验证流程，不能将其视为通用的已打包 PostgreSQL 扩展。
