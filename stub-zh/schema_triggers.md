## 用法

来源：

- [已归档的官方 README](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/README.md)
- [扩展 SQL 与事件信息类型](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/schema_triggers--0.1.sql)
- [扩展控制文件](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/schema_triggers.control)

`schema_triggers` 版本 `0.1` 为 PostgreSQL 事件触发器机制增加关系、列与触发器变更事件。该项目已归档，最后更新于 2015 年，并且会挂接 PostgreSQL 内部机制；只能用于经过明确验证的旧服务器构建。

### 启用方式

扩展既要安装数据库对象，也要加载钩子库。把库加入配置并重启 PostgreSQL，然后在需要辅助类型与函数的每个数据库中安装扩展：

```conf
shared_preload_libraries = 'schema_triggers.so'
```

```sql
CREATE EXTENSION schema_triggers;
```

仅在隔离测试中，上游 README 允许在单个会话执行 `LOAD 'schema_triggers.so'`，代替服务器预加载。

### 核心流程

```sql
CREATE FUNCTION audit_relation_create()
RETURNS event_trigger
LANGUAGE plpgsql
AS $$
DECLARE
  event_info schema_triggers.relation_create_eventinfo;
BEGIN
  event_info := schema_triggers.get_relation_create_eventinfo();
  RAISE NOTICE 'created relation %', event_info.relation;
END;
$$;

CREATE EVENT TRIGGER audit_relation_create_event
ON relation_create
EXECUTE PROCEDURE audit_relation_create();
```

### 事件与辅助函数

新增事件是 `relation_create`、`relation_alter`、`relation_drop`、`column_add`、`column_alter`、`column_drop`、`trigger_create` 和 `trigger_drop`。每个事件都有对应的 `get_*_eventinfo()` 函数，返回包含相关目录行与标识符的复合记录。

### 安全与兼容性

控制文件要求超级用户安装。触发器函数在 DDL 期间同步运行，可能拒绝或扰乱模式变更；应保持简短并妥善处理失败。事件信息记录暴露了对版本敏感的 `pg_class`、`pg_attribute` 与 `pg_trigger` 行布局。由于仓库明确无人维护且早于现代 PostgreSQL 版本，不能根据历史上的“PostgreSQL 9.3 及以上”说明推断二进制或目录兼容性。
