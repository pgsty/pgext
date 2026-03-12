

## 用法

> [plpgsql_check: PL/pgSQL 函数的增强检查器](https://github.com/okbob/plpgsql_check)

`plpgsql_check` 是 PL/pgSQL 函数的代码检查和分析工具，能够在开发阶段而非运行时检测错误。

```sql
CREATE EXTENSION plpgsql_check;
```

### 检查函数

```sql
SELECT * FROM plpgsql_check_function('my_function()');
SELECT * FROM plpgsql_check_function('my_function(int, text)');
SELECT * FROM plpgsql_check_function('my_function()', fatal_errors := false);
```

### 输出格式

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### 检查触发器函数

```sql
-- 触发器函数需要关联的表
SELECT * FROM plpgsql_check_function('my_trigger_func()', 'my_table');

-- 使用过渡表
SELECT * FROM plpgsql_check_function(
  'my_trigger_func()', 'my_table',
  newtable := 'newtab', oldtable := 'oldtab'
);
```

### 警告类别

```sql
SELECT * FROM plpgsql_check_function('fx()',
  extra_warnings := true,         -- 死代码、未使用的参数
  performance_warnings := true,   -- 索引和类型转换问题
  security_warnings := true,      -- SQL 注入检查
  compatibility_warnings := true  -- 过时的模式
);
```

### 批量检查所有函数

```sql
SELECT p.oid, p.proname, plpgsql_check_function(p.oid)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON pronamespace = n.oid
JOIN pg_catalog.pg_language l ON p.prolang = l.oid
WHERE l.lanname = 'plpgsql' AND p.prorettype <> 2279;
```

### 被动模式（执行时检查）

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'every_start';  -- 每次执行前检查
```

或在 `postgresql.conf` 中配置：

```
shared_preload_libraries = 'plpgsql,plpgsql_check'
plpgsql_check.mode = 'every_start'
```

### 性能分析器

```sql
-- 启用性能分析
SELECT plpgsql_check_profiler(true);

-- 执行函数以收集数据
SELECT my_function();

-- 查看每行执行时间
SELECT lineno, avg_time, source
FROM plpgsql_profiler_function_tb('my_function()');

-- 每条语句的性能分析
SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('my_function()');

-- 所有函数的统计信息
SELECT * FROM plpgsql_profiler_functions_all();

-- 重置性能分析数据
SELECT plpgsql_profiler_reset_all();
```

### 依赖追踪

```sql
SELECT * FROM plpgsql_show_dependency_tb('my_function(int)');
```

### 覆盖率指标

```sql
SELECT * FROM plpgsql_coverage_statements('my_function()');
SELECT * FROM plpgsql_coverage_branches('my_function()');
```

### Pragma 指令

在函数注释中嵌入检查选项：

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
BEGIN
  /* @plpgsql_check_options: anyelementtype = text */
  RETURN $1;
END;
$$ LANGUAGE plpgsql;
```
