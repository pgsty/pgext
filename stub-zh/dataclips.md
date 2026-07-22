## 用法

来源：

- [官方项目说明](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/docs/dataclip.md)
- [扩展控制文件](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/dataclips.control)
- [扩展 SQL](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/sql/dataclips.sql)
- [C 实现](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/src/dataclips.c)

`dataclips` 是一个已归档、未完成的原型，原计划用于从 PostgreSQL 查询 Heroku Dataclips。上游明确说明它尚不可用；它既不是可用的 Dataclips 客户端，也不是 FDW。

### 声明的接口

发行包声明的扩展名为 `dataclips`，并提供一个返回多态记录的函数：

```sql
CREATE EXTENSION dataclips;

SELECT *
FROM dataclip('gqhkylscctbsgrohaythgdmlfcnn')
  AS t(name text, data json, active boolean);
```

这个查询只是上游设计示例，不是可运行的工作流。由于 `dataclip(text)` 返回 `SETOF record`，每次调用都必须根据预期的 Dataclip 输出提供列定义列表。

### 实现状态

C 函数中只有对 HTTP 获取、CSV/JSON 解析、元组描述符和行生成的计划注释，实际逻辑均未实现。它硬编码了三次调用，并返回零 `Datum`，没有构造任何元组。项目也没有创建任何外部数据包装器对象。

软件包本身也不一致：Makefile 构建名为 `dataclips` 的动态库，而 `dataclips.control` 请求 `$libdir/dataclip`；SQL 脚本的保护提示还写成 `CREATE EXTENSION dataclip`，但规范扩展名实际是 `dataclips`。因此标准构建甚至可能在进入未完成函数前就失败。

### 注意事项

不要部署此扩展，也不要向用户开放 `dataclip(text)`。它没有实现认证、响应校验、错误模型或稳定的行映射，而且仓库已经归档。该源码只能作为历史原型材料；真正的集成需要维护中的客户端或 FDW，并明确处理凭据、网络和模式。
