


## 用法

> [omni_yaml: YAML 工具箱](https://docs.omnigres.org/omni_yaml/yaml/)

`omni_yaml` 扩展提供 YAML 与 JSON 之间的转换函数。

### 函数

**`omni_yaml.to_json(text)`** -- 将 YAML 文本转换为 JSON：

```sql
SELECT omni_yaml.to_json('key: value');
-- 返回: {"key": "value"}
```

**`omni_yaml.to_yaml(json)`** -- 将 JSON 转换为 YAML 文本：

```sql
SELECT omni_yaml.to_yaml('{"key": "value"}'::json);
-- 返回: key: value
```

### JSONB 用法

不提供直接的 JSONB 支持。根据需要将 JSON 结果转换为 `jsonb`：

```sql
SELECT omni_yaml.to_json('key: value')::jsonb;
```
