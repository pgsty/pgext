## 用法

来源：

- [Official README](https://github.com/ccutrer/pg_collkey/blob/42d93fd5182ca94efbc45d7c02b86d95e8e2070a/README)
- [Extension control file](https://github.com/ccutrer/pg_collkey/blob/42d93fd5182ca94efbc45d7c02b86d95e8e2070a/pg_collkey.control)
- [SQL function definitions](https://github.com/ccutrer/pg_collkey/blob/42d93fd5182ca94efbc45d7c02b86d95e8e2070a/collkey_icu.sql)

pg_collkey 封装 ICU 排序键生成，让查询可以独立于数据库默认排序规则选择区域设置和比较强度。它适合生成可排序的字节键，包括忽略重音或识别数字的排序，但这份年代较久的上游代码只支持 UTF-8 数据库。

### 核心流程

可以在 `ORDER BY` 中生成排序键；如果区域设置及选项保持稳定，也可将其放入表达式索引：

```sql
CREATE EXTENSION pg_collkey;

SELECT name
FROM product
ORDER BY collkey(name, 'fr_FR');

CREATE INDEX product_name_root_idx
ON product (collkey(name, 'root'));
```

双参数重载会采用后置标点处理、默认强度以及数字排序。单参数重载还会选择 ICU root 排序规则。

### 重要对象

- `collkey(text, text, boolean, integer, boolean)` 以 `bytea` 返回 ICU 排序键。
- locale 参数选择 ICU root 或 `fr_FR` 等命名区域设置。
- 第一个 boolean 会把标点移到第四比较层级。
- 强度 1 比较基本字符；2 包含重音；3 包含大小写；4 包含其他差异；5 比较规范化后的码点。
- 最后一个 boolean 启用数字排序，使数字序列按数值排列。
- `collkey(text, text)` 和 `collkey(text)` 是便捷重载。

### 运维说明

排序键来自链接的 ICU 库。ICU 升级或区域数据变化可能使既有排序假设和表达式索引失效，因此升级后要重建索引并回归测试结果。在同一查询中使用许多区域设置会增加初始化开销。README 记录的是非常老的 PostgreSQL 与 ICU 基线，且不支持非 UTF-8 数据库。
