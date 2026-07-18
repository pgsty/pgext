## 用法

来源：

- [上游 README](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/README.md)
- [0.0.5 版控制文件](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/postgis_letters.control.in)
- [几何函数定义](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/sql_bits/postgis_letters.sql.in)
- [内置字体数据](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/sql_bits/font_set.sql.in)

postgis_letters 0.0.5 是一个纯 SQL 的 PostGIS 附加组件，可将文本转换为多多边形字形几何。安装会创建 font_set 表并载入内置 Kankin 矢量字体。

### 将文本渲染为几何

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_letters;
SELECT ST_LettersAsGeometry('Circle', 'kankin');
```

可选的 SRID、文字高度与起始位置参数用于放置和缩放结果。尺寸与位置必须使用目标几何的坐标单位。

### 注意事项

- 指定 SRID 只是标记输出坐标，并不会投影字形。当涉及地理或投影精度时，应先得到合适的平面结果再进行转换。
- 所选字体中不存在的字符不会生成字形。存储结果前应验证支持的字符集。
- 复杂文本会生成大型多多边形值，其存储、渲染、索引或相交计算成本可能很高。
- SQL 函数引用 font_set 时没有模式限定。应安装在受控模式中并使用受控搜索路径，避免解析到非预期表。
- 扩展转储只保留 packaged_font 标志为 false 的 font_set 行；内置行会被安装或升级脚本替换。应备份自定义字体并谨慎设置该标志。
- README 面向 PostgreSQL 9.1 与 PostGIS 2.0 及以上版本，但上游没有发布当前兼容矩阵或自动化测试。
