## 用法

来源：

- [官方 README](https://github.com/freevryheid/units/blob/eb52d6246e6fda6668e3e50cb5048e80453dc8d0/README.md)
- [官方扩展 SQL](https://github.com/freevryheid/units/blob/eb52d6246e6fda6668e3e50cb5048e80453dc8d0/units--0.0.1.sql)

`units` 是一个小型 PL/pgSQL 学习扩展，把长度换算系数存入表中，再用选中的系数乘以输入值。它适合作为清晰示例或本地维护的换算表，而不是完整的单位系统。

### 核心流程

实现会把两个单位参数都当作 PostgreSQL 正则表达式。若要准确选择一对单位，应锚定名称。

```sql
CREATE EXTENSION units;

SELECT unitsConvert(1, '^mile$', '^meter$');
SELECT unitsConvert(12, '^inch$', '^foot$');

SELECT msr, uin, niu, mop
FROM units
WHERE uin = 'meter' AND niu = 'foot';
```

### 对象

- `units(msr, uin, niu, mop)` 存储度量类别、输入单位、输出单位和乘数，主键为 `(uin, niu)`。
- `unitsConvert(numeric, text, text) RETURNS numeric` 使用 `uin ~ input_pattern` 与 `niu ~ output_pattern` 查找系数，再返回输入值乘以该系数的结果。

内置数据覆盖 kilometer、meter、centimeter、millimeter、micrometer、nanometer、mile、yard、foot、inch 与 nautical mile 之间的长度换算。

### 注意事项

由于查询使用正则表达式和非 strict 的 `SELECT INTO`，宽泛模式可能匹配多行，并从中选择未指定的一行。应使用 `^mile$` 之类的锚定模式；无匹配时返回 `NULL`。若该表成为应用数据，应复核换算精度并有意维护其内容。

上游 README 将项目描述为学习实验；版本 `0.0.1` 不提供更广泛的量纲分析、单位别名、复合单位、跨度量类别校验或已记录的许可证。若扩展未安装在默认模式中，应限定 `units` 表与 `unitsConvert` 函数的模式。
