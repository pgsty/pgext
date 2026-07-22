## 用法

来源：

- [官方 README](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/README.md)
- [基础扩展 SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.0.0.sql)
- [向量类型升级 SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.0.0--1.1.0.sql)
- [模型 API 升级 SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.1.0--1.2.0.sql)
- [当前 API 升级 SQL](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/sql/pgdl--1.2.0--1.3.0.sql)
- [模型管理器实现](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/src/pgdl/model_manager.cpp)
- [扩展控制文件](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/pgdl.control)

`pgdl` 1.3.0 也称 MorphingDB，在 PostgreSQL 中嵌入深度学习推理与带形状的向量存储。它通过 libtorch 在数据库后端加载 TorchScript 模型，并提供自定义 `mvec` 类型。复核的项目面向 PostgreSQL 12，需要携带大量原生依赖和模型专用 C++ 处理器进行源码构建；应把它视为研究原型，而不是即装即用的推理服务。

### 向量存储流程

验证本地修补后的安装后，可先使用 API 中相对独立的部分，把数组及其形状一同存储：

```sql
CREATE EXTENSION pgdl;

CREATE TABLE feature_vectors (
    id bigint PRIMARY KEY,
    value mvec NOT NULL
);

INSERT INTO feature_vectors VALUES
    (1, '[1.0,2.2,3.123,4.2]{2,2}'),
    (2, ARRAY[1.0,2.0,3.0,1.2345]::float4[]::mvec);

SELECT id, get_mvec_shape(value), get_mvec_data(value)
FROM feature_vectors;

UPDATE feature_vectors SET value = value + value WHERE id = 1;
```

`mvec` 接受文本或数值数组，可以保留多维形状，支持加、减与相等运算；可用 `get_mvec_data(mvec)` 转回数组，并用 `get_mvec_shape(mvec)` 查看形状。

### 模型推理流程

以模型为中心的用法从 `create_model` 开始，它在 `model_info` 中记录服务端本地模型路径与可选基础模型。`register_process()` 绑定编译进扩展的前/后处理回调，`predict_float` 或 `predict_text` 在 CPU 或 GPU 上运行已注册模型。移动窗口聚合 `predict_batch_float8` 与 `predict_batch_text` 提供批量推理。1.3.0 版增加了面向任务的 `ai_operator`、`api_load_model` 与 `api_predict` 接口。

这不是纯数据驱动的注册机制：上游要求用户在 `src/external_process` 下实现模型输入与输出处理器，然后重新构建并安装共享库。`image_to_vector`、`text_to_vector` 和 `image_classification` 则为随附示例提供图像与 SentencePiece 预处理。

### 重要对象

- `model_info`、`model_layer_info` 和 `base_model_info`：模型路径、元数据、层以及基础模型关系。
- `create_model`、`modify_model`、`drop_model`、`load_base_model` 和 `register_process`：模型生命周期与编译式回调注册。
- `predict_float`、`predict_text`、`predict_batch_float8` 和 `predict_batch_text`：标量与窗口式推理入口。
- `ai_operator`、`api_load_model` 和 `api_predict`：1.3.0 引入的任务式 API。
- `mvec`、`get_mvec_data`、`get_mvec_shape` 和 `text_to_mvec`：带形状的向量存储与转换。

### 安全、兼容与打包注意事项

模型和分词器路径由 PostgreSQL 服务进程打开，图像预处理接受由 OpenCV 在服务器环境中打开的路径。只能把这些函数授权给可信角色：获准调用者能够消耗大量 CPU/GPU 内存、访问服务器可达文件，并在数据库后端内运行模型代码。请在有价值集群之外验证模型来源、资源限制、并发、崩溃隔离和 GPU 行为。

复核的源码没有直接的 `pgdl--1.3.0.sql` 安装脚本，而是依赖从 1.0.0 开始的安装/升级链。1.1.0 升级中的多个向量函数引用 `MODULE_PATHNAME/pgdl.so`，但 `MODULE_PATHNAME` 本身已经解析为动态库路径；未经补丁的安装很可能得到无效库名。README 示例也与当前 SQL 签名不一致。保留任何数据前，必须针对实际 PostgreSQL 12 制品构建并运行 `CREATE EXTENSION`、完整升级链、向量往返测试和代表性模型。
