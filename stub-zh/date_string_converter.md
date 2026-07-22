## 用法

来源：

- [项目 README](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/README.md)
- [date_string_converter 实现](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/src/lib.rs)
- [扩展控制文件](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/date_string_converter.control)

`date_string_converter` 0.0.0 是一个 pgrx 演示扩展。虽然仓库级 README 描述了将 `1w 2d 4h 30m` 之类字符串转换为时长的目标，但固定版本的实际扩展只暴露问候函数，并不包含日期时长解析器或转换 API。

### 可用流程

```sql
CREATE EXTENSION date_string_converter;

SELECT hello_date_string_converter();
-- Hello, date_string_converter
```

此版本唯一面向用户的对象是 `hello_date_string_converter() RETURNS text`。扩展不可重定位，并且控制文件要求由超级用户安装。

### 限制

不要依据仓库级说明设计转换逻辑：此版本不存在接收 `1w 2d 4h 30m` 的函数，也无法生成 `220.5` 这样的结果。在上游实现并记录真实转换函数之前，应把它视作最小扩展脚手架。其 Cargo 配置通过 pgrx 0.10.1 面向 PostgreSQL 11 至 16；该源码快照未确认对更新服务器版本的兼容性。
