## 用法

来源：

- [Pinned control file](https://github.com/lmwnshn/hypostats/blob/0181b25d84430d7d7cfc1971b1c94c270cfebbb4/hypostats.control)
- [Pinned implementation](https://github.com/lmwnshn/hypostats/blob/0181b25d84430d7d7cfc1971b1c94c270cfebbb4/src/lib.rs)

`hypostats` 是实验性的 pgrx 扩展，可将某列内部的 `pg_statistic` 行序列化为 JSON 文本，并能把该行写回。它的狭窄用途是用复制或合成的统计信息测试规划器行为；它会直接修改系统目录，并不是通用的统计信息备份格式。

### 核心流程

以超级用户安装，找出关系 OID 和属性编号，再导出当前统计信息：

```sql
CREATE EXTENSION hypostats;

SELECT attnum, attname
FROM pg_attribute
WHERE attrelid = 'public.events'::regclass
  AND attnum > 0
  AND NOT attisdropped;

SELECT pg_statistic_dump(
  'public.events'::regclass::oid::integer,
  2::smallint
);
```

`pg_statistic_dump(integer, smallint)` 返回 JSON 文本；没有匹配的统计元组时返回 `NULL`。把未经修改的导出内容传给 `pg_statistic_load(text)`，它会按 JSON 内嵌的 OID 插入或替换目标行：

```sql
SELECT pg_statistic_load('<json-returned-by-pg_statistic_dump>');
```

目录写入完成后，加载器返回 `true`。此外还暴露了低层辅助函数 `anyarray_elemtype(anyarray)`，用于返回元素类型 OID。

### 安全边界

JSON 包含关系、运算符、排序规则和元素类型的 OID。这些标识符只在本集群内有效，执行转储/恢复、重建扩展或替换对象后都可能改变；加载来自其他集群或对应不同关系的文档，可能破坏规划器统计信息或直接失败。

实现只支持元素类型为 `float4`、`float8`、`int4` 或 `int8` 的统计值数组；其他元素类型会报错。它在行排他锁下直接写入 `pg_statistic`，绕过 `ANALYZE` 的正常计算与校验；之后执行 `ANALYZE` 可能覆盖已加载的行。

只应在可丢弃的测试系统中使用，并保留恢复路径，通过 `EXPLAIN` 验证计划。该扩展版本为 `0.0.0`，不可迁移，安装需要超级用户。它不要求 `shared_preload_libraries`。
