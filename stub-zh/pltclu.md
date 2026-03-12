

## 用法

> [pltclu: PL/Tcl 不受信过程语言](https://www.postgresql.org/docs/current/pltcl.html)

PL/Tcl 不受信版本提供完整的 Tcl 功能，包括文件系统访问和外部程序执行。只有超级用户可以使用此语言创建函数。

```sql
CREATE EXTENSION pltclu;

-- 从服务器文件系统读取文件
CREATE FUNCTION read_file(filename text) RETURNS text
LANGUAGE pltclu AS $$
  set fd [open $1 r]
  set content [read $fd]
  close $fd
  return $content
$$;

-- 执行外部命令
CREATE FUNCTION run_command(cmd text) RETURNS text
LANGUAGE pltclu AS $$
  return [exec {*}$1]
$$;

-- 访问环境变量
CREATE FUNCTION get_env(varname text) RETURNS text
LANGUAGE pltclu AS $$
  if {[info exists ::env($1)]} {
    return $::env($1)
  }
  return ""
$$;

SELECT get_env('HOME');
```
