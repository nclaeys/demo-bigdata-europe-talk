# Demo Bigdata Europe 2025 

## Duckdb investigation for Delta

- Currently only read support for delta lake
- add write support using custom plugin:
  - https://github.com/l-mds/demo-dbt-duckdb-delta-plugin
  - https://github.com/milicevica23/dbt-duckdb/blob/feature/support-delta-plugin-write/dbt/adapters/duckdb/plugins/delta.py