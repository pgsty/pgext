## 用法

来源：

- [README.md](https://api.pgxn.org/src/uint128/uint128-1.3.0/README.md)
- [uint128.control](https://api.pgxn.org/src/uint128/uint128-1.3.0/uint128.control)
- [uint128--1.2.0--1.3.0.sql](https://api.pgxn.org/src/uint128/uint128-1.3.0/uint128--1.2.0--1.3.0.sql)
- [1.3.0](https://github.com/pg-uint/pg-uint128/releases/tag/1.3.0)


`uint128` 提供无符号整数，以及有符号 8 位和 128 位整数类型。1.3.0 支持 PostgreSQL 12–18，并修复了可能使索引比较返回错误结果的操作符元数据。

### 保存与查询整数

安装需要超级用户；控制文件未将扩展标记为可信。扩展可迁移模式，不要求预加载。

```sql
CREATE EXTENSION uint128;
CREATE TABLE ids (id uint16 PRIMARY KEY);
INSERT INTO ids VALUES ('340282366920938463463374607431768211455'::uint16);
SELECT * FROM ids WHERE id > '18446744073709551615'::uint16;
SELECT * FROM generate_series(1::uint4, 3::uint4);
```

无符号类型 `uint1`、`uint2`、`uint4`、`uint8` 和 `uint16` 分别占用 1、2、4、8 和 16 字节，范围为零至 2 的相应位数次方减一。`int1` 和 `int16` 则是有符号的 1 字节和 16 字节类型。

### 操作符与索引

这些类型提供算术、比较、位运算、B-tree 和哈希索引、二进制输入输出、类型转换及 `generate_series`。对应范围类型（包括 `uint1range` 和 `uint16range`）支持 GiST 索引；多范围类型取决于服务器支持。`SUM` 和 `AVG` 内部使用 `numeric`，也提供 `MIN` 和 `MAX`。

有符号与无符号整数的混合运算受操作数顺序和结果类型影响。结果可能为负数时应显式转换类型，因为无符号结果无法表示负数。转换和算术运算可能报越界错误，不能假设数值会回绕。

### 升级至 1.3.0

安装匹配的文件后，以超级用户执行：

```sql
ALTER EXTENSION uint128 UPDATE TO '1.3.0';
```

该更新修正比较操作符的交换关系，并移除混合类型加减运算中无效的交换关系；同时修复 PostgreSQL 17.12 更严格的操作符检查导致的安装问题。提供的 SQL 原地更新操作符元数据，保留数据、索引、视图和物化视图。升级后应检查已安装版本，并验证有代表性的索引比较查询。
