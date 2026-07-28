## 用法

来源：

- [官方上游 README](https://github.com/openuruunofficial/moss/blob/ace66c80de4e3388b94db8f0f95b5a615bc92027/postgresql/README)
- [官方扩展控制文件 (moss_uuidgen.control)](https://github.com/openuruunofficial/moss/blob/ace66c80de4e3388b94db8f0f95b5a615bc92027/postgresql/moss_uuidgen.control)
- [官方扩展 SQL (moss_uuidgen--1.0.sql)](https://github.com/openuruunofficial/moss/blob/ace66c80de4e3388b94db8f0f95b5a615bc92027/postgresql/moss_uuidgen--1.0.sql)

`moss_uuidgen` — 如果您使用的是 >= 9.1，您需要创建一个扩展：psql -U postgres moss -c "CREATE EXTENSION moss_uuidgen"。在需要这些特殊函数或聚合时使用它。使用上方链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION moss_uuidgen;
```

在目标数据库中安装扩展，在可用时运行上方的最小上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `uuid()` 是一个扩展函数并返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源代码进行比对。
