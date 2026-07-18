## 用法

来源：

- [官方扩展控制文件](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx.control)

`lsm_idx` — 基于 PostgreSQL B-tree 内部接口实现的实验性 LSM 风格索引访问方法。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "lsm_idx";
SELECT extversion
FROM pg_extension
WHERE extname = 'lsm_idx';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
