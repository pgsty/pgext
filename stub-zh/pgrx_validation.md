## 用法

来源：

- [官方控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/pgrx_validation/pgrx_validation.control)
- [官方 PGRX Validation README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/pgrx_validation/README.md)
- [生成的 PostgreSQL 17 扩展 SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/pgrx_validation/extension/usr/share/postgresql/17/extension/pgrx_validation--0.0.0.sql)

`pgrx_validation` 将 EMI 项目共享的校验规则暴露为 PostgreSQL 函数。它们用于数据库约束，让字符串、数值、UUID、时间戳和图标检查与其他应用目标保持一致。

### 核心流程

```sql
CREATE EXTENSION pgrx_validation;

CREATE TABLE contacts (
    email text NOT NULL CHECK (must_be_email(email)),
    score double precision NOT NULL CHECK (must_be_strictly_positive_f64(score))
);
```

生成的函数都是 `STRICT`，因此空输入得到空结果。PostgreSQL 的 `CHECK` 接受空值；必须拒绝缺失值时，应另加 `NOT NULL`。

### 校验函数

- 字符串：`must_be_email(text)`、`must_not_be_empty(text)`、`must_not_be_padded(text)`、`must_not_contain_consecutive_whitespace(text)`、`must_not_contain_control_characters(text)`、`must_be_paragraph(text)` 和 `must_be_distinct(text, text)`。
- 整数与 UUID：`must_be_positive_i16(smallint)`、`must_be_positive_i32(int)`、`must_be_strictly_positive_i16(smallint)`、`must_be_strictly_positive_i32(int)`、`must_be_distinct_i16(smallint, smallint)`、`must_be_distinct_i32(int, int)` 和 `must_be_distinct_uuid(uuid, uuid)`。
- 浮点数：`must_be_strictly_positive_f32(real)`、`must_be_strictly_positive_f64(double precision)`，以及边界检查 `must_be_greater_than_f32`、`must_be_smaller_than_f32`、`must_be_strictly_greater_than_f32`、`must_be_strictly_smaller_than_f32`、`must_be_strictly_greater_than_f64` 和 `must_be_strictly_smaller_than_f64`。
- 其他检查：`must_be_smaller_than_utc(timestamptz, timestamptz)`、`must_be_strictly_smaller_than_utc(timestamptz, timestamptz)` 和 `must_be_font_awesome_class(text)`。

### 运维说明

所核验包版本为 0.0.0，API 仍可能变化。扩展不可重定位；控制文件设置 `superuser = true` 和 `trusted = false`，因此必须由超级用户创建。它未声明前置扩展或预加载要求。应把这些谓词视为应用数据校验而不是安全净化，并在共享规则升级时回归测试约束。
