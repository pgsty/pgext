## 用法

来源：

- [目录源码版本的官方 README](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/README.md)
- [目录源码版本的扩展控制文件](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid.control)
- [0.1.4 版扩展 SQL](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid--0.1.4.sql)
- [PGXN 发行页面](https://pgxn.org/dist/is_jsonb_valid/)

is_jsonb_valid 使用原生代码，根据 JSON Schema Draft 4 或 Draft 7 对 JSONB 值进行布尔校验。它适合只需要通过或失败结果的检查与谓词；校验失败时不会返回路径或关键字诊断信息。

### 核心流程

选择与模式草案匹配的函数，第一个参数传模式，第二个参数传被校验的值。

```sql
CREATE EXTENSION is_jsonb_valid;

SELECT is_jsonb_valid(
    '{"type":"object","required":["id"],"properties":{"id":{"type":"integer"}}}'::jsonb,
    '{"id":42}'::jsonb
);

SELECT is_jsonb_valid_draft_v7(
    '{"if":{"exclusiveMaximum":0},"else":{"multipleOf":2}}'::jsonb,
    '4'::jsonb
);
```

0.1.4 版中的两个函数都具有严格与不可变属性。当模式保持稳定时，可以把它们用于约束和索引。

```sql
CREATE TABLE events (
    payload jsonb NOT NULL,
    CONSTRAINT events_payload_valid CHECK (
        is_jsonb_valid(
            '{"type":"object","required":["id"]}'::jsonb,
            payload
        )
    )
);
```

### 草案边界与限制

基础校验器遵循 Draft 4，名称中明确标出的另一个函数遵循 Draft 7。必须主动选择，因为不同草案中的关键字含义不同，扩展也不会根据模式声明自动推断草案。

在固定的源码版本中，不支持远端引用。本地引用仅限于从模式根可达的定义，也不会解析引用链上的标识符变化。此外不支持格式校验。因此，返回真只说明数据满足已实现的子集，并不代表完整 JSON Schema 规范中的全部行为。

扩展没有声明预加载或重启要求。应把模式字面量与应用预期纳入版本控制，并在约束依赖该扩展时加入有代表性的有效与无效用例。
