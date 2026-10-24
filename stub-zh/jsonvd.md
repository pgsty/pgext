## 用法

来源：

- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/README.md)
- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/jsonvd.control)
- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/sql/jsonvd--1.0.0.sql)
- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/docs/02-design/api-contracts.md)

`jsonvd` 1.0.0 是面向 PostgreSQL 14–18 的 SQL/PLpgSQL JSONB 工具扩展，提供嵌套值搜索、文档比较与合并，以及类型推断。它是可信扩展，使用固定的 `jsonvd` 模式，无须动态库、预加载或重启；安装者仍须具有在数据库中创建扩展的权限。

### 基本流程

```sql
CREATE EXTENSION jsonvd;
SELECT jsonvd.find_path('{"items":[{"id":1},{"id":2}]}'::jsonb, '2'::jsonb);
SELECT * FROM jsonvd.diff('{"a":1,"b":[1,2]}'::jsonb,
                          '{"a":2,"b":[1,3]}'::jsonb);
SELECT jsonvd.merge_deep('{"db":{"host":"localhost","port":5432}}'::jsonb,
                        '{"db":{"port":6543}}'::jsonb);
SELECT jsonvd.merge_patch('{"a":1,"b":2}'::jsonb, '{"a":null}'::jsonb);
```

路径为 `text[]`，可用于 PostgreSQL JSON 运算符及更新函数。匹配使用 JSONB 相等规则。`find_path` 返回首个匹配路径或 SQL NULL，`find_paths` 返回全部匹配。`contains_deep` 判断是否包含目标，`flatten` 返回每个节点的路径、值、类型与深度。

### 转换与推断

`diff` 返回新增、删除或替换记录，数组按位置比较。`merge_deep` 递归合并对象，其他冲突由右侧覆盖，忽略 SQL NULL 一侧。`merge_patch` 遵循 JSON Merge Patch，用 JSON null 删除对象键。SQL NULL 与 JSON null 是不同输入。

`set_path` 创建缺少的中间对象，`delete_key_deep` 递归删除键，`push` 添加或替换对象字段。`parse_value` 解析合法 JSON 文本，其他文本转为 JSON 字符串。`infer_schema` 和 `recordset_definition` 根据样本对象推断列定义；该定义描述样本，并不约束未来数据。

### 使用边界

函数声明为不可变且并行安全。严格函数遇到 NULL 参数会返回 SQL NULL；合并和差异函数则按各自文档处理。无效路径或参数形状可能引发 SQL 错误。用这些辅助函数替换应用原有合并规则前，应核对空值、标量和数组的行为。
