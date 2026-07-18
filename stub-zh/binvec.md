## 用法

来源：

- [上游 README](https://github.com/sany1231/binvec/blob/1bca9cab92fb199717c206e7d6601b82c727cf04/README.md)
- [0.1 版安装 SQL](https://github.com/sany1231/binvec/blob/1bca9cab92fb199717c206e7d6601b82c727cf04/binvec--0.1.sql)
- [C 实现](https://github.com/sany1231/binvec/blob/1bca9cab92fb199717c206e7d6601b82c727cf04/binvec.c)

binvec 0.1 提供 vec_sum_bin(anyarray, integer)，这是一个将整数中的置位位累加到 32 元素向量的转换函数。README 展示的聚合函数并不会随扩展安装，需要时必须另行创建。

### 基本调用

只能在隔离测试数据库中使用 integer 数组：

```sql
CREATE EXTENSION binvec;
SELECT vec_sum_bin(ARRAY[1, 0, 1]::integer[], 5);
```

结果包含 32 个元素。第零位对应第一个元素，因此数值 5 会使第一和第三个元素递增。

### 注意事项

- C 代码虽然接受多种数值数组元素类型，却始终按 32 位整数读写每个元素。integer 以外的数组可能损坏或产生无效结果。
- 输入数组的空值标记会被忽略，多维数组会被拒绝。空数组返回空值，而右侧数值为零时会直接返回原数组，不会扩展到 32 个元素。
- 有符号最高位与负数输入没有得到安全处理。应将右侧数值限制为不小于零且小于 2^31。
- 这是一个没有测试、许可证、发布历史或当前 PostgreSQL 兼容性说明的旧原型。任何非一次性用途都必须先验证精确构建版本。
