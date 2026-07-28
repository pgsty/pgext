## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_xnode/pg_xnode-0.7.2/README)
- [官方扩展控制文件 (xnode.control)](https://api.pgxn.org/src/pg_xnode/pg_xnode-0.7.2/src/xnode.control)
- [官方扩展 SQL (xnode.sql)](https://api.pgxn.org/src/pg_xnode/pg_xnode-0.7.2/src/sql/xnode.sql)

`xnode` — 使用 DOM 实现 XML。当应用程序数据需要此类型、域或其操作符时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION xnode;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add(doc, path, node, add_mode)` 是一个扩展函数，返回 `doc`。
- `children(node)` 是一个扩展函数，返回 `node[]`。
- `doc_in(cstring)` 是一个扩展函数，返回 `doc`。
- `doc_out(doc)` 是一个扩展函数，返回 `cstring`。
- `doc_to_node(doc)` 是一个扩展函数，返回 `node`。
- `element(text, text[][2], node)` 是一个扩展函数，返回 `node`。
- `fragment_sfunc(node, node)` 是一个扩展函数，返回 `node`。
- `node(xnt, text[], record)` 是一个扩展函数，返回 `node`。
- `node_debug_print(node)` 是一个扩展函数，返回 `text`。
- `node_in(cstring)` 是一个扩展函数，返回 `node`。
- `node_kind(node)` 是一个扩展函数，返回 `text`。
- `node_out(node)` 是一个扩展函数，返回 `cstring`。
- `node_to_doc(node)` 是一个扩展函数，返回 `doc`。
- `path(path, doc)` 是一个扩展函数，返回 `pathval`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.7.2`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
