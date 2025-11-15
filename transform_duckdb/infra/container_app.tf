locals {
  rg_name = "dm-playground-rg"
  location = "Germany West Central"
}

resource "azurerm_log_analytics_workspace" "workspace" {
  name                = "dbt-logs-workspace"
  location            = local.location
  resource_group_name = local.rg_name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

resource "azurerm_container_app_environment" "dbt-environment" {
  name                       = "dbt-dev-environment"
  location                   = local.location
  resource_group_name        = local.rg_name
  logs_destination            = "azure-monitor"
}

resource "azurerm_container_app_job" "conf-job" {
  name                         = "dbt-conf-job"
  location                     = local.location
  resource_group_name          = local.rg_name
  container_app_environment_id = azurerm_container_app_environment.dbt-environment.id
  identity {
    type = "SystemAssigned"
  }
  replica_timeout_in_seconds = 600
  replica_retry_limit        = 10

  manual_trigger_config {
    parallelism              = 1
    replica_completion_count = 1
  }

  # schedule_trigger_config {
  #   cron_expression = "0 1 * * *"
  #   parallelism = 1
  # }

  template {
    container {
      image = "nilli9990/demo-dbt-conf-analysis:v2"
      name  = "dbt"
      cpu    = 1
      memory = "2Gi"
      env {
        name = "DBT_ENV_SECRET_STORAGE_ACCOUNT_KEY"
        secret_name = "storage-account-key"
      }
    }

  }
  # You can also refer to keyvault for this
  secret {
    name = "storage-account-key"
    value = "<FILL_STORAGE_ACCOUNT_KEY>"
  }
}