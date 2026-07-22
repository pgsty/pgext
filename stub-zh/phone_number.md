## 用法

来源：

- [Official README](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/README.md)
- [Cargo version and PostgreSQL feature](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/Cargo.toml)
- [SQL-facing type and generator](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/mod.rs)
- [Input and output implementation](https://github.com/Igr1x/phone_number/blob/5e3347ae90f0cc3b96c142af73523c9d52b78b05/src/phone_number/implementation.rs)

`phone_number` 0.0.0 是一个最小化 pgrx 原型，定义五段式 `PhoneNumber` 类型和 `create_random_number()`。它没有实现 E.164 解析、地区校验、格式策略、运营商元数据或真实号码分配。

### 原型接口

输入解析器只接受五组以空格分隔的数字，长度依次为 1、3、3、2、2：

```sql
CREATE EXTENSION phone_number;

SELECT '1 212 555 12 34'::phone_number;
SELECT create_random_number();
```

输出格式类似 `+1 212 555-12-34`。随机函数从伪随机范围填充每个数字分量，不会确认国家、城市、交换局或用户号码是否真实分配或可拨通。

### 严重的类型输入输出缺陷

0.0.0 的输出不能被自己的输入解析器接受：解析器要求空格，并会拒绝输出中的 `+` 和连字符。PostgreSQL 类型通常必须能通过文本往返，因此这个缺陷可能破坏类型转换、COPY、逻辑交换以及已存值的备份恢复。在修复输入/输出契约并通过备份恢复回归测试前，不要持久化该类型。

### 其他边界

Cargo 只启用 PostgreSQL 16，control 文件要求超级用户，上游 README 也只有一行。解析器只检查分组长度和数字语法，不检查有意义的编号规划约束；全零分量等值可以通过。生成值使用普通伪随机数生成器，既不全局唯一，也不具备密码学安全性。

该扩展只适合作为 pgrx 类型开发示例。生产电话号码数据应保留规范文本表示，并在受控校验层使用持续维护的编号规划库。
