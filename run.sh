#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: ./run.sh <folder_name>"
  exit 1
fi

folder_name="$1"

# Create the folder
mkdir tutorial/"$folder_name"

# Create and populate files
touch "$folder_name/main.go"
touch "$folder_name/logic.md"
touch "$folder_name/io.md"

echo "Folder '$folder_name' created and initialized with files."
