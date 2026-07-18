## 用法

来源：

- [上游 README](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/README.md)
- [扩展 control 文件](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/pgfaker.control)
- [Rust 包清单](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/Cargo.toml)
- [SQL 函数实现](https://github.com/rustprooflabs/pgfaker/blob/034a697c4270c30caaab81da5c35a1005e913c62/src/lib.rs)

`pgfaker` `0.0.3` 版在固定的 `pgfaker` 模式中安装美式英语假数据生成器。它能生成公司名称与标语、姓名与前后缀、电话号码、域名、电子邮件地址和用户名。当前包面向 PostgreSQL 14–18。

### 生成可丢弃的样例行

```sql
CREATE EXTENSION pgfaker;

SELECT pgfaker.company(),
       pgfaker.email(),
       pgfaker.phone();

SELECT g AS sample_id,
       pgfaker.person_full_name()
FROM generate_series(1, 5) AS g;
```

结果是随机的，而且扩展没有提供种子或可复现接口。生成值不保证唯一、有效、可投递、具有文化代表性、隐私安全，也不保证符合应用模式中的约束。它们只应用于可丢弃的演示和测试，并配合明确的唯一性/有效性处理；绝不能把它们用于密码、API 令牌、身份凭证或生产数据匿名化。除非操作经过隔离且可逆，否则不要在含有真实记录的表中批量运行生成器。
