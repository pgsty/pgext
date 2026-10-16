## 用法

来源：

- [README.md](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/README.md)
- [postcode.control](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/postcode.control)
- [postcode--2.1.0.sql](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/postcode--2.1.0.sql)
- [META.json](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/META.json)
- [.github/workflows/test.yml](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/.github/workflows/test.yml)
- [CHANGELOG.md](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/CHANGELOG.md)

`postcode` 2.1.0 是原 patchsoft 项目在 PGXN 上由维护分叉延续的版本。它保留紧凑的英国 `postcode` 及投递点后缀 `dps` 类型，并新增 64 位国际 `postal_code` 类型。带国家代码的文本形式可以区分含义不同的各国邮编。

### 核心用法

```sql
CREATE EXTENSION postcode;
SELECT 'SW1A 1AA'::postcode;
SELECT 'US-90210-1234'::postal_code, 'CA-K1A 0B1'::postal_code;
CREATE TABLE addresses (code postcode);
INSERT INTO addresses VALUES ('SW1A 1AA'), ('LS2 4AA');
SELECT * FROM addresses
WHERE code >= range_lower('SW1A') AND code < range_upper('SW1A');
```

### 运行边界

`postcode` 支持格式化、校验及 B-tree 索引。`%` 前缀过滤在常量片段上可使用规划器支持；`range_lower` 与 `range_upper` 可明确表达索引范围。早期 1.3.0 错误地将前缀匹配注册为 B-tree 相等关系，维护版本修复了这个问题及一个堆溢出缺陷。替换库后应执行 `ALTER EXTENSION postcode UPDATE`；仅替换二进制不会修复 SQL 运算符元数据。

国际 `postal_code` 值必须包含国家代码，形式见上方示例。`country` 提取国家代码，`postal_prefix`、`lower_bound` 与 `upper_bound` 支持前缀操作。`is_valid_postal_code` 和 `to_postal_code` 用于导入不洁净输入，也可用国家类型修饰符约束列。格式合法不代表该邮编实际分配或地址真实存在。

该发行版测试 PostgreSQL 14–18。安装需要超级用户，扩展可迁移模式，不要求预加载。2.0 以追加方式升级，保留原有英国类型与已存储值。

2.1.0 为两种邮编类型增加 `part`、`parts` 和 `prefix_of`。字符串字面量须显式转换类型，避免重载歧义；`postal_code_parts` 列出可用名称。`to_char(postcode, text)` 已弃用，但会保留到 2.x 系列结束。该版本还修复了编译大型自定义模式时的释放后使用缺陷，内置模式不受影响。安装新版库后执行 `ALTER EXTENSION postcode UPDATE`；2.1 升级保留既有存储值。
