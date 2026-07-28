## 用法

来源：

- [官方上游 README](https://github.com/karajan1001/dummy_index_am/blob/5468e7325bd930af02152b0bc0c6f99eae5e64af/README.md)
- [官方扩展控制文件 (dummy_index_am.control)](https://github.com/karajan1001/dummy_index_am/blob/5468e7325bd930af02152b0bc0c6f99eae5e64af/dummy_index_am.control)
- [官方实现源代码](https://github.com/karajan1001/dummy_index_am/blob/5468e7325bd930af02152b0bc0c6f99eae5e64af/src/lib.rs)

`dummy_index_am` — 假定索引 AM 是一个用于测试任何可用于索引访问方法的设施的模块，其代码保持尽可能简单。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION dummy_index_am;
```

在目标数据库中安装扩展，如果有可用示例，则运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 该目录记录了版本信息 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
