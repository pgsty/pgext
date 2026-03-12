

## 用法

> [pg_polyline: PostgreSQL 的 Google 编码折线编解码](https://github.com/yihong0618/pg_polyline)

快速的 Google 编码折线编码与解码 PostgreSQL 扩展。基于 `pgrx` 构建。

```sql
CREATE EXTENSION pg_polyline;

-- 将点数组编码为折线字符串
SELECT polyline_encode(
  ARRAY[point(-120.2, 38.5), point(-120.95, 40.7), point(-126.453, 43.252)], 6
);
--          polyline_encode
-- ----------------------------------
--  _izlhA~rlgdF_{geC~ywl@_kwzCn`{nI

-- 将折线字符串解码回点数组
SELECT polyline_decode('_ibE_seK_seK_seK', 6);
--       polyline_decode
-- ---------------------------
--  {"(0.2,0.1)","(0.4,0.3)"}
```

第二个参数是精度（小数位数）。
