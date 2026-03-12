

## 用法

> [plr: 加载 R 解释器并在数据库中执行 R 脚本](https://github.com/postgres-plr/plr)

`plr` 允许在 PostgreSQL 中使用 R 编程语言编写函数，提供对 R 统计和数据分析功能的完整访问。

```sql
CREATE EXTENSION plr;
```

### 创建函数

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (arg1 > arg2)
  return(arg1)
else
  return(arg2)
' LANGUAGE plr STRICT;

SELECT r_max(10, 20);  -- 20
```

使用命名参数：

```sql
CREATE OR REPLACE FUNCTION sd(vals float8[]) RETURNS float AS '
sd(vals)
' LANGUAGE plr STRICT;

SELECT sd(ARRAY[1.0, 2.0, 3.0, 4.0, 5.0]);
```

### 参数处理

- 参数以 `arg1`、`arg2` 等形式或按命名参数方式访问
- NULL 参数变为 R 的 `NA` 值（除非函数声明为 `STRICT`）
- 复合类型（行）以 R data.frame 形式传递
- 数组以 R 向量形式传递

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (is.null(arg1) && is.null(arg2))
  return(NULL)
if (is.null(arg1))
  return(arg2)
if (is.null(arg2))
  return(arg1)
if (arg1 > arg2)
  return(arg1)
return(arg2)
' LANGUAGE plr;
```

### 通过 SPI 访问数据库

```sql
CREATE OR REPLACE FUNCTION test_spi(text) RETURNS SETOF record AS '
pg.spi.exec(arg1)
' LANGUAGE plr;

SELECT * FROM test_spi('SELECT oid, typname FROM pg_type LIMIT 5')
  AS t(oid oid, typname name);
```

预备语句：

```sql
-- 预备
sp <<- pg.spi.prepare('SELECT * FROM pg_type WHERE typname = $1', c(NAMEOID))
-- 执行
pg.spi.execp(sp, list('text'))
```

### 集合返回函数

返回 data.frame 以实现集合返回函数：

```sql
CREATE OR REPLACE FUNCTION get_numbers(n int) RETURNS SETOF integer AS '
1:arg1
' LANGUAGE plr;

SELECT * FROM get_numbers(5);
```

### 窗口函数

```sql
CREATE OR REPLACE FUNCTION r_regr_slope(float8, float8, int)
RETURNS float8 AS '
slope <- NA
y <- farg1
x <- farg2
if (fnumrows == arg3 + 1L)
  try(slope <- lm(y ~ x)$coefficients[2])
return(slope)
' LANGUAGE plr WINDOW;
```

窗口函数接收 `farg1..fargN`（窗口帧内的值向量）、`fnumrows`（帧大小）和 `prownum`（分区中的当前行位置）。

### 全局变量

使用 R 的全局环境在函数调用之间保持数据：

```sql
CREATE OR REPLACE FUNCTION set_state(key text, val text) RETURNS void AS '
assign(arg1, arg2, env=.GlobalEnv)
' LANGUAGE plr;
```

### 实用辅助函数

```sql
SELECT load_r_typenames();  -- 加载类型 OID 变量
SELECT * FROM r_typenames(); -- 列出可用的类型 OID
SELECT plr_version();        -- PL/R 版本
```

### 触发器函数

PL/R 支持触发器函数，可以访问 `pg.tg.name`、`pg.tg.relname`、`pg.tg.when`、`pg.tg.level`、`pg.tg.op`、`pg.tg.new` 和 `pg.tg.old`。
