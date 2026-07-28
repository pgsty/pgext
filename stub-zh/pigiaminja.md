## 用法

来源：

- [官方上游 README](https://github.com/leomos/pigiaminja/blob/a9b476b1405deb2926f57ac577ddf716a3f34b9b/README.md)
- [官方扩展控制文件 (pigiaminja.control)](https://github.com/leomos/pigiaminja/blob/a9b476b1405deb2926f57ac577ddf716a3f34b9b/pigiaminja.control)
- [官方实现源代码](https://github.com/leomos/pigiaminja/blob/a9b476b1405deb2926f57ac577ddf716a3f34b9b/src/lib.rs)

`pigiaminja` — 一个 PostgreSQL 扩展，为 COPY TO 命令添加了 Jinja 模板格式支持。在从 PostgreSQL 移动、转换或集成相应数据时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pigiaminja;
```

在目标数据库中安装扩展，如果有可用示例，请运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
