## 用法

来源：

- [固定提交的 1.0 版安装 SQL](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/nvlfunc--1.0.sql)
- [固定提交的扩展 control 文件](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/nvlfunc.control)
- [固定提交的上游 README](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/README.md)

`nvlfunc` `1.0` 提供一个 Oracle 风格的 `NVL` 便利函数。它实际的 SQL 接口刻意保持很窄：`public.NVL(smallint, smallint)` 通过 PostgreSQL `COALESCE` 返回第一个非空参数。

### 核心流程

```sql
CREATE EXTENSION nvlfunc;

SELECT public.NVL(NULL::smallint, 7::smallint);  -- 7
SELECT public.NVL(3::smallint, 7::smallint);     -- 3
SELECT public.NVL(NULL::smallint, NULL::smallint);
```

### 重要对象

- `public.NVL(smallint, smallint) RETURNS smallint`：封装 `COALESCE($1, $2)` 的 `IMMUTABLE` SQL 函数。PostgreSQL 会把未加引号的声明折叠为函数名 `nvl`。

### 运维说明

该扩展不实现 Oracle `NVL` 的广泛类型转换行为，也没有面向 `integer`、`bigint`、numeric、text、date 或 timestamp 的重载。当字面量或 `NULL` 导致重载解析不明确时，应把两个参数都显式转换为 `smallint`；其他类型直接使用内置 `COALESCE`。函数被明确安装在 `public`，扩展不可重定位，因此安装前应检查名称冲突与 search-path 策略。上游仓库只包含这一个函数和极少文档；应在目标 PostgreSQL 版本上验证兼容性，不要据此推断存在更广泛的 Oracle 兼容层。
