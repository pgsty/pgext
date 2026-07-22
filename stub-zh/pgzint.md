## 用法

来源：

- [0.1.4 版官方 README](https://github.com/davidbeauchamp/pgzint/blob/47cdb014cd81752c9d42b1d06c85871872a1793f/README.md)
- [0.1.4 版扩展 SQL](https://github.com/davidbeauchamp/pgzint/blob/47cdb014cd81752c9d42b1d06c85871872a1793f/pgzint--0.1.4.sql)
- [PGXN 发行页面](https://pgxn.org/dist/pgzint/)

`pgzint` 通过 Zint 条码库在 PostgreSQL 内生成条码图像。其 SQL 函数以 `bytea` 返回 PNG 字节，适合能够显示二进制图像、但不希望直接集成 Zint 的应用。

### 核心流程

安装扩展、选择码制，然后返回或存储生成的二进制结果：

```sql
CREATE EXTENSION pgzint;

-- Convenience wrapper with project-provided QR defaults.
SELECT bc_qrcode('SAMPLE');

-- Inspect the installed symbol catalog before choosing an integer ID.
SELECT symbol_id, symbol_name
FROM barcodes
ORDER BY symbol_id;
```

`barcodes` 的实际列来自已安装的 0.1.4 版 SQL，因此客户端集成时应使用 `\d+ barcodes` 检查。主 API 接收载荷、Zint 码制编号及渲染选项：

```sql
SELECT bc_generate(
  'SAMPLE', 58, NULL, 2, 0, NULL, NULL,
  NULL, NULL, NULL, NULL, 14, NULL, 0
);
```

该示例使用上游 README 所述的同一种 QR 码制与默认值。渲染参数会传递给 Zint，应以扩展实际搭配的 Zint 版本确认其语义。

### 重要对象

- `bc_generate` 是通用的 C 实现生成器。参数控制码制、高度、缩放、留白、边框、输出标志、前景/背景颜色、文字显示、三个码制专用选项与旋转角度。
- `bc_qrcode`、`bc_excode39`、`bc_pdf417`、`bc_maxicode` 与 `bc_code128` 是使用项目预设默认值的 SQL 便利包装器。
- `bc_symbols` 保存扩展的码制元数据，`barcodes` 以视图形式呈现这些信息。
- `getzintsymbolconstant` 与 `getzintsymbolid` 在生成器使用的数字常量和文本标识之间转换。

此版本只有 `bc_generate` 直接以 C 实现；便利函数均委托给它。若包装器默认值不符合所需条码大小或编码选项，应直接调用主函数。

### 前置条件与注意事项

0.1.4 版要求 PostgreSQL 9.1 或更高版本，并在构建与运行时具备 Zint 和 PNG 库。扩展控制文件没有声明预加载或其他 PostgreSQL 扩展依赖。由早期版本升级的现有数据库在安装 0.1.4 文件后必须执行扩展升级：

```sql
ALTER EXTENSION pgzint UPDATE TO '0.1.4';
```

图像生成会消耗数据库 CPU 和内存，并可能返回较大的二进制值，因此应限制载荷大小，避免在延迟敏感查询中未经评估地批量生成。该项目较旧，内置码制元数据反映的是其所处时代的 Zint 定义；应使用应用要求的扫码器与标准验证输出。
