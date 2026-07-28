## Usage

Sources:

- [Official upstream README](https://github.com/postgres-artificialintelligence/pgai/blob/5c5f01d833ab4ee80779e0958c7ca0f6022e5c94/README.md)
- [Official extension control file (pgai.control)](https://github.com/postgres-artificialintelligence/pgai/blob/5c5f01d833ab4ee80779e0958c7ca0f6022e5c94/pgai.control)
- [Official extension SQL (pgai--1.0.sql)](https://github.com/postgres-artificialintelligence/pgai/blob/5c5f01d833ab4ee80779e0958c7ca0f6022e5c94/pgai--1.0.sql)

`pgai` — This guide provides step-by-step instructions to install PostgreSQL on various operating systems. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgai;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `fetch_apple_stock(input_date DATE)` is an extension function and returns `TABLE`.
- `fetch_google_stock(input_date DATE)` is an extension function and returns `TABLE`.
- `fetch_tesla_stock(input_date DATE)` is an extension function and returns `TABLE`.
- `get_arima_prediction(target_date TEXT)` is an extension function and returns `FLOAT`.
- `get_db_config()` is an extension function and returns `TABLE`.
- `get_prediction_for_date(prediction_date DATE)` is an extension function and returns `TABLE`.
- `predict_stock_close_value_apple(input_date_str TEXT)` is an extension function and returns `FLOAT`.
- `predict_stock_close_value_tesla(input_date_str TEXT)` is an extension function and returns `FLOAT`.
- `apple_stocks` is an extension-defined materialized view.
- `google_stocks` is an extension-defined materialized view.
- `tesla_stocks` is an extension-defined materialized view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
