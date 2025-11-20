# Demo Bigdata Europe 2025 

This repository contains a practical demonstration of my talk at BigData Conference Europe 2025.
The focus is on showing how to use DuckDB as well as Databricks on conference data.

Using the conference data I will demonstrate the local development flow as well as how to productionize the solution for both setups.

## Databricks setup
Classic Databricks setup for running dbt models for local development as well as scheduling them in production.
The infrastructure is setup for Databricks on Azure is out of scope for this demo.

The raw data is stored in ADLS Gen2 storage account and the relevant tables are created in Unity Catalog.
For more details about the setup please refer to the [Databricks README](transform_databricks/README.md)

## DuckDB setup
DuckDB setup for running dbt models for local development as well as scheduling them in production.
The raw data is stored as CSV files locally (in `data` folder) as well as in ADLS Gen2 storage account.

For more details about the setup please refer to the [DuckDB README](transform_duckdb/README.md)

### Duckdb usage Delta

- Currently only read support for delta lake
- add write support using custom plugin:
  - https://github.com/l-mds/demo-dbt-duckdb-delta-plugin
  - https://github.com/milicevica23/dbt-duckdb/blob/feature/support-delta-plugin-write/dbt/adapters/duckdb/plugins/delta.py

### Duckdb azure blob integration

There are 2 plugins to interact with azure blob storage:
- fsspec and adlfs: which is used in this demo and supports reading and writing on blob storage when using Python but has limited support for authentication.
- [Azure extension](https://duckdb.org/docs/stable/core_extensions/azure) of Duckdb. This only supports reading from Azure Blob storage but has better secrets integration.
  The writing support is currently in progress, see the following [PR](https://github.com/duckdb/duckdb-azure/pull/131)

For more details about how to use both plugins, take a look at [the following blogpost](https://medium.com/datamindedbe/quacking-queries-in-the-azure-cloud-with-duckdb-14be50f6e141).