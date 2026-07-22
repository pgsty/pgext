## 用法

来源：

- [官方 README](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/README.md)
- [官方控制文件](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/pg_substrait.control)
- [官方构建清单](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/Cargo.toml)

`pg_substrait` 通过 PostgreSQL 原生执行器运行 Substrait 查询计划。它接受 protobuf 字节或 JSON，并返回动态推断的记录集，适合与能够生成 Substrait 计划的系统做互操作实验；当前版本为 `0.1.0`，支持的关系、表达式与类型映射都必须实测，不能想当然。

### 核心流程

```sql
CREATE EXTENSION pg_substrait;

SELECT *
FROM from_substrait_json('{
  "version": {"minorNumber": 54},
  "relations": [{
    "root": {
      "names": ["result"],
      "input": {
        "project": {
          "expressions": [{"literal": {"i32": 42}}]
        }
      }
    }
  }]
}');
```

扩展会推断输出名称与类型，因此 `AS` 子句可选。调用者需要显式稳定的行契约时应提供该子句。二进制计划把 protobuf 数据传给 `from_substrait(bytea)`；`from_substrait_json(text)` 处理 JSON。

### 重要对象

- `from_substrait(plan bytea)` 解码并执行 protobuf 计划。
- `from_substrait_json(json_plan text)` 解码并执行 JSON 计划。
- 两者都返回带扩展推断模式的 `SETOF RECORD`。

### 限制

Substrait 计划是可执行输入，不是被动数据。应限制函数执行权、设置语句和资源上限，不能因为输入是有效 JSON 或 protobuf 就接受不可信计划。还要在打包构建上验证 PostgreSQL 权限确实限制计划引用的每个关系和函数。

README 声明支持 PostgreSQL 13 至 17，而构建清单默认 PostgreSQL 15，并固定 pgrx 与 Substrait 库版本。不同版本的生产者可能发出本版本无法理解的新扩展或语义。应维护生产者与消费者兼容测试，覆盖类型、空值、函数、连接、错误、事务行为，以及畸形或超大计划。
