## 用法

来源：

- [官方扩展控制文件](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/advcopy.control)
- [官方上游文档](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/README.md)

`advcopy` — 用于在 PostgreSQL 与 S3 之间导入导出数据的 PL/Python 扩展。

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`。

```sql
CREATE EXTENSION "advcopy";
SELECT extversion
FROM pg_extension
WHERE extname = 'advcopy';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
