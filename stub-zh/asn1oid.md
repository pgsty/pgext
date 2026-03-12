

## 用法

> [asn1oid: PostgreSQL 的 ASN.1 OID 数据类型](https://github.com/df7cb/pgsql-asn1oid)

`asn1oid` 扩展提供了用于存储和比较 ASN.1 对象标识符（OID）的数据类型。

```sql
CREATE EXTENSION asn1oid;

SELECT '1.3.6.1.4.1'::asn1oid;
   asn1oid
─────────────
 1.3.6.1.4.1
```

### 数据类型

`asn1oid` 类型以点分十进制表示法存储 ASN.1 OID 值（例如 `1.3.6.1.4.1.311`）。这些是 SNMP、LDAP、X.509 证书和其他标准中使用的层级标识符。

### 运算符

支持标准比较运算符用于排序和相等判断：`=`、`<>`、`<`、`>`、`<=`、`>=`。

### 类型转换

该类型支持与 `text` 类型之间的相互转换。
