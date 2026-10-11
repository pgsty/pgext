## 用法

来源：

- [docs/assert.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/assert.rst)
- [assert/assert--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/assert/assert--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [assert/assert.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/assert/assert.control)

`assert` 是 Oliphant 的 SQL 断言工具。检查成功时正常返回，失败时抛出异常，可使测试事务失败。

### 核心用法

```sql
CREATE EXTENSION "assert";
SELECT assert_equals(1, 1);
SELECT assert_is_not_null(1);
SELECT assert_raises('UTA08', 'SELECT assert_equals(1, 2)');
```

### 运行边界

这些是最初面向 PostgreSQL 9.x 编写文档的纯 SQL/PLpgSQL 工具，尚未验证当前主版本兼容性。安装沿用 control 默认权限，需要管理员；不含共享库，也不需要预加载。 `assert_equals` 和 `assert_not_equals` 不会因 NULL 参数而失败，需要时应结合 `assert_is_null` 或 `assert_is_not_null`。对象检查包括 `assert_table_exists`、`assert_column_exists`、`assert_function_exists` 和 `assert_trigger_exists`。`assert_raises` 会执行调用者提供的 SQL 并检查 SQLSTATE，应仅使用可信的测试语句。
