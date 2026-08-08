## 用法

来源：

- [plpgsql_check 2.10.4 README](https://github.com/okbob/plpgsql_check/blob/v2.10.4/README.md)
- [plpgsql_check 2.10.4 发行版](https://github.com/okbob/plpgsql_check/releases/tag/v2.10.4)
- [plpgsql_check 2.10.4 控制文件](https://github.com/okbob/plpgsql_check/blob/v2.10.4/plpgsql_check.control)
- [plpgsql_check 2.10.3 到 2.10.4 的变更](https://github.com/okbob/plpgsql_check/compare/v2.10.3...v2.10.4)

`plpgsql_check` 是面向 PL/pgSQL 的检查器、代码规范检查器、性能分析器、跟踪器和覆盖率工具。它利用 PostgreSQL 自身的解析器和执行器基础设施分析 PL/pgSQL 函数体，因此许多原本只会在运行时出现的问题，可以在开发或 CI 阶段被发现。

软件包发行版 2.10.4 安装的 SQL 扩展版本为 `2.10`。该发行版还会在依赖项输出中报告被用作声明类型的关系。

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

返回表的变体更容易筛选、存储或用于 CI 输出：

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

### 触发器函数

检查触发器函数时，需要指定其操作的关系：

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

- `extra_warnings` 覆盖缺少返回值、死代码、变量遮蔽和未使用参数。
- `performance_warnings` 覆盖隐式类型转换、类型修饰符以及可能妨碍索引使用的模式。
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

可以在迁移流水线中使用这一模式，在发布前发现依赖项变化、列被删除、不安全的类型转换以及 PL/pgSQL 错误。

### 被动检查

被动模式会在函数启动时执行检查。它适用于开发和预生产环境，但会增加开销。

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

要共享性能分析器统计信息并确保可靠的早期初始化，请在 `plpgsql_check` 之前预加载 `plpgsql`：

```conf
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

如果未共享预加载，性能分析器数据仅限当前活动会话。

### 跟踪器与覆盖率

跟踪功能会在进入和退出函数及语句时发出通知，并可能暴露变量值。它默认禁用，必须通过由超级用户控制的设置启用。

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true, 'terse');

SELECT * FROM plpgsql_coverage_statements('public.refresh_totals()');
SELECT * FROM plpgsql_coverage_branches('public.refresh_totals()');
```

### 编译指示

在函数内使用编译指示调用来描述动态 SQL、临时表、推断出的记录类型或局部检查设置：

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

版本 2.10 新增了 `plpgsql_make_pragma(regprocedure)`，它会规划临时表创建但不实际执行，并返回可提供给检查器的表编译指示。发行版 2.10.2 将其扩展至更多 `CREATE TABLE` 语句形式：

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

- `plpgsql_check` 依赖 `plpgsql`。
- 主动检查不强制要求预加载，但共享性能分析器存储以及可靠的跟踪器/性能分析器初始化需要预加载。
- 跟踪器输出可能包含函数参数和局部变量值；不要在敏感的生产工作负载中广泛启用。
- 检查器无法完美理解所有动态 SQL 字符串。请使用编译指示记录预期的动态对象，以减少误报。
- 发行版 2.10.2 和 2.10.3 修复了分析复合参数时可能发生的崩溃，以及为带多态参数的函数生成覆盖率或性能分析报告时可能发生的崩溃。版本 2.10.4 保留这些修复并改进了依赖项报告。
