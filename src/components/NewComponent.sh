# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

read -p 'New Folder Name: ' foldervar

read -p 'New Component name: ' uservar
echo

command git clone https://source.developers.google.com/p/future-309012/r/template-go-component

# Rename root folder
command mv template-go-component $foldervar

# Delete all git stuff to prevent commits to template repo
command cd $foldervar
command rm .gitignore
command rm -rf .git

command cd v1

# Rename main component
FILE_EXT=".go"
COMP_NAME=$foldervar$FILE_EXT
command mv component.go "$COMP_NAME"

filename="$COMP_NAME"
search="NameManager"
replace=$uservar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="$COMP_NAME"
search="*NameManager"
replace="*"$uservar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="$COMP_NAME"
search="template_go_component"
replace=$foldervar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="init.go"
replace=$foldervar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="flags.go"
replace=$foldervar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="utils.go"
replace=$foldervar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="state.go"
replace=$foldervar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

filename="deps.go"
replace=$foldervar
if [[ $search != "" && $replace != "" ]]; then
  sed -i "s/$search/$replace/" $filename
fi

# Update all build files & dependencies
command bazel run //:gazelle

# Build all sources
command bazel build //:build

command echo
command echo "DONE! Build NEW Component!"
command echo "Inspect and run git add ."
command echo
