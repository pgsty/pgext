## 用法

来源：

- [官方上游文档](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/readme_jp.txt)

`vmatch` — vmatch：文本模糊匹配扩展，提供相似度函数与 /= 操作符。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "vmatch";
SELECT extversion
FROM pg_extension
WHERE extname = 'vmatch';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
