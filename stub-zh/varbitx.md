## 用法

来源：

- [阿里云 RDS varbitx 文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-varbitx-plug-in)

`varbitx` 版本 `1.0` 是阿里云 ApsaraDB RDS for PostgreSQL 扩展，为变长位串增加切片、批量修改、位置查找、生成和计数函数。它是供应商托管功能，并没有被文档描述为可移植的社区软件包。

### 核心流程

阿里云文档说明该扩展用于受支持的 RDS PostgreSQL 10 与 11 实例。

```sql
CREATE EXTENSION varbitx;

SELECT get_bit('111110000011'::varbit, 3, 5);
SELECT bit_count('1111000011110000'::varbit, 1, 5, 4);
SELECT set_bit_array_record(
  '111100001111'::varbit,
  1,
  0,
  ARRAY[1,4,5,6,7],
  2
);
```

文档中函数的位置从零开始。`get_bit` 提取切片。`set_bit_array` 返回修改后的位串；`set_bit_array_record` 还会返回发生变化的位置。可选限制可在达到指定变更次数后停止。

### 函数组与注意事项

提取使用 `get_bit` 和 `get_bit_array`；生成使用 `bit_fill` 与概率性的 `bit_rand`；计数使用 `bit_count` 与 `bit_count_array`；`bit_posite` 按升序或降序返回匹配位置。

修改函数返回新值，并不会原地更新表。填充值、结果长度、顺序和限制参数都会实质性改变结果，因此应测试边界位置，以及短于目标索引的位串。可用性与升级由 RDS 实例系列和供应商发行控制；将其写入存储过程前，应在目标实例上核实确切数据库版本与函数签名。
