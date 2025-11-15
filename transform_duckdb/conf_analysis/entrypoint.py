#!/usr/bin/env python

import logging
import sys
from dbt.cli.main import dbtRunner, dbtRunnerResult


def main() -> None:
    logging.info("Starting scheduled dbt run...")

    try:
        dbt = dbtRunner()
        cli_args = ["run", "--target", "azure"]
        res: dbtRunnerResult = dbt.invoke(cli_args)

        logging.info("dbt run completed successfully.")

    except subprocess.CalledProcessError as e:
        logging.error("dbt run failed.")
        logging.error(e.stderr)
        raise e

if __name__ == "__main__":
    logging.basicConfig(stream=sys.stdout, level=logging.INFO, force=True)
    main()