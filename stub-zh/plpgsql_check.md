## 用法

来源：[official README](https://github.com/okbob/plpgsql_check)、[v2.9.1 release notes](https://github.com/okbob/plpgsql_check/releases/tag/v2.9.1)、[local package metadata](../db/extension.csv)。

`plpgsql_check` 是面向 PL/pgSQL 的检查器、linter、性能分析器、跟踪器和覆盖率工具。它可以在开发阶段发现许多错误，而不是等到运行时才失败。

```sql
CREATE EXTENSION plpgsql_check;
```

### 检查函数

```sql
SELECT * FROM plpgsql_check_function('my_function()');
SELECT * FROM plpgsql_check_function('my_function(int, text)');
SELECT * FROM plpgsql_check_function('my_function()', fatal_errors := false);
```

返回表的变体适合生成结构化报告：

```sql
SELECT *
FROM plpgsql_check_function_tb('my_function()');
```

### 输出格式

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### 检查触发器函数

```sql
SELECT * FROM plpgsql_check_function('my_trigger_func()', 'my_table');

SELECT * FROM plpgsql_check_function(
  'my_trigger_func()',
  'my_table',
  newtable := 'newtab',
  oldtable := 'oldtab'
);
```

### 警告类别

```sql
SELECT * FROM plpgsql_check_function(
  'fx()',
  extra_warnings := true,
  performance_warnings := true,
  security_warnings := true,
  compatibility_warnings := true
);
```

- `extra_warnings` 覆盖缺少返回、死代码、未使用参数等问题。
- `performance_warnings` 覆盖性能相关检查。
- `security_warnings` 包括 SQL 注入风险等检查。
- `compatibility_warnings` 报告过时或对版本敏感的 PL/pgSQL 模式。

### 批量检查所有函数

```sql
SELECT p.oid, p.proname, plpgsql_check_function(p.oid)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON p.pronamespace = n.oid
JOIN pg_catalog.pg_language l ON p.prolang = l.oid
WHERE l.lanname = 'plpgsql'
  AND p.prorettype <> 'pg_catalog.trigger'::pg_catalog.regtype;
```

### 被动模式

被动模式会在函数启动时检查函数。它适合开发或预生产环境，因为会带来额外开销。

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'every_start';
```

常用设置：

```text
plpgsql_check.mode = [ disabled | by_function | fresh_start | every_start ]
plpgsql_check.fatal_errors = [ yes | no ]
plpgsql_check.show_nonperformance_warnings = false
plpgsql_check.show_performance_warnings = false
```

### 性能分析器

```sql
SELECT plpgsql_check_profiler(true);

SELECT my_function();

SELECT lineno, avg_time, source
FROM plpgsql_profiler_function_tb('my_function()');

SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('my_function()');

SELECT * FROM plpgsql_profiler_functions_all();
SELECT plpgsql_profiler_reset_all();
```

若要使用共享性能分析统计，请在 `plpgsql_check` 之前预加载 `plpgsql`：

```text
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

没有共享预加载时，性能分析数据仅限于当前活动会话。

### 跟踪器

跟踪会为函数和语句的进入/退出发出 notice，并可能暴露变量值。它默认关闭，只应通过超级用户控制的设置启用。

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true);
SET plpgsql_check.tracer_verbosity = terse;
```

### 依赖追踪

```sql
SELECT *
FROM plpgsql_show_dependency_tb('my_function(int)');
```

### 覆盖率指标

```sql
SELECT * FROM plpgsql_coverage_statements('my_function()');
SELECT * FROM plpgsql_coverage_branches('my_function()');
```

### Pragma 指令

在函数内部使用 pragma 调用，告诉 `plpgsql_check` 动态 SQL、临时表、推断出的 record 类型，或局部检查设置：

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
BEGIN
  PERFORM plpgsql_check_pragma('type: r (id int, processed bool)');
  RETURN $1::text;
END;
$$ LANGUAGE plpgsql;
```

### 注意事项

- Pigsty 将 `plpgsql_check` 2.9.1 作为 PostgreSQL 14-18 的 RPM 打包；DEB 包来自 PGDG。
- 该扩展要求 `plpgsql`。主动检查不强制要求预加载，但共享性能分析存储以及跟踪器/性能分析器的可靠早期初始化需要预加载。
- v2.9.1 修复了被跟踪 inline block 失败时可能发生的崩溃；这个补丁发布没有记录新的面向用户 SQL API。
- 跟踪器可能暴露函数参数或变量值。对 security-definer 函数或敏感数据使用时要格外谨慎。
