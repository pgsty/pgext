## 用法

来源：

- [官方上游 README](https://gitlab.com/RekGRpth/postgresql-functions/-/blob/master/README.md)
- [官方扩展控制文件](https://gitlab.com/RekGRpth/postgresql-functions/-/blob/master/tsearch/dict_maxlen/dict_maxlen.control)
- [官方扩展 SQL](https://gitlab.com/RekGRpth/postgresql-functions/-/blob/master/tsearch/dict_maxlen/dict_maxlen--1.0.sql)

`dict_maxlen` — 一个自定义 PostgreSQL 函数和扩展的仓库。使用它来实现相应的文本搜索、解析或语言工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION dict_maxlen;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `dictmaxlen_init(internal)` 是一个扩展函数，返回 `internal`。
- `dictmaxlen_lexize(internal, internal, internal, internal)` 是一个扩展函数，返回 `internal`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
