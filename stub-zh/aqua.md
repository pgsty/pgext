## 用法

来源：

- [官方上游 README](https://github.com/zjudbxai/peps/blob/1923b7e389a8028dc71cfcfaa2bddd8b491c62c4/README.md)
- [官方扩展控制文件 (aqua.control)](https://github.com/zjudbxai/peps/blob/1923b7e389a8028dc71cfcfaa2bddd8b491c62c4/aqua_extension/aqua.control)
- [官方扩展 SQL (aqua--0.0.1.sql)](https://github.com/zjudbxai/peps/blob/1923b7e389a8028dc71cfcfaa2bddd8b491c62c4/aqua_extension/sql/aqua--0.0.1.sql)

`aqua` — 在数据库中实现的神经用户定义函数，用于机器学习推理。将其用于相应的向量、模型或检索工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION aqua;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `argmax(real[])` 是一个扩展函数，返回 `integer`。
- `array_2d_agg_combine(internal, internal)` 是一个扩展函数，返回 `internal`。
- `array_2d_agg_final_array(internal)` 是一个扩展函数，返回 `real[]`。
- `array_2d_agg_state(internal, real[][])` 是一个扩展函数，返回 `internal`。
- `array_agg_array_deserialize(bytea, internal)` 是一个扩展函数，返回 `internal`。
- `array_agg_array_serialize(internal)` 是一个扩展函数，返回 `bytea`。
- `avgpool(input_f real[][], kernel INT = 2, padding INT = 0, stride INT = 2, op_cost REAL = 0.0)` 是一个扩展函数，返回 `real[]`。
- `avgpool_conv(kmatrix real[][], fmatrix real[][], bias real[], kernel INT = 2, padding INT = 0, stride INT = 2)` 是一个扩展函数，返回 `real[]`。
- `batchnorm(fmatrix real[][], args real[])` 是一个扩展函数，返回 `real[]`。
- `check_weight_pool()` 是一个扩展函数，返回 `VOID`。
- `concat_array(num int, input1 REAL[][] DEFAULT NULL, input2 REAL[][] DEFAULT NULL, input3 REAL[][] DEFAULT NULL, input4 REAL[][] DEFAULT NULL, input5 REAL[][] DEFAULT NULL, input6 REAL[][] DEFAULT NULL, input1_offset INT [] DEFAULT NULL)` 是一个扩展函数，返回 `real[]`。
- `conv_text(kmatrix text, fmatrix real[][], kernel_H INT, kernel_W INT, padding_H INT, padding_W INT, stride INT)` 是一个扩展函数，返回 `real[]`。
- `group_kfm_im2col(kmatrix text, fmatrix real[][][], kernel INT = 3, padding INT = 1, stride INT = 1, groups INT = 8)` 是一个扩展函数，返回 `real[]`。
- `im2col(input_f real[][], kernel INT = 3, padding INT = 1, stride INT = 1)` 是一个扩展函数，返回 `real[]`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.0.1`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
