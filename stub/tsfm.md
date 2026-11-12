## Usage

Sources:

- [Official README](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/README.md)
- [Control file](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/extension/tsfm.control)
- [Installation SQL](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/extension/tsfm--0.1.sql)
- [Settings implementation](https://github.com/barddoo/tsfm-pg/blob/eec8098d7f3cbda55cba1a0d207d4daa4741ad02/src/guc.zig)

`tsfm` runs local time-series forecasting models inside PostgreSQL 16 or 17. The SQL/control version is 0.1. Chronos-2 supports univariate input; Timestral also accepts multiple variates and covariates. An administrator must install the library and provision model files on the database host. SQL calls do not download models or send series to an external service.

### Core Workflow

After provisioning a supported checkpoint, install as a superuser and check the model before forecasting:

```sql
CREATE EXTENSION tsfm;
SELECT tsfm_version();
SELECT jsonb_pretty(tsfm_model_status());
SELECT tsfm_load();

SELECT step, quantile, value
FROM tsfm_forecast(ARRAY[10,11,12,13,14]::float8[], 3)
WHERE quantile = 0.5
ORDER BY step;
```

History must be ordered oldest to newest. Steps and series numbers start at 1. The default model directory is beneath PostgreSQL's shared-data directory. `tsfm.model_dir` and `tsfm.model` are fixed after a backend loads its engine; reconnect to change them. No shared preload is required.

### Privileges and Settings

Installation revokes all function privileges from `PUBLIC`. Grant only the functions an application needs:

```sql
GRANT EXECUTE ON FUNCTION tsfm_forecast(double precision[], integer) TO analyst;
```

The wrapper functions for model information and status also call underlying text helpers; callers need the corresponding helper privileges. `tsfm_forecast_ts` similarly requires permission on its underlying forecast function.

`tsfm.model_dir`, `tsfm.model`, `tsfm.threads`, `tsfm.enabled` and `tsfm.stub` are superuser settings. `tsfm.threads` defaults to 1 and sizes the worker pool when loading; many connections with multiple workers can oversubscribe the host. `tsfm.stub` selects a test predictor, not a trained model. `tsfm.enabled = off` blocks forecasting, warmup and quantile access.

The decoding settings include `tsfm.max_context`, `tsfm.sort_quantiles` and Timestral-specific controls. Quote the keyword when configuring `tsfm."symmetric"`:

```sql
SET tsfm."symmetric" = off;
```

### Function Families

| Function | Purpose |
| --- | --- |
| `tsfm_model_info()` | Report the model and actual engine configuration |
| `tsfm_model_status()` | Inspect checkpoint availability without loading weights |
| `tsfm_quantiles()` | Return model quantile levels |
| `tsfm_forecast_ex()` | Forecast multiple variates with optional covariates |
| `tsfm_forecast_batch()` | Process independent series in batches |
| `tsfm_forecast_matrix()` | Return a step-by-quantile array |
| `tsfm_forecast_matrix_ex()` | Return a variate-by-step-by-quantile array |
| `tsfm_forecast_ts()` | Attach regular future timestamps |

Extended histories flatten variates × time in row-major order; batches flatten series × variates × time. Past covariates cover the input history; known covariates also cover the horizon. Chronos-2 accepts one variate and no covariates. Timestamp forecasting assumes regular spacing and does not detect calendar gaps.

### Resource and Model Boundaries

Values are converted to float32. Missing elements and NaNs are treated as missing data; infinities and out-of-range values also become missing. Large integers can lose precision. The documented horizon limit is 1024, and one result is capped at 16 million float32 values.

Inference runs inside the backend. Cancellation is checked between calls or batch chunks, but an individual forward pass cannot be interrupted. Filter and bound workloads before invoking it. Checkpoint pages are memory-mapped; model state and scratch allocations can outlive a transaction.

GPU support is optional; the default build uses the CPU. The extension's Apache-2.0 license does not grant rights to every model: Timestral weights have separate terms and are not redistributed by the project. A managed service must explicitly support the custom library and model filesystem access before this extension can be used there.
