## 用法

来源：

- [官方上游 README](https://github.com/guepard-corp/gfs/blob/461d07820e253828d085a3b387f6912753e25d7c/crates/extensions/gfs/README.md)
- [官方扩展控制文件 (gfs.control)](https://github.com/guepard-corp/gfs/blob/461d07820e253828d085a3b387f6912753e25d7c/crates/extensions/gfs/gfs.control)
- [官方扩展 SQL (gfs--0.0.1.sql)](https://github.com/guepard-corp/gfs/blob/461d07820e253828d085a3b387f6912753e25d7c/crates/extensions/gfs/c-ref/gfs--0.0.1.sql)

`gfs` — 克隆一个远程 PostgreSQL 的 **copy-on-read**：一个空的本地数据库，仅在查询触及其数据时从源端获取数据，因此一个多 TB 的源可以瞬间克隆，并且克隆始终为 **部分状态**（它永远不会拉取应用程序未读取的数据）。在将数据从 PostgreSQL 移动、转换或集成时使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION gfs;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gfs.register_clone(local regclass, source_ref text, key_col text DEFAULT 'id')` 是一个扩展函数，返回 `void`。
- `gfs.unregister_clone(local regclass)` 是一个扩展函数，返回 `void`。
- `gfs.warm(local regclass)` 是一个扩展函数，返回 `bigint`。
- `gfs_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `gfs.clones` 是一个由扩展定义的视图。
- `gfs.clone_source` 是一个由扩展安装或管理的表。
- `gfs.clone_stats` 是一个由扩展安装或管理的表。
- `gfs` 是一个由扩展创建的模式。
- `gfs` 是一个由扩展定义的访问方法。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件标记该扩展为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
