## 用法

来源：

- [官方扩展控制文件](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/ext/gitgres.control)
- [官方上游文档](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/README.md)

`gitgres` — 在 PostgreSQL 表中存储 Git 对象与引用，并提供 Git 数据查询函数和原生 git_oid 类型。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pgcrypto`。

```sql
CREATE EXTENSION "gitgres";
SELECT extversion
FROM pg_extension
WHERE extname = 'gitgres';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
