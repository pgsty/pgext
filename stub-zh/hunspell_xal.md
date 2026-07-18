## 用法

来源：

- [官方扩展控制文件](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/hunspell_xal.control)
- [官方上游文档](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/README.md)

`hunspell_xal` — 用于 PostgreSQL 全文搜索的卡尔梅克语 Hunspell 词典。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "hunspell_xal";
SELECT extversion
FROM pg_extension
WHERE extname = 'hunspell_xal';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
