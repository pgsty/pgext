## 用法

来源：[README](https://github.com/plv8/plv8/blob/r3.2/README.md)，[Docs site](https://plv8.github.io/)，[Built-ins](https://github.com/plv8/plv8/blob/r3.2/docs/BUILTINS.md)，[Runtime configuration](https://github.com/plv8/plv8/blob/r3.2/docs/CONFIGURATION.md)，[Tag v3.2.4](https://github.com/plv8/plv8/tree/v3.2.4)

`plv8` 是 PostgreSQL 的 trusted JavaScript procedural language，由 V8 引擎驱动。上游当前将扩展标记为 `v3.2.4`；Pigsty 的 `3.2.4-2` 包版本只是打包修订，不代表新的上游扩展发布。

### 基本使用

```sql
CREATE EXTENSION plv8;

SELECT plv8_version();
SELECT plv8_info();

DO $$ plv8.elog(NOTICE, plv8.version); $$ LANGUAGE plv8;

CREATE FUNCTION plv8_test(keys text[], vals text[]) RETURNS json AS $$
  let out = {};
  for (let i = 0; i < keys.length; i++) out[keys[i]] = vals[i];
  return out;
$$ LANGUAGE plv8 IMMUTABLE STRICT;
```

### 常用 built-ins

- `plv8.elog(level, ...)`：输出 PostgreSQL 日志或客户端消息。
- `plv8.execute(sql [, args])`：执行 SQL，并返回结果行或受影响行数。
- `plv8.prepare(...)`、`PreparedPlan.execute()`、`PreparedPlan.cursor()`：提供预编译 SPI 访问。
- `plv8.subtransaction(fn)`：以原子方式执行一组 SPI 操作。
- `plv8.find_function(...)`：按名称调用另一个 PLV8 函数。
- `plv8.memory_usage()`：查看当前会话的 V8 heap 使用情况。
- `plv8.run_script(source, name)`：执行具名脚本文本。

### 运行时设置

```sql
SET plv8.start_proc = 'plv8_init';
SET plv8.execution_timeout = 60;
SET plv8.memory_limit = 512;
```

- `plv8.start_proc`
- `plv8.v8_flags`
- `plv8.execution_timeout`
- `plv8.memory_limit`

### 注意事项

- 当前文档说明支持 PostgreSQL 13 及以上版本。
- 每个 session 都有独立的全局 JavaScript runtime；切换 role 会初始化单独的 runtime context。
- `plv8.execution_timeout` 仅在扩展以 execution-timeout 支持编译时生效。
