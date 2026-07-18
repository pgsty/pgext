


## 用法

来源：

- [PGXN plpgsql_check 2.9.2](https://pgxn.org/dist/plpgsql_check/2.9.2/)
- [plpgsql_check README](https://github.com/okbob/plpgsql_check)
- [plpgsql_check control file](https://pgxn.org/dist/plpgsql_check/2.9.2/)

`plpgsql_check` 是 PL/pgSQL 检查器、linter、性能分析器、跟踪器和覆盖率工具。它使用 PostgreSQL 自身的解析器和执行基础设施分析 PL/pgSQL 函数体，因此许多原本只会在运行时出现的问题可以在开发或 CI 阶段提前发现。

PGXN 发布包版本是 2.9.2，但扩展 control 文件中的 SQL `default_version` 仍为 `2.9`。上游 README 声明支持 PostgreSQL 14-18。

```sql
CREATE EXTENSION IF NOT EXISTS plpgsql_check;
```

### 检查函数

```sql
SELECT *
FROM plpgsql_check_function('public.refresh_totals()');

SELECT *
FROM plpgsql_check_function('public.refresh_totals(int, text)', fatal_errors := false);
```

返回表的版本更适合过滤、落表或作为 CI 输出：

```sql
SELECT functionid, lineno, statement, sqlstate, message, level
FROM plpgsql_check_function_tb('public.refresh_totals()');
```

输出格式支持 text、JSON 和 XML：

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### 触发器函数

检查触发器函数时需要传入其操作的关系：

```sql
SELECT *
FROM plpgsql_check_function('public.audit_trigger()', 'public.accounts');

SELECT *
FROM plpgsql_check_function(
  'public.audit_trigger()',
  'public.accounts',
  newtable := 'new_rows',
  oldtable := 'old_rows'
);
```

### 警告级别

```sql
SELECT *
FROM plpgsql_check_function(
  'fx()',
  extra_warnings         := true,
  performance_warnings   := true,
  security_warnings      := true,
  compatibility_warnings := true
);
```

- `extra_warnings` 覆盖缺少返回、死代码、变量遮蔽和未使用参数。
- `performance_warnings` 覆盖隐藏 cast、类型修饰符以及可能阻止索引使用的写法。
- `security_warnings` 包括动态 SQL 和 SQL 注入风险检查。
- `compatibility_warnings` 报告过时或对版本敏感的 PL/pgSQL 模式。

### 批量检查

```sql
SELECT n.nspname, p.proname, c.*
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON p.pronamespace = n.oid
JOIN pg_catalog.pg_language l ON l.oid = p.prolang
CROSS JOIN LATERAL plpgsql_check_function_tb(p.oid) AS c
WHERE l.lanname = 'plpgsql'
  AND n.nspname NOT IN ('pg_catalog', 'information_schema');
```

可以在迁移流水线中使用这种模式，在发布前发现依赖变更、列删除、不安全 cast 和 PL/pgSQL 错误。

### 被动检查

被动模式会在函数启动时检查函数。它适合开发和预生产环境，但会引入额外开销。

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'fresh_start';
```

常用设置：

```text
plpgsql_check.mode = disabled | by_function | fresh_start | every_start
plpgsql_check.fatal_errors = yes | no
plpgsql_check.show_nonperformance_warnings = false
plpgsql_check.show_performance_warnings = false
```

### 性能分析器

```sql
SELECT plpgsql_check_profiler(true);

SELECT public.refresh_totals();

SELECT lineno, exec_stmts, total_time, avg_time, source
FROM plpgsql_profiler_function_tb('public.refresh_totals()');

SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('public.refresh_totals()');

SELECT * FROM plpgsql_profiler_functions_all();
SELECT plpgsql_profiler_reset_all();
```

若要使用共享 profiler 统计并获得可靠的早期初始化，应在 `plpgsql_check` 之前预加载 `plpgsql`：

```conf
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

如果没有共享预加载，profiler 数据仅限当前会话。

### Tracer 与覆盖率

Tracer 会输出函数和语句进入/退出 notice，并可能暴露变量值。它默认关闭，必须通过 superuser 控制的设置启用。

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true, 'terse');

SELECT * FROM plpgsql_coverage_statements('public.refresh_totals()');
SELECT * FROM plpgsql_coverage_branches('public.refresh_totals()');
```

### 编译指令

可以在函数内部用 pragma 调用描述动态 SQL、临时表、推断出的 record 类型或局部检查设置：

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
DECLARE
  r record;
BEGIN
  PERFORM plpgsql_check_pragma('type: r (id int, processed bool)');
  RETURN $1::text;
END;
$$ LANGUAGE plpgsql;
```

### 注意事项

- `plpgsql_check` 依赖 `plpgsql`。
- 主动检查不一定需要 preload，但共享 profiler 存储和稳定的 tracer/profiler 早期初始化需要 preload。
- tracer 输出可能包含函数参数和局部变量值；不要在敏感生产负载上大范围启用。
- 检查器无法完美理解所有动态 SQL 字符串。可使用 pragmas 描述预期的动态对象，减少误报。
