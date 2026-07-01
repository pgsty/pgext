## 用法

来源：[README](https://github.com/fmbiete/pgdisablelogerror/blob/v1.0/README.md)、[v1.0 release](https://github.com/fmbiete/pgdisablelogerror/releases/tag/v1.0)、[control file](https://github.com/fmbiete/pgdisablelogerror/blob/v1.0/pgdisablelogerror.control)

`pgdisablelogerror` 会抑制配置的 SQLSTATE error codes 在 PostgreSQL server log 中的记录。对于 duplicate-key violations 等预期应用错误过于频繁、污染 server log 的场景，它比较有用。

### 启用 Hook

加载 module 并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pgdisablelogerror'
```

在 `postgres` 数据库中创建一次扩展：

```sql
CREATE EXTENSION pgdisablelogerror;
```

### 配置 SQLSTATE Codes

将 `pgdisablelogerror.sqlerrcode` 设为逗号分隔的 SQLSTATE codes：

```conf
pgdisablelogerror.sqlerrcode = '23505,23503'
```

空值或 NULL 会禁用 suppression：

```conf
pgdisablelogerror.sqlerrcode = ''
```

要在普通 PostgreSQL 日志中识别 SQLSTATE 值，可以在 `log_line_prefix` 中添加 `%e`。

### 注意事项

- 版本 1.0 支持 PostgreSQL 14-18。
- 该扩展影响日志，不影响错误行为。客户端仍会收到原始错误。
- SQLSTATE 列表应尽量窄。抑制过宽的错误类别可能隐藏真实运维问题。
