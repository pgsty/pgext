## 用法

来源：

- [官方扩展控制文件](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/jalali_utils.control)

`jalali_utils` — 将公历日期和时间转换为 Jalali 历值，并提取 Jalali 日期字段。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "jalali_utils";
SELECT extversion
FROM pg_extension
WHERE extname = 'jalali_utils';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
