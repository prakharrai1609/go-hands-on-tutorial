#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: ./run.sh <folder_name>"
  exit 1
fi

folder_name="$1"
folder_name_capitalized="$(tr '[:lower:]' '[:upper:]' <<< ${folder_name:0:1})${folder_name:1}"

# Create the folder
mkdir tutorial/"$folder_name"

# Create and populate files
touch "tutorial/$folder_name/main.go"
touch "tutorial/$folder_name/logic.md"

echo "Folder 'tutorial/$folder_name' created and initialized with files."

# Update the README
readme_file="README.md"

cat << EOF >> "$readme_file"

. [$folder_name_capitalized](tutorial/$folder_name)

- [Code](tutorial/$folder_name/main.go)
- [Explanation](tutorial/$folder_name/logic.md)
EOF

echo "README updated."