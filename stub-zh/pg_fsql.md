
## 用法

> 来源：[README](https://github.com/yurc/pg_fsql/blob/main/README.md), [control file](https://raw.githubusercontent.com/yurc/pg_fsql/main/pg_fsql.control)

`pg_fsql` 是 PostgreSQL 的递归 SQL 模板引擎。它把基于 C 的占位符渲染、PL/pgSQL 模板执行、层级化模板组合，以及可选的 SPI plan cache 结合在一起。上游项目强调它不需要 superuser 权限。

### 快速开始

```sql
CREATE EXTENSION pg_fsql;

INSERT INTO fsql.templates (path, cmd, body)
VALUES ('user_count', 'exec',
        'SELECT jsonb_build_object(''total'', count(*))
         FROM users WHERE status = {d[status]!r}');

SELECT fsql.run('user_count', '{"status":"active"}');
SELECT fsql.render('user_count', '{"status":"active"}');
```

### 目录表

扩展会安装两个主要目录表：

```sql
fsql.templates (
    path varchar(500) primary key,
    cmd varchar(50),
    body text,
    defaults text,
    cached boolean default false
)

fsql.params (
    key_param varchar(255) primary key,
    type_param varchar(255) not null
)
```

`path` 采用点号分隔，用于定义模板层级。

### 命令与占位符

README 记录了六种命令类型：

- `exec`，执行 SQL 并返回 `jsonb`
- `ref`，重定向到另一个模板
- `if`，选择子分支
- `exec_tpl`，执行 SQL 后把结果重新作为模板渲染
- `map`，把子节点收集成 JSON 对象
- `NULL`，作为插入父模板的文本片段

渲染器支持这些占位符：

- `{d[key]}`
- `{d[key]!r}`，对应 `quote_literal`
- `{d[key]!j}`，对应 JSONB 字面量
- `{d[key]!i}`，对应 `quote_identifier`

特殊键 `_self` 会注入完整输入 JSON 对象。

### 公共 API

上游公开函数包括：

- `fsql.run(path, data, debug)`，执行模板树
- `fsql.render(path, data)`，预览渲染后的 SQL
- `fsql.tree(path)`，查看层级结构
- `fsql.explain(path, data)`，跟踪展开过程
- `fsql.validate()`，检查模板
- `fsql.depends_on(path)`，查看依赖关系
- `fsql.clear_cache()`，释放缓存的 SPI plans

### 层级模板示例

```sql
INSERT INTO fsql.templates (path, cmd, body) VALUES
    ('report', 'exec',
     'SELECT jsonb_build_object(''data'', array_agg(row_to_json(t)))
      FROM (SELECT {d[cols]} FROM {d[src]} {d[where]}) t'),
    ('report.cols',  NULL, 'id, name, email'),
    ('report.src',   NULL, 'customers'),
    ('report.where', NULL, 'WHERE city = {d[city]!r}');

SELECT fsql.run('report', '{"city":"Moscow"}');
SELECT fsql.render('report', '{"city":"Moscow"}');
```

### 说明

README 列出 PostgreSQL 14+ 和 `plpgsql`。control file 当前仍声明 SQL extension version `1.0`，尽管包任务跟踪的 release 为 `1.1.0`。仓库没有额外的官方 release notes，因此用户侧行为仍应以当前 README 为准。
