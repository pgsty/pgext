## 用法

来源：

- [上游 README](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/README.md)
- [扩展 control 文件](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/pgspeck.control)
- [1.0 版 SQL 接口](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/pgspeck--1.0.sql)
- [PGXN 包元数据](https://github.com/johto/pgspeck/blob/6120dffe2d293ed1a3e201ca299fba348a7b381d/META.json)

`pgspeck` `1.0` 版提供确定性的 Speck 分组密码置换，可把连续整数值变成不那么直观的整数序列。32 位函数使用明文的低 32 位和一个 64 位密钥；48 位函数使用明文的低 48 位以及两个非零的 48 位密钥分量。

### 验证可逆映射

```sql
CREATE EXTENSION pgspeck;

SELECT pgspeck_encrypt32(42, 123456789);

SELECT pgspeck_decrypt32(
    pgspeck_encrypt32(42, 123456789),
    123456789
);
```

它只能用于可逆的标识符混淆。小型确定性分组会暴露重复输入，不提供随机数或认证，也不保护完整性；回绕或截断还需要谨慎分析取值域。SQL 中传入的密钥可能通过函数定义、日志、监控、查询文本或角色访问而泄漏。不要把 `pgspeck` 用于密码、个人数据、机密、访问令牌或现代应用加密。授权必须独立于混淆后的标识符；只要涉及保密性，就应采用经过审查的认证加密设计。
