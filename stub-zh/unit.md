

## 用法

> [unit: PostgreSQL 的 SI 单位数据类型](https://github.com/df7cb/postgresql-unit)

`unit` 扩展提供了 SI 单位数据类型，可在 SQL 中直接进行量纲分析和单位换算。

```sql
CREATE EXTENSION unit;

SELECT '9.81 m/s^2'::unit;
SELECT '120 km/h'::unit @ 'm/s' AS velocity;  -- 33.3333333333333 m/s
```

### 基本单位

米 (m)、千克 (kg)、秒 (s)、安培 (A)、开尔文 (K)、摩尔 (mol)、坎德拉 (cd)、字节 (B)。

### 运算符

| 运算符 | 说明 | 示例 |
|----------|-------------|---------|
| `+`, `-` | 加/减（需相同量纲） | `'1 m'::unit + '50 cm'::unit` |
| `*`, `/` | 乘/除 | `'5 kg'::unit * '9.81 m/s^2'::unit` |
| `^` | 整数次幂 | `'2 m'::unit ^ 3` |
| `@` | 转换单位（返回 unit） | `'2 MB/min'::unit @ 'GB/d'` |
| `@@` | 转换单位（返回 double precision） | `'1 km'::unit @@ 'm'` |

### 函数

数学函数：`sqrt()`、`exp()`、`ln()`、`log2()`、`cbrt()`、`asin()`、`tan()` 等。

聚合函数：`sum(unit)`、`avg(unit)`、`min(unit)`、`max(unit)`、`stddev()`、`variance()`。

### 输入格式

```sql
SELECT '3|4 m'::unit;            -- 分数：0.75 m
SELECT '10:05:30 s'::unit;       -- 时间格式：36330 s
SELECT 'm⁻²'::unit;              -- Unicode 上标
```

### 单位换算

```sql
SELECT '2 MB/min'::unit @ 'GB/d';       -- 2.88 GB/d
SELECT '1 hl'::unit @ '0.5 l';          -- 200 * 0.5 l
SELECT '100 degC'::unit @ 'degF';        -- 华氏温度转换
```

### 范围类型

```sql
SELECT unitrange('earthradius_polar', 'earthradius_equatorial');
```

### 配置

- `unit.byte_output_iec`：二进制前缀（Ki, Mi, Gi）
- `unit.output_base_units`：仅显示基本单位
- `unit.time_output_custom`：使用分/时/日格式化时间
- `unit.output_superscript`：Unicode 上标指数
