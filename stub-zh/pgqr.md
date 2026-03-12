

## 用法

> [pgqr: PostgreSQL 二维码生成](https://github.com/AbdulYadi/pgqr)

### 函数

```sql
pgqr(t text, correction_level integer, model_number integer, scale integer) RETURNS bytea
```

参数：
- `t` -- 要编码到二维码中的文本
- `correction_level` -- 纠错级别：0 到 3
- `model_number` -- 二维码模型编号：0 到 2
- `scale` -- 每个点的像素数（缩放因子）

### 示例

生成单色位图格式的二维码：

```sql
SELECT pgqr('QR Code with PostgreSQL', 0, 0, 4);
```

输出为单色位图图像（BMP 格式），以 `bytea` 类型返回，可直接显示或存储。
