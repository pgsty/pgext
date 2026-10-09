## 用法

来源：

- [README.md](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/README.md)
- [tensor_dancer/sql/tensor_dancer--0.1.0.sql](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/sql/tensor_dancer--0.1.0.sql)
- [tensor_dancer/copilota.cpp](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/copilota.cpp)
- [tensor_dancer/meson.build](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/meson.build)
- [tensor_dancer/tensor_dancer.control](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/tensor_dancer/tensor_dancer.control)

`tensor_dancer` 是 Tensor Dancer 的早期原型。固定源码版本的 SQL 提供返回文本的 `copilota(content text)`，实现会将提示内容发送给本机 Ollama 端点。

### 核心用法

```sql
CREATE EXTENSION tensor_dancer;
SELECT oid::regprocedure FROM pg_proc WHERE proname = 'copilota';
```

### 运行边界

端点硬编码为数据库服务器视角下的 `http://localhost:11434`。同仓库的 `vector_extra` 有独立的控制文件身份，不会自动安装。构建需要 BLAS、LAPACK、GGML 及 HTTP/JSON 依赖；尚未核实 PostgreSQL 主版本支持范围。SQL 虽将函数声明为 `IMMUTABLE` 和 `PARALLEL SAFE`，实现却会发起外部请求，这些声明不能保证网络结果不变，也不能使请求随事务回滚。应将其视为实验功能，保护提示内容，避免用于索引表达式或假定重复调用没有副作用。
