## 用法

来源：

- [已复核 commit 的 pg_median_utils 文档](https://github.com/greenape/pg_median_utils/blob/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/doc/pg_median_utils.md)
- [已复核 commit 的 pg_median_utils.control](https://github.com/greenape/pg_median_utils/blob/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/pg_median_utils.control)
- [已安装 SQL API](https://github.com/greenape/pg_median_utils/blob/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/sql/pg_median_utils--0.0.1.sql)
- [窗口函数实现](https://github.com/greenape/pg_median_utils/tree/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/src)
- [当前 PGXN 发行页面](https://pgxn.org/dist/pg_median_utils/)

`pg_median_utils` 提供五个 `double precision` 窗口函数：`median_filter`、`iterated_median_filter`、`rolling_median`、`backfilled_rolling_median` 和 `rolling_median_impute`。它们处理每个窗口分区中的有序行，并要求窗口大小为奇数。

### 滚动中位数

```sql
CREATE EXTENSION pg_median_utils;

WITH samples(sample_no, value) AS (
  VALUES (1, 1.0), (2, 5.0), (3, 2.0), (4, 8.0), (5, 3.0)
)
SELECT sample_no, value,
       rolling_median(value, 3) OVER (ORDER BY sample_no) AS median_3
FROM samples
ORDER BY sample_no;
```

前两行返回 null，因为此时还没有完整的三行尾随窗口。

### 注意事项

- 目录、控制文件和实际扩展安装脚本报告 `0.0.1`，而文档、tag、元数据和 PGXN 把发行包标为 `0.0.7`。除非存在匹配的 SQL 升级路径，否则不要把发行标签用作 `ALTER EXTENSION` 目标。
- 排序来自 `OVER` 子句；这些函数检查分区位置，而不是调用者定义的 SQL window frame。务必指定确定性顺序。
- 窗口大小必须为正奇数。C 代码会拒绝偶数，但会在校验前把有符号 SQL 参数转换为无符号整数，因此负奇数可能变成巨大的内存分配量或循环上限。
- 迭代过滤器会反复处理整个分区直至收敛，插补函数则构建分区大小的数组。应限制分区和窗口规模，并测试后端内存与延迟。
- 上游声称支持 PostgreSQL 9.6 及以上版本，但 C 窗口 API 代码没有当前兼容性矩阵。
