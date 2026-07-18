## 用法

来源：

- [已复核 commit 的 phone_number README](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/README.md)
- [已复核 commit 的 phone_number Cargo 配置](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/Cargo.toml)
- [已复核 commit 的 phone_number SQL 类型与函数](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/mod.rs)
- [已复核 commit 的 phone_number 输入输出实现](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/implementation.rs)

`phone_number` 0.0.0 是一个最小化 pgrx 原型，定义结构化电话号码值及一个 SQL 函数 `create_random_number`。

```sql
CREATE EXTENSION phone_number;

SELECT create_random_number();
```

输入实现只接受五组以空格分隔的数字，例如 `1 212 555 12 34`；输出格式化器则生成类似 `+1 212 555-12-34` 的值。

### 注意事项

- 0.0.0 版本及只有一行的 README 表明它仍是原型；control 文件要求超级用户。
- Cargo 只启用 PostgreSQL 16，没有其他 PostgreSQL 特性或兼容性声明。
- 输入输出格式无法往返：解析器会拒绝格式化器生成的加号与连字符。没有规范化层时，不要把文本输出持久化后再尝试转换回来。
- 随机生成的组成部分既不是经过验证的真实电话号码分配，也不是密码学安全标识符。
