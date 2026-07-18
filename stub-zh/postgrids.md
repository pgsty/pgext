## 用法

来源：

- [官方扩展控制文件](https://github.com/BIS-Brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/postgrids.control)
- [官方上游文档](https://github.com/BIS-Brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/README.md)
- [官方 Rust 包清单](https://github.com/BIS-Brecon/postgrids/blob/ba0a85a344cfefc8e630c7428e930e6eb8df03aa/Cargo.toml)

`postgrids` — 基于 pgrx 的英国 OSGB 与爱尔兰 OSI 国家格网引用扩展，支持解析、格式化及精度重算。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postgrids";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgrids';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
