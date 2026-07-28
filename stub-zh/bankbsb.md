## 用法

来源：

- [官方上游 README](https://github.com/timwmillard/bankbsb/blob/fa62a91aa3c3845527e29449fd2feb328427f5c4/README.md)
- [官方扩展控制文件 (bankbsb.control)](https://github.com/timwmillard/bankbsb/blob/fa62a91aa3c3845527e29449fd2feb328427f5c4/bankbsb.control)
- [官方扩展 SQL (bankbsb--0.0.1.sql)](https://github.com/timwmillard/bankbsb/blob/fa62a91aa3c3845527e29449fd2feb328427f5c4/bankbsb--0.0.1.sql)

`bankbsb` — 一个用于添加澳大利亚银行 BSB 数字类型的 PostgreSQL 扩展。当应用程序数据需要此类型、域或其操作符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION bankbsb;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `bankbsb_cmp(bsb, bsb)` 是一个扩展函数，返回 `integer`。
- `bankbsb_eq(bsb, bsb)` 是一个扩展函数，返回 `boolean`。
- `bankbsb_ge(bsb, bsb)` 是一个扩展函数，返回 `boolean`。
- `bankbsb_gt(bsb, bsb)` 是一个扩展函数，返回 `boolean`。
- `bankbsb_in(cstring)` 是一个扩展函数，返回 `bsb`。
- `bankbsb_le(bsb, bsb)` 是一个扩展函数，返回 `boolean`。
- `bankbsb_lt(bsb, bsb)` 是一个扩展函数，返回 `boolean`。
- `bankbsb_ne(bsb, bsb)` 是一个扩展函数，返回 `boolean`。
- `bankbsb_out(bsb)` 是一个扩展函数，返回 `cstring`。
- `bsb` 是一个扩展定义的类型。
- `bankbsb_ops` 是一个扩展定义的操作符类。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
