## 用法

来源：

- [README](https://github.com/plv8/plv8/blob/v3.2.5/README.md)
- [Built-ins](https://github.com/plv8/plv8/blob/v3.2.5/docs/BUILTINS.md)
- [Configuration](https://github.com/plv8/plv8/blob/v3.2.5/docs/CONFIGURATION.md)
- [Changes](https://github.com/plv8/plv8/blob/v3.2.5/Changes)
- [Control](https://github.com/plv8/plv8/blob/v3.2.5/plv8.control.common)
- [SQL template](https://github.com/plv8/plv8/blob/v3.2.5/plv8.sql.common)

`plv8` 基于 V8 为 PostgreSQL 提供可信的 JavaScript 存储过程语言。本文依据上游 3.2.5，包含有效用户上下文相关的崩溃修复与 PostgreSQL 19 支持。

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

- 3.2.5 的 CI 矩阵覆盖 PostgreSQL 14–19；上游主版本支持与下游软件包可用范围需分别看待。
- 创建扩展需要超级用户；所安装的 JavaScript 语言属于可信语言，并不意味着普通用户可直接安装该扩展。安装 SQL 撤销了 PUBLIC 对 `plv8_info()` 的执行权限。
- 3.2.5 修复了缓存函数经 `SECURITY DEFINER`、`SET ROLE` 切换有效用户或在 `plv8_reset()` 后执行时的崩溃，也修复了异常消息无法转换为字符串时的崩溃。
- 每个 session 都有独立的全局 JavaScript runtime；切换 role 会初始化单独的 runtime context。
- `plv8.execution_timeout` 仅在扩展以 execution-timeout 支持编译时生效。
