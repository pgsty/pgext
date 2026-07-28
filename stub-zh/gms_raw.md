## 用法

来源：

- [官方上游 README](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/README)
- [官方扩展控制文件 (gms_raw.control)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_raw/gms_raw.control)
- [官方扩展 SQL (gms_raw--1.0.sql)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_raw/gms_raw--1.0.sql)

`gms_raw` — 处理 PL/SQL 应用程序的原始类型数据。在移植或模拟相应的数据库 API 时使用。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION gms_raw;
```

在目标数据库中安装扩展，当有可用示例时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gms_raw.bit_and(r1 in raw, r2 in raw)` 是一个扩展函数，返回 `raw`。
- `gms_raw.bit_complement(r1 in raw)` 是一个扩展函数，返回 `raw`。
- `gms_raw.bit_or(r1 in raw, r2 in raw)` 是一个扩展函数，返回 `raw`。
- `gms_raw.bit_xor(r1 in raw, r2 in raw)` 是一个扩展函数，返回 `raw`。
- `gms_raw.cast_from_binary_double(n in binary_double, endianess in integer default 1)` 是一个扩展函数，返回 `raw`。
- `gms_raw.cast_from_binary_float(n in float, endianess in integer default 1)` 是一个扩展函数，返回 `raw`。
- `gms_raw.cast_from_binary_integer(n in bigint, endianess in integer default 1)` 是一个扩展函数，返回 `raw`。
- `gms_raw.cast_from_number(n in number)` 是一个扩展函数，返回 `raw`。
- `gms_raw.cast_to_binary_double(r in raw, endianess in integer default 1)` 是一个扩展函数，返回 `binary_double`。
- `gms_raw.cast_to_binary_float(r in raw, endianess in integer default 1)` 是一个扩展函数，返回 `float4`。
- `gms_raw.cast_to_binary_integer(r in raw, endianess in integer default 1)` 是一个扩展函数，返回 `binary_integer`。
- `gms_raw.cast_to_number(r in raw)` 是一个扩展函数，返回 `number`。
- `gms_raw.cast_to_nvarchar2(r in raw)` 是一个扩展函数，返回 `nvarchar2`。
- `gms_raw.cast_to_raw(c in varchar2)` 是一个扩展函数，返回 `raw`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
