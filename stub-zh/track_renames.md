## 用法

来源：

- [目录版本对应的官方 README](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/README.md)
- [目录版本对应的扩展 SQL](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/track_renames--0.1.sql)
- [目录版本对应的事件触发器实现](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/track_renames.c)

`track_renames` 0.1 提供一个事件触发器函数，在 PostgreSQL 重命名命令后调用应用回调。实际实现要求回调接受五个 text 参数；旧 README 展示的四参数回调与固定版本 C 源码不符。

### 核心流程

创建一个能以非限定名称找到的回调，设置回调名称，再注册事件触发器。创建事件触发器需要超级用户权限。

```sql
CREATE EXTENSION track_renames;

CREATE TABLE rename_log (
    object_type text,
    schema_name text,
    object_name text,
    subobject_name text,
    new_name text,
    changed_at timestamptz DEFAULT clock_timestamp()
);

CREATE FUNCTION log_rename(text, text, text, text, text)
RETURNS void
LANGUAGE sql
AS $$
  INSERT INTO rename_log
  VALUES ($1, $2, $3, $4, $5, clock_timestamp());
$$;

SET track_renames.function = 'log_rename';

CREATE EVENT TRIGGER track_object_renames
ON ddl_command_end
WHEN TAG IN ('ALTER TABLE', 'ALTER FUNCTION', 'ALTER TYPE', 'ALTER VIEW')
EXECUTE FUNCTION track_renames();
```

只有确认每个受影响会话的 search path 都能找到回调后，才应把设置持久化到数据库或角色级别。

### 回调契约

- 预期的回调返回 void，并且必须接受五个 text 参数：对象类型、schema 名、对象名、子对象名和新名称。
- 对于解析树未暴露关系或子对象的重命名形式，部分参数会是 NULL。
- 被拦截命令的解析树不是重命名语句时，事件触发器会忽略它。

### 运维限制

- 回调按非限定函数名和会话 search path 查找。应锁定 search path，并禁止不可信用户在可见 schema 中创建影子函数。
- 回调错误发生在 DDL 事务内，可能使重命名失败。日志代码应保持短小、确定并且只使用本地事务。
- 源码面向老版本 PostgreSQL 解析树 API 和枚举。必须针对每个目标服务器大版本以及依赖的每类重命名执行编译和回归测试。
