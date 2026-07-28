## 用法

来源：

- [官方上游 README](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/README)
- [官方扩展控制文件 (gms_tcp.control)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_tcp/gms_tcp.control)
- [官方扩展 SQL (gms_tcp--1.0--1.1.sql)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_tcp/gms_tcp--1.0--1.1.sql)

`gms_tcp` — 提供了 PL/SQL 中的 TCP/IP 客户端访问功能。在移植或模拟相应的数据库 API 时使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION gms_tcp;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gms_tcp.available(c in gms_tcp.connection, timeout in int default 0)` 是一个扩展函数，返回 `integer`。
- `gms_tcp.available_real(c in gms_tcp.connection, timeout in int)` 是一个扩展函数，返回 `integer`。
- `gms_tcp.close_all_connections()` 是一个扩展函数，返回 `void`。
- `gms_tcp.close_connection(c in gms_tcp.connection)` 是一个扩展函数，返回 `void`。
- `gms_tcp.connection_in(cstring)` 是一个扩展函数，返回 `gms_tcp`。
- `gms_tcp.connection_out(gms_tcp.connection)` 是一个扩展函数，返回 `cstring`。
- `gms_tcp.crlf()` 是一个扩展函数，返回 `varchar2`。
- `gms_tcp.flush(c in gms_tcp.connection)` 是一个扩展函数，返回 `void`。
- `gms_tcp.get_line(c in gms_tcp.connection, remove_crlf in boolean default false, peek in boolean default false)` 是一个扩展函数，返回 `text`。
- `gms_tcp.get_line_real(c in gms_tcp.connection, remove_crlf in boolean, peek in boolean, ch_charset in boolean default false)` 是一个扩展函数，返回 `text`。
- `gms_tcp.get_raw(c in gms_tcp.connection, len in integer default 1, peek in boolean default false)` 是一个扩展函数，返回 `raw`。
- `gms_tcp.get_raw_real(c in gms_tcp.connection, len in integer, peek in boolean)` 是一个扩展函数，返回 `raw`。
- `gms_tcp.get_text(c in gms_tcp.connection, len in integer default 1, peek in boolean default false)` 是一个扩展函数，返回 `text`。
- `gms_tcp.get_text_real(c in gms_tcp.connection, len in integer, peek in boolean, ch_charset in boolean default false)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.1`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
