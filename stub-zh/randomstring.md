## 用法

来源：

- [官方扩展 SQL](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring--1.0.0.sql)
- [官方 C 实现](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring.c)
- [官方扩展控制文件](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring.control)

`randomstring` 版本 `1.0.0` 提供用于测试数据的简单伪随机文本与字节生成器。它使用 C 库的 `random()` 生成器，因此两个函数都不适合密码、令牌、密钥、盐值或其他安全敏感材料。

### 核心流程

两个函数都要求输出长度严格为正数：

```sql
CREATE EXTENSION randomstring;

SELECT random_string(24);
SELECT encode(random_bytea(16), 'hex');
```

`random_string(integer)` 返回文本，`random_bytea(integer)` 返回二进制数据。长度为零或负数时会报错。

### 分布与安全注意事项

已核验实现只索引所声明文本字母表的前 62 个字节。实际文本输出可能包含空格、大小写字母和数字 0 到 8；数字 9 与列出的标点都不会出现。二进制生成使用模 255，因此永远不会产生字节值 255。

扩展没有提供种子管理或唯一性保证。值可能重复，分布也没有覆盖所宣称的完整字母表。只有在这些偏差可接受时，才能用于非敏感测试夹具；安全材料应使用 PostgreSQL 密码学功能或应用侧 CSPRNG。
