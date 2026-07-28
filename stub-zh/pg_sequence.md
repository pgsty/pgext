## 用法

来源：

- [官方上游 README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)
- [官方扩展控制文件 (pg_sequence.control)](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_sequence/pg_sequence.control)
- [官方实现源代码](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_sequence/src/lib.rs)

`pg_sequence` — ERP 文档编号使用格式化、限定和自增的序列。当 SQL 需要这些特殊功能或聚合时使用它。上游明确表示该项目尚未准备好用于生产环境。

### 核心工作流

```sql
CREATE EXTENSION pg_sequence;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了版本 `0.2.0`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游明确表示该项目尚未准备好用于生产环境。
- 上游将该项目描述为概念验证项目。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
