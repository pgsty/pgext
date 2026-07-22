## 用法

来源：

- [固定提交的上游 README](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/README.md)
- [固定提交的 1.1 版安装 SQL](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/zson--1.1.sql)
- [固定提交的扩展 control 文件](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/zson.control)
- [固定提交的官方 benchmark 说明](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/docs/benchmark.md)

`zson` `1.1` 是一种使用字典压缩、兼容 JSONB 操作的存储类型。它从代表性 JSONB 文档中学习频繁重复的字符串，包括 key、字符串值和数组元素，然后在编码新的 `zson` 值时使用该共享字典。

### 核心流程

插入压缩值前先训练字典。训练参数是由表名与 JSONB 列名组成的二维数组：

```sql
CREATE EXTENSION zson;

CREATE TABLE zson_training (doc jsonb);
INSERT INTO zson_training VALUES
  ('{"kind":"event","service":"billing","status":"ok"}'::jsonb),
  ('{"kind":"event","service":"billing","status":"failed"}'::jsonb);

SELECT zson_learn('{{"zson_training","doc"}}'::text[][]);
SELECT * FROM zson_dict ORDER BY dict_id, word_id;

CREATE TABLE zson_events (id bigint GENERATED ALWAYS AS IDENTITY, payload zson);
INSERT INTO zson_events(payload)
SELECT doc::zson FROM zson_training;

SELECT payload ->> 'service' FROM zson_events;
SELECT zson_info(payload) FROM zson_events LIMIT 1;
```

文档词汇发生明显变化后应再次运行 `zson_learn()`。新值使用新字典，已有值则保留原来的字典标识。

### 重要对象

- `zson`：变长压缩类型，提供从 `jsonb` 到自身的 assignment cast 和到 `jsonb` 的 implicit cast，因此解压后可使用 JSONB 操作符。
- `zson_learn(tables_and_columns, max_examples, min_length, max_length, min_count)`：采样 JSONB 列，创建下一个最多包含 65,534 个高频字符串的字典。
- `zson_dict(dict_id, word_id, word)`：由扩展管理、包含在 extension configuration dump 中的字典表。
- `zson_info(zson)`：报告已存值的格式和字典版本信息。
- `zson_extract_strings(jsonb)`：递归提取训练时使用的对象 key 与字符串值。

### 运维说明

扩展只支持安装在 `public` schema。`zson_learn()` 会根据表名和列名构造动态 SQL 并扫描样本数据，因此只能允许可信管理员执行；应限制 `max_examples`，并使用具有代表性的分布训练。压缩率和吞吐收益取决于数据与负载；应在真实数据集上复现官方 benchmark 方法，不要把已发布数字当作保证。

字典会被缓存，上游指南说明新字典可能约需一分钟才会在所有后端生效。只要仍有已存值引用旧 `dict_id`，就绝不能删除它，否则这些文档可能无法读取；保留几 KB 的旧字典数据更安全。用 `zson` 替换 JSONB 列前，应测试包含 `zson_dict` 的备份恢复、混合字典版本、转换为 `jsonb` 的表达式索引以及解压 CPU 成本。
