## 用法

来源：

- [plpgsql_check 2.10.1 README](https://github.com/okbob/plpgsql_check/blob/v2.10.1/README.md)
- [plpgsql_check 2.10.1 release](https://github.com/okbob/plpgsql_check/releases/tag/v2.10.1)
- [plpgsql_check 2.10.1 control file](https://github.com/okbob/plpgsql_check/blob/v2.10.1/plpgsql_check.control)
- [plpgsql_check 2.9.2 to 2.10.1 changes](https://github.com/okbob/plpgsql_check/compare/v2.9.2...v2.10.1)

`plpgsql_check` 是一个 PL/pgSQL 检查器、语法检查工具、性能分析器、跟踪器和覆盖率工具。它使用 PostgreSQL 自身的解析器和执行器基础设施来分析 PL/pgSQL 函数体，因此许多问题可以在开发或 CI 阶段被发现而不会在运行时才出现。

2.10.1 版本安装的是 SQL 扩展 `2.10`。支持 PostgreSQL 14-18；上游源代码还包含对后续 PostgreSQL 开发分支的兼容性工作。

```sql
CREATE EXTENSION IF NOT EXISTS plpgsql_check;
```

### 检查一个函数

```sql
SELECT *
FROM plpgsql_check_function('public.refresh_totals()');

SELECT *
FROM plpgsql_check_function('public.refresh_totals(int, text)', fatal_errors := false);
```

返回表的变体更容易过滤、存储或用于 CI 输出：

```sql
SELECT functionid, lineno, statement, sqlstate, message, level
FROM plpgsql_check_function_tb('public.refresh_totals()');
```

输出格式包括文本、JSON 和 XML：

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### 触发函数

触发函数需要它们操作的关系：

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

- `extra_warnings` 包括缺少返回值、死代码、遮蔽变量和未使用的参数。
- `performance_warnings` 包括隐式转换、类型修饰符以及可能阻止索引使用的模式。
- `security_warnings` 包括动态 SQL 和 SQL 注入风险检查。
- `compatibility_warnings` 报告过时或版本敏感的 PL/pgSQL 模式。

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

在迁移管道中使用此模式来捕获更改依赖关系、删除列、不安全转换和 PL/pgSQL 错误，以确保发布前的安全性。

### 被动检查

被动模式会在函数启动时进行检查。它适用于开发和预生产环境，但会增加开销。

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

### 分析器

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

为了共享分析器统计信息并可靠地早期初始化，在加载 `plpgsql` 之前预加载 `plpgsql_check`：

```conf
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

如果没有共享预加载，分析器数据将仅限于当前会话。

### 跟踪和覆盖率

跟踪会在函数和语句的入口/出口处发出通知，并可以暴露变量值。默认情况下禁用此功能，必须由超级用户控制设置启用。

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true, 'terse');

SELECT * FROM plpgsql_coverage_statements('public.refresh_totals()');
SELECT * FROM plpgsql_coverage_branches('public.refresh_totals()');
```

### 断言

在函数内部使用断言调用来描述动态 SQL、临时表、推断出的记录类型或局部检查设置：

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

版本 2.10 添加了 `plpgsql_make_pragma(regprocedure)`，它会计划 `CREATE TEMP TABLE ... AS SELECT|VALUES|TABLE` 语句而不执行它们，并返回可以提供给检查的表断言。

```sql
SELECT *
FROM plpgsql_check_function(
  'public.refresh_stage()'::regprocedure,
  pragmas => ARRAY(
    SELECT plpgsql_make_pragma('public.refresh_stage()'::regprocedure)
  )
);
```

### 注意事项

- `plpgsql_check` 需要 `plpgsql`。
- 活动检查可选地预加载，但共享分析器存储和稳健的跟踪/分析器初始化需要预加载。
- 跟踪输出可能包括函数参数和局部变量值；不要在敏感生产工作负载中广泛启用它。
- 检查器无法完全理解每个动态 SQL 字符串。使用断言来记录预期的动态对象并减少误报。
- 2.10.1 版本修复了 PostgreSQL 14 中的共享内存 LWLock tranche 注册问题；在预加载分析器/跟踪状态时，请使用此版本而不是初始 2.10.0 构建。
