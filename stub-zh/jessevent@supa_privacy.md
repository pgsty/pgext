## 用法

来源：

- [Official database.dev 包页面](https://database.dev/jessevent/supa_privacy)

`jessevent@supa_privacy` — 格式保留匿名化和数据屏蔽。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "jessevent@supa_privacy";
```

在目标数据库中安装扩展，当可用时运行上方的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `create_masked_view` 是一个扩展函数。
- `generalize_date` 是一个扩展函数。
- `generalize_numeric` 是一个扩展函数。
- `mask_email` 是一个扩展函数。
- `mask_phone` 是一个扩展函数。
- `mask_phone_flexible` 是一个扩展函数。
- `mask_text` 是一个扩展函数。
- `partial_mask` 是一个扩展函数。
- `perturb_numeric` 是一个扩展函数。
- `perturb_numeric_deterministic` 是一个扩展函数。
- `salted_hash` 是一个扩展函数。
- `shift_date_deterministic` 是一个扩展函数。

### 要求与注意事项

- 该扩展的版本记录在目录中，名称为 `1.0.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行对比。
