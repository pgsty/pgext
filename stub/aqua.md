## Usage

Sources:

- [Official upstream README](https://github.com/zjudbxai/peps/blob/1923b7e389a8028dc71cfcfaa2bddd8b491c62c4/README.md)
- [Official extension control file (aqua.control)](https://github.com/zjudbxai/peps/blob/1923b7e389a8028dc71cfcfaa2bddd8b491c62c4/aqua_extension/aqua.control)
- [Official extension SQL (aqua--0.0.1.sql)](https://github.com/zjudbxai/peps/blob/1923b7e389a8028dc71cfcfaa2bddd8b491c62c4/aqua_extension/sql/aqua--0.0.1.sql)

`aqua` — In-database neural user-defined functions for machine-learning inference. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION aqua;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `argmax(real[])` is an extension function and returns `integer`.
- `array_2d_agg_combine(internal, internal)` is an extension function and returns `internal`.
- `array_2d_agg_final_array(internal)` is an extension function and returns `real[]`.
- `array_2d_agg_state(internal, real[][])` is an extension function and returns `internal`.
- `array_agg_array_deserialize(bytea, internal)` is an extension function and returns `internal`.
- `array_agg_array_serialize(internal)` is an extension function and returns `bytea`.
- `avgpool(input_f real[][], kernel INT = 2, padding INT = 0, stride INT = 2, op_cost REAL = 0.0)` is an extension function and returns `real[]`.
- `avgpool_conv(kmatrix real[][], fmatrix real[][], bias real[], kernel INT = 2, padding INT = 0, stride INT = 2)` is an extension function and returns `real[]`.
- `batchnorm(fmatrix real[][], args real[])` is an extension function and returns `real[]`.
- `check_weight_pool()` is an extension function and returns `VOID`.
- `concat_array(num int, input1 REAL[][] DEFAULT NULL, input2 REAL[][] DEFAULT NULL, input3 REAL[][] DEFAULT NULL, input4 REAL[][] DEFAULT NULL, input5 REAL[][] DEFAULT NULL, input6 REAL[][] DEFAULT NULL, input1_offset INT [] DEFAULT NULL)` is an extension function and returns `real[]`.
- `conv_text(kmatrix text, fmatrix real[][], kernel_H INT, kernel_W INT, padding_H INT, padding_W INT, stride INT)` is an extension function and returns `real[]`.
- `group_kfm_im2col(kmatrix text, fmatrix real[][][], kernel INT = 3, padding INT = 1, stride INT = 1, groups INT = 8)` is an extension function and returns `real[]`.
- `im2col(input_f real[][], kernel INT = 3, padding INT = 1, stride INT = 1)` is an extension function and returns `real[]`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
