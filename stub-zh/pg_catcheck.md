

## 用法

> [pg_catcheck: 诊断系统目录损坏](https://github.com/EnterpriseDB/pg_catcheck)

`pg_catcheck` 是一个命令行工具，通过验证目录表之间的交叉引用来检查 PostgreSQL 系统目录是否损坏。它接受与其他 PostgreSQL 工具相同的连接参数（`-h`、`-p`、`-U`、`-d`）。

### 基本用法

```bash
pg_catcheck -h localhost -p 5432 -d mydb
```

未发现问题时的正常输出：

```
progress: done (0 inconsistencies, 0 warnings, 0 errors)
```

### 存在损坏时的示例输出

```
notice: pg_class row has invalid relnamespace "24580": no matching entry in pg_namespace
row identity: oid="24581" relname="foo" relkind="r"
notice: pg_type row has invalid typnamespace "24580": no matching entry in pg_namespace
row identity: oid="24583"
progress: done (4 inconsistencies, 0 warnings, 0 errors)
```

### 结果分类

- **不一致**：目录交叉引用中的逻辑问题（例如悬空的 OID 引用）
- **警告**：更严重的问题
- **错误**：无法读取目录

### 选项

```bash
pg_catcheck --help                      # 完整选项列表
pg_catcheck --select-from-relations     # 同时检查缺失/不可访问的关系文件
```

### 连接

支持与 `psql` 相同的选项：`-h` 主机、`-p` 端口、`-U` 用户、`-d` 数据库，或连接字符串/URL。
