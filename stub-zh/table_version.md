

## 用法

> [table_version: PostgreSQL 表版本控制扩展](https://github.com/linz/postgresql-tableversion)

PostgreSQL 表版本控制扩展，记录行的修改及其历史。该扩展提供 API 用于获取表在特定版本的快照，以及任意两个版本之间的差异。它基于 PL/PgSQL 触发器系统来记录和访问行版本。

### 快速开始

```sql
CREATE EXTENSION table_version;

CREATE SCHEMA foo;
SET search_path TO foo, public;

CREATE TABLE foo.bar (
    id INTEGER NOT NULL PRIMARY KEY,
    baz TEXT
);

-- 启用版本控制
SELECT table_version.ver_enable_versioning('foo', 'bar');

-- 创建版本并插入数据
SELECT table_version.ver_create_revision('Insert data');
INSERT INTO foo.bar (id, baz) VALUES
  (1, 'foo bar 1'),
  (2, 'foo bar 2'),
  (3, 'foo bar 3');
SELECT table_version.ver_complete_revision();

-- 查看两个版本之间的差异
SELECT * FROM table_version.ver_get_foo_bar_diff(1001, 1002);
```


## 工作原理

启用表版本控制后，原始表数据保持不变，系统会创建一张新的版本表，包含原表所有字段加上 `_revision_created` 和 `_revision_expired` 字段。原表上会设置行级触发器，记录每一次插入、更新和删除到版本数据表中。同时设置语句级触发器，禁止 TRUNCATE 操作。

### 表的前置条件

- 表必须有唯一的非复合 integer、bigint、text 或 varchar 列
- 表不能是临时表


## 自动版本

如果你不想显式调用 `ver_create_revision` 和 `ver_complete_revision`，自动版本模式会按事务对编辑进行分组：

```sql
SELECT table_version.ver_enable_versioning('foo', 'bar');

BEGIN;
INSERT INTO foo.bar (id, baz) VALUES (1, 'foo bar 1');
INSERT INTO foo.bar (id, baz) VALUES (2, 'foo bar 2');
COMMIT;

BEGIN;
UPDATE foo.bar SET baz = 'foo bar 1 edit' WHERE id = 1;
COMMIT;

SELECT * FROM table_version.foo_bar_revision;
```

版本消息会根据事务 ID 自动创建。


## 使用表差异复制数据

要在远程系统上维护一份表数据的副本：

```sql
-- 1. 确定哪些表已启用版本控制
SELECT * FROM table_version.ver_get_versioned_tables();

-- 2. 获取基准版本
SELECT table_version.ver_get_table_base_revision('foo', 'bar');

-- 3. 创建基准快照
CREATE TABLE foo_bar_copy AS
SELECT * FROM table_version.ver_get_foo_bar_revision(
    table_version.ver_get_table_base_revision('foo', 'bar')
);

-- 4. 获取差异以进行增量更新
SELECT * FROM table_version.ver_get_foo_bar_diff(
    my_last_revision,
    table_version.ver_get_table_last_revision('foo', 'bar')
);
```


## 安全模型

- 任何人都可以创建版本
- 版本只能由其创建者完成
- 只有拥有表所有权权限的人才能启用/禁用版本控制
- 只有空版本可以被删除
- 只有版本创建者才能删除该版本

注意：禁用表的版本控制会导致该表的所有历史记录被删除。


## 主要函数

| 函数 | 说明 |
|------|------|
| `ver_enable_versioning(schema, table)` | 启用表的版本控制 |
| `ver_disable_versioning(schema, table)` | 禁用版本控制并删除历史 |
| `ver_create_revision(comment)` | 创建新版本 |
| `ver_complete_revision()` | 标记当前版本为已完成 |
| `ver_get_<schema>_<table>_diff(rev1, rev2)` | 获取两个版本之间的差异 |
| `ver_get_<schema>_<table>_revision(rev)` | 获取特定版本的快照 |
| `ver_get_versioned_tables()` | 列出所有已版本控制的表 |
| `ver_get_last_revision()` | 获取最后的版本号 |
