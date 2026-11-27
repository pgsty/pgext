## 用法

来源：

- [Official source](https://github.com/df7cb/postgresql-unit/blob/7.11/README.md)
- [Official source](https://github.com/df7cb/postgresql-unit/blob/7.11/unit.control)
- [Official source](https://github.com/df7cb/postgresql-unit/blob/7.11/Makefile)
- [Official source](https://github.com/df7cb/postgresql-unit/releases/tag/7.11)

`unit` 7.11 提供具有量纲运算和转换能力的物理量类型。软件包发行版保留 SQL/control 版本 7，依赖 `plpgsql`。它包含编译代码并安装单位定义表，需要管理权限安装，无需共享预加载。

### 物理量与转换

```sql
CREATE EXTENSION unit;
SELECT '9.81 m/s^2'::unit;
SELECT '120 km/h'::unit @ 'm/s';
SELECT '1 km'::unit @@ 'm';
SELECT '5 kg'::unit * '9.81 m/s^2'::unit;
```

`@` 返回以目标单位表示的文本，`@@` 返回双精度数值。加减要求量纲兼容，乘除及整数幂会组合量纲。八个基本维度包含七个 SI 维度及字节。物理量以浮点数保存，并非任意精度十进制运算。

### 解析与输出

输入支持分数、时间记法、前缀及 Unicode 指数。7.11 新增数值字面量中的下划线支持，并将内置 GNU 单位定义更新到 3.26。对精确常量或温度偏移有要求时，应验证转换结果。

`unit.byte_output_iec` 选择二进制字节前缀，`unit.output_base_units`、`unit.time_output_custom`、`unit.output_superscript` 控制显示。扩展提供数学函数、聚合和 `unitrange`，使用时需保证量纲兼容。

### 维护

`unit_prefixes` 和 `unit_units` 表可以安装到选定模式，但安装后扩展不可重定位。升级时保留本地自定义单位。安装匹配的库和数据文件，并重连使用旧库的会话；选择可用升级路径前检查当前 SQL 版本。7.11 发行版本身仍使用 SQL 版本 7。
