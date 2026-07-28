## 用法

来源：

- [官方上游 README](https://github.com/kurtbuilds/sid/blob/cfd88358d24675b49cf6907e42397f06264a3809/README.md)
- [官方扩展控制文件 (sid.control)](https://github.com/kurtbuilds/sid/blob/cfd88358d24675b49cf6907e42397f06264a3809/pg/sid.control)
- [官方实现源代码](https://github.com/kurtbuilds/sid/blob/cfd88358d24675b49cf6907e42397f06264a3809/pg/src/lib.rs)

`sid` — 可排序带标签的 128 位标识符类型，具有紧凑的人类可读编码。当应用程序数据需要此类型、域或其操作符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION sid;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `sid_from_uuid` 是一个扩展函数。
- `sid_new` 是一个扩展函数。
- `sid_null` 是一个扩展函数。

### 要求与注意事项

- 元组记录版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
