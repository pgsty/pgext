## 用法

来源：

- [pgzint 0.2.0 README](https://github.com/davidbeauchamp/pgzint/blob/v0.2.0/README.md)
- [pgzint 0.2.0 发行说明](https://github.com/davidbeauchamp/pgzint/releases/tag/v0.2.0)
- [pgzint 0.2.0 扩展 SQL](https://github.com/davidbeauchamp/pgzint/blob/v0.2.0/pgzint--0.2.0.sql)
- [pgzint 0.2.0 控制文件](https://github.com/davidbeauchamp/pgzint/blob/v0.2.0/pgzint.control)

`pgzint` 使用 Zint 库在 PostgreSQL 内生成条码图像，并以 `bytea` 返回 PNG 字节。适用于应用可以消费二进制图像、但不希望直接集成 Zint 的场景。

### 核心流程

安装扩展、检查码制目录，再调用便利包装器或通用生成器：

```sql
CREATE EXTENSION pgzint;

SELECT bc_symbol_zint_id, bc_symbol_zint_constant, bc_symbol_name
FROM bc_symbols
ORDER BY bc_symbol_zint_id;

SELECT bc_qrcode('SAMPLE');

SELECT bc_generate(
  'SAMPLE', 58, NULL, 2, 0, NULL, NULL,
  NULL, NULL, NULL, NULL, 14, NULL, 0
);
```

`bc_generate` 接收载荷、Zint 码制 ID、高度、缩放、留白和边框宽度、输出标志、颜色、文字标志、三个码制专用选项以及旋转角度。0.2.0 中高度参数改为 `float8`。

### 重要对象

- `bc_generate` 是以 C 实现的通用生成器。
- `bc_qrcode`、`bc_excode39`、`bc_pdf417`、`bc_maxicode` 与 `bc_code128` 是带项目预设值的 SQL 包装器。
- `bc_symbols` 映射 Zint 数字 ID、常量与显示名称。
- `getzintsymbolid(text)` 与 `getzintsymbolconstant(integer)` 用于在两种标识之间转换。
- `pgzint_version()` 返回已安装的 pgzint 版本。

0.2.0 删除了旧的 `barcodes` 视图并简化了 `bc_symbols`；集成代码必须使用上例中的三个列，而不能继续依赖本版本删除的元数据列。

### 升级与要求

安装 0.2.0 软件包文件后，在每个已经安装 pgzint 的数据库中执行升级：

```sql
ALTER EXTENSION pgzint UPDATE TO '0.2.0';
```

pgzint 0.2.0 要求 PostgreSQL 9.4 或更高版本，以及启用了 PNG 支持的 Zint 2.14 或更高版本。它改用 Zint 的内存 PNG 输出，不再经过旧版的 BMP 到 PNG 转换，也不再直接维护 libpng 转换层。

图像生成会消耗数据库 CPU，并可能返回很大的二进制值。应限制载荷大小，避免在延迟敏感查询中未经评估地批量生成，并用应用要求的扫码器与条码标准验证输出。
