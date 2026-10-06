## 用法

来源：

- [README.md](https://github.com/ancoron/pg-uuid-v1/blob/577b1db78f16351199613d88e3e847c08b45118e/README.md)
- [uuid_v1--0.1.sql](https://github.com/ancoron/pg-uuid-v1/blob/577b1db78f16351199613d88e3e847c08b45118e/uuid_v1--0.1.sql)
- [uuid_v1.control](https://github.com/ancoron/pg-uuid-v1/blob/577b1db78f16351199613d88e3e847c08b45118e/uuid_v1.control)

`uuid_v1` 用一个优先按时间戳分量排序的类型保存第 1 版 UUID，提供比较算子、与 uuid 的相互转换、B-tree 支持及分量提取函数。

### 核心用法

```sql
CREATE EXTENSION uuid_v1;
SELECT uuid_v1_get_timestamp('b647e96b-862d-11e9-ae2b-db6f0f573554'::uuid_v1);
SELECT uuid_v1_get_epoch('b647e96b-862d-11e9-ae2b-db6f0f573554'::uuid_v1);
```

### 运行边界

`uuid_v1_get_timestamp` 返回 timestamptz，`uuid_v1_get_epoch` 返回不可变的纪元时间值，`uuid_v1_get_clockseq` 与 `uuid_v1_get_node` 提取其余分量。时间戳比较使用完整精度，精确到秒的时间不一定等于带小数秒的 UUID。

此扩展处理已有的第 1 版 UUID，并不生成它们。安装要求超级用户及 C 库，不需要预加载。上游未声明当前 PostgreSQL 主版本矩阵；需要不可变表达式索引时，应按上游说明选用 epoch 辅助函数。
