## 用法

来源：

- [上游 README](https://github.com/patchsoft/postcode/blob/13ccded329c2e12081987d197cfe44413c35bd4e/README.md)
- [扩展 control 文件](https://github.com/patchsoft/postcode/blob/13ccded329c2e12081987d197cfe44413c35bd4e/postcode.control)
- [SQL 安装脚本](https://github.com/patchsoft/postcode/blob/13ccded329c2e12081987d197cfe44413c35bd4e/postcode--1.3.0.sql)
- [PGXN 发布页面](https://pgxn.org/dist/postcode/)

`postcode` `1.3.0` 版提供紧凑的 `postcode` 与 `dps` 类型，分别表示英国邮编和投递点后缀。它会规范化输入，提供 B-tree 比较支持，用 `%` 执行部分匹配，并提供验证与格式化辅助函数。

### 示例

```sql
CREATE EXTENSION postcode;
CREATE TABLE addresses (code postcode);
INSERT INTO addresses VALUES ('SW1A 1AA'), ('LS2 4AA');
SELECT postcode_validate('SW1A 1AA');
SELECT * FROM addresses WHERE code % 'SW1A';
SELECT to_char(code, 'AD SW') FROM addresses;
```

语法验证不能证明邮编当前已分配或地址真实存在；此类结论必须用当前权威地址数据集复核。部分形式可能存在歧义（例如邮区与扇区），因此应用代码应明确允许的输入层级。上游版本发布于 2015 年，没有当前兼容矩阵，内置的英国邮编格式假设也可能已经过时。
