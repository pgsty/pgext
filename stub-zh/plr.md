## 用法

来源：

- [8.4.8.7 README](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/README.md)
- [Versioned user guide](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/userguide.md)
- [Control file](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/plr.control)
- [Version 8.4.8.7 SQL](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/plr--8.4.8.7.sql)

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

- 未命名参数以 `arg1`、`arg2` 等形式访问；显式命名的参数会替换对应的 `argN` 变量。
- SQL 标量 NULL 转为 R NULL，数组内的空元素转为 R `NA`。`STRICT` 函数会跳过整个参数为 NULL 的调用，但不会跳过包含空元素的数组。
- 复合类型（行）以 R data.frame 形式传递
- 一维数组转为 R 向量，二维数组转为矩阵，三维数组转为 R 数组；不支持更高维数。

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

调用准备参数化查询的函数之前，应在同一连接中初始化类型 OID 变量：

```sql
SELECT load_r_typenames();

CREATE OR REPLACE FUNCTION lookup_type(type_name text)
RETURNS SETOF record AS $$
sp <- pg.spi.prepare(
  'SELECT oid, typname FROM pg_type WHERE typname = $1',
  c(NAMEOID)
)
pg.spi.execp(sp, list(type_name))
$$ LANGUAGE plr;

SELECT * FROM lookup_type('text') AS t(oid oid, typname name);
```

### 集合返回函数

返回 R 向量以产生标量值集合：

```sql
CREATE OR REPLACE FUNCTION get_numbers(n int) RETURNS SETOF integer AS '
1:n
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
assign(key, val, env=.GlobalEnv)
' LANGUAGE plr;
```

### 实用辅助函数

```sql
SELECT load_r_typenames();  -- Load type OID variables
SELECT * FROM r_typenames(); -- List available type OIDs
SELECT plr_version();        -- PL/R version
```

### 触发器函数

PL/R 支持触发器函数，可以访问 `pg.tg.name`、`pg.tg.relname`、`pg.tg.when`、`pg.tg.level`、`pg.tg.op`、`pg.tg.new` 和 `pg.tg.old`。

### 运行环境与权限

本文对应 PL/R 8.4.8.7。PL/R 是不受信任的过程语言：创建其函数需要超级用户，R 代码能够以 PostgreSQL 操作系统用户的权限访问文件和进程，因此应审查函数体和 EXECUTE 授权。R 共享库必须可用；在 Unix 系统上，上游要求启动前将 `R_HOME` 设置到 PostgreSQL 服务进程的环境中，仅在交互式客户端终端设置并不能配置服务。

SQL 标量 NULL 转换为 R NULL，数组内的空元素则转换为 R NA；将函数声明为 STRICT 可以避免以空参数调用。R 全局状态属于单个后端进程，既不是跨会话共享状态，也不是持久数据库表。上游撤销了 PUBLIC 对 `plr_set_rhome(text)` 等修改环境的辅助函数的执行权限；应由管理员管理运行环境，而不是向应用角色开放这些函数。普通使用无需共享预加载。
