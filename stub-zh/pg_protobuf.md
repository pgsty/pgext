
## 用法

> [pg_protobuf: PostgreSQL 的 Protocol Buffers 支持](https://github.com/afiskon/pg_protobuf)

该扩展提供无需 schema 定义即可在 SQL 中直接解码 Protocol Buffer 二进制数据的函数。

### 函数

- `protobuf_decode(bytea) RETURNS cstring` -- 将 protobuf 二进制数据解码为可读格式
- `protobuf_get_int(bytea, int) RETURNS bigint` -- 按字段编号提取整数字段
- `protobuf_get_string(bytea, int) RETURNS text` -- 按字段编号提取字符串字段
- `protobuf_get_bytes(bytea, int) RETURNS bytea` -- 按字段编号提取原始字节
- `protobuf_get_bool(bytea, int) RETURNS boolean` -- 按字段编号提取布尔字段
- `protobuf_get_float(bytea, int) RETURNS real` -- 按字段编号提取浮点字段
- `protobuf_get_double(bytea, int) RETURNS double precision` -- 按字段编号提取双精度字段
- `protobuf_get_sfixed32(bytea, int) RETURNS int` -- 提取有符号定长 32 位字段
- `protobuf_get_sfixed64(bytea, int) RETURNS bigint` -- 提取有符号定长 64 位字段
- `protobuf_get_int_multi(bytea, int) RETURNS bigint[]` -- 提取重复整数字段
- `protobuf_get_string_multi(bytea, int) RETURNS text[]` -- 提取重复字符串字段
- `protobuf_get_bytes_multi(bytea, int) RETURNS bytea[]` -- 提取重复字节字段
- `protobuf_get_bool_multi(bytea, int) RETURNS boolean[]` -- 提取重复布尔字段
- `protobuf_get_float_multi(bytea, int) RETURNS real[]` -- 提取重复浮点字段
- `protobuf_get_double_multi(bytea, int) RETURNS double precision[]` -- 提取重复双精度字段
- `protobuf_get_sfixed32_multi(bytea, int) RETURNS int[]` -- 提取重复 sfixed32 字段
- `protobuf_get_sfixed64_multi(bytea, int) RETURNS bigint[]` -- 提取重复 sfixed64 字段

### 示例

```sql
CREATE EXTENSION pg_protobuf;

-- 创建包含 protobuf 数据的表
CREATE TABLE heroes (x bytea);

-- 为特定字段定义访问函数
CREATE FUNCTION hero_name(x bytea) RETURNS text AS $$
BEGIN
  RETURN protobuf_get_string(x, 512);
END $$ LANGUAGE plpgsql IMMUTABLE;

CREATE FUNCTION hero_hp(x bytea) RETURNS bigint AS $$
BEGIN
  RETURN protobuf_get_int(x, 2);
END $$ LANGUAGE plpgsql IMMUTABLE;

-- 在 protobuf 字段上创建索引
CREATE INDEX hero_name_idx ON heroes USING btree(hero_name(x));

-- 按 protobuf 字段查询
SELECT hero_name(x) FROM heroes ORDER BY hero_name(x) LIMIT 10;
```

### 限制

- 不支持修改 Protobuf 数据
- 枚举值可通过 `protobuf_get_int` 读取
- 不直接支持无符号类型（PostgreSQL 中没有无符号整数）
- `*_multi` 函数不支持 `[packed=true]`（可改用 `protobuf_get_bytes*`）
