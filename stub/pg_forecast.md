## Usage

Sources:

- [README.md](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/README.md)
- [src/pg_forecast/Makefile](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/src/pg_forecast/Makefile)
- [src/pg_forecast/sql/functions/_common.sql](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/src/pg_forecast/sql/functions/_common.sql)
- [src/pg_forecast/pg_forecast.control](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/src/pg_forecast/pg_forecast.control)

`pg_forecast` fits ARIMA models, forecasts time-series values and updates model state through triggers. This research implementation links to libnlopt; the SQL version 1.0 is generated from the source SQL fragments by its Makefile.

### Core Workflow

```sql
CREATE EXTENSION pg_forecast;
CREATE TABLE time_series_data (timestamp timestamp, value double precision);
```

Populate the source table with representative ordered time-series observations before calling the training and forecasting functions.

```sql
SELECT create_forecast('autoarima', 'time_series_data', 'timestamp', 'value');
SELECT * FROM run_forecast('autoarima', 'time_series_data', 'timestamp', 'value', 10, '1 hour');
```

### Operational Boundaries

`create_forecast()` trains a named model for a source table, date column and value column. `run_forecast()` returns forecast dates and values for a horizon and interval; `remove_forecast()` removes the registered model. The implemented model is autoarima. Model tables and insertion triggers are persistent database objects. Several routines use SECURITY DEFINER and the public schema: deploy only in a database whose object-creation privileges and function access are controlled. Installation uses the control-file superuser default. Upstream gives a PostgreSQL 18 installation example but no full compatibility matrix or clear license declaration.
