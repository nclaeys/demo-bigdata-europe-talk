# dbt duckdb transformations

## Prerequisites
- Install python 3.8+

## Getting started
1. Create and activate a virtual environment

```bash
python -m venv venv
```

2. Activate the virtual environment
```bash
source venv/bin/activate
```

3. Install dependencies
```bash
pip install -r requirements.txt
```

4. Test setup by running dbt debug
```bash
dbt debug --project-dir conf_analysis --profiles-dir conf_analysis
```

5. Run all dbt models
```bash
dbt run --project-dir conf_analysis --profiles-dir conf_analysis
```

## Architecture diagram

## Production setup
To run the dbt transformations in production, we use an Azure Container App Job.
The necessary steps to deploy and run the job are:
1. Build and push the Docker image. The Dockerfile is located in the `conf_analysis` folder.
2. Deploy the Container App Job using the provided Terraform configuration.
3. Start the job using the Azure CLI OR schedule it using the built-in scheduling capabilities of Azure Container Jobs.

This setup works well if there are not too many dependencies between different jobs.
If you want to make more complex workflows, a good first step is using Github Actions to orchestrate multiple jobs.
When the amount of workflows increases significantly, consider using a more advanced orchestrator like Airflow.

### Starting the job manually
```bash
az containerapp job start --name <job-name> --resource-group <resource-group>
```