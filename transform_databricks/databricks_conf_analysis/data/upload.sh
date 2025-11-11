#!/usr/bin/env bash
set -euo pipefail
STORAGE_ACCOUNT="databricksbigdataconf"
CONTAINER_NAME="data"
LOCAL_DIR="./"

STORAGE_KEY="${AZURE_STORAGE_KEY:-}"

if [ ! -d "$LOCAL_DIR" ]; then
  echo "Error: Directory '$LOCAL_DIR' does not exist."
  exit 1
fi


for file in "$LOCAL_DIR"/*.csv; do
  [ -e "$file" ] || { echo "No CSV files found."; exit 0; }
  filename=$(basename "$file")

  echo "Uploading $filename ..."
  az storage blob upload \
    --account-name "$STORAGE_ACCOUNT" \
    --container-name "$CONTAINER_NAME" \
    --name "raw/$filename" \
    --file "$file" \
    --overwrite \
    --auth-mode login \
    --only-show-errors
  echo "Uploaded: $filename"
done

echo
echo "All CSV files uploaded successfully to container '$CONTAINER_NAME'!"
