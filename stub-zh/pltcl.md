

## 用法

> [pltcl: PL/Tcl 受信过程语言](https://www.postgresql.org/docs/current/pltcl.html)

PL/Tcl 允许使用 Tcl 编程语言编写 PostgreSQL 函数。作为受信语言，它限制了对文件系统和其他系统资源的访问。

```sql
CREATE EXTENSION pltcl;

-- 返回问候语的简单函数
CREATE FUNCTION tcl_hello(name text) RETURNS text
LANGUAGE pltcl AS $$
  return "Hello, $1!"
$$;

SELECT tcl_hello('world');

-- 包含条件逻辑的函数
CREATE FUNCTION tcl_max(a integer, b integer) RETURNS integer
LANGUAGE pltcl AS $$
  if {$1 > $2} {return $1}
  return $2
$$;

-- 集合返回函数
CREATE FUNCTION tcl_sequence(count integer) RETURNS SETOF integer
LANGUAGE pltcl AS $$
  for {set i 1} {$i <= $1} {incr i} {
    return_next $i
  }
$$;

-- 触发器函数
CREATE FUNCTION tcl_audit() RETURNS trigger
LANGUAGE pltcl AS $$
  set NEW(modified_at) [clock format [clock seconds] -format "%Y-%m-%d %H:%M:%S"]
  return [array get NEW]
$$;
```

在 PL/Tcl 中通过 `spi_exec` 命令访问数据库：

```sql
CREATE FUNCTION tcl_count_rows(tbl text) RETURNS integer
LANGUAGE pltcl AS $$
  spi_exec "SELECT count(*) AS cnt FROM $1"
  return $cnt
$$;
```
