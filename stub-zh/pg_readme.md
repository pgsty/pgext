

## 用法

> [pg_readme: 从 PostgreSQL COMMENT 对象自动生成 README 文档](https://github.com/bigsmoke/pg_readme)

根据 `pg_description` 系统目录中的 `COMMENT` 对象，为扩展或模式生成 `README.md` 文档。

### 生成扩展 README

```sql
SELECT pg_extension_readme('my_extension'::name);
```

### 生成模式 README

```sql
SELECT pg_schema_readme('my_schema'::regnamespace);
```

### 处理指令

在 `COMMENT ON EXTENSION` 或 `COMMENT ON SCHEMA` 中包含以下 XML 处理指令：

- `<?pg-readme-reference?>` -- 替换为完整的对象引用
- `<?pg-readme-colophon?>` -- 添加包含 pg_readme 信息的版权页

### 典型工作流程

在扩展中创建一个生成 README 的函数：

```sql
CREATE FUNCTION my_ext_readme() RETURNS text
    VOLATILE SET search_path FROM current
    SET pg_readme.include_view_definitions TO 'true'
    LANGUAGE plpgsql AS $$
DECLARE _readme text;
BEGIN
    CREATE EXTENSION IF NOT EXISTS pg_readme WITH VERSION '0.1.0';
    _readme := pg_extension_readme('my_extension'::name);
    RAISE transaction_rollback;
EXCEPTION WHEN transaction_rollback THEN RETURN _readme;
END; $$;
```

然后通过以下命令生成：`make README.md`

### 设置

| 设置项 | 默认值 |
|--------|--------|
| `pg_readme.include_view_definitions` | `true` |
| `pg_readme.include_routine_definitions_like` | `'{test__%}'` |
| `pg_readme.include_this_routine_definition` | `null` |
