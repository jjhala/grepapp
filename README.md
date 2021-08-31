# grepapp

Few examples on how to run this app:

## Exact match string
FILE_PATH="hamlet.txt" KEY_STRING="Fortinbras"  ./basicapps

## Case insensitive search
FILE_PATH="hamlet.txt" KEY_STRING="FORTINBRAS" IGNORE_CASE="true" ./basicapps

## Using regex
FILE_PATH="hamlet.txt" KEY_STRING="FORTIN+" REGEX="true" ./basicapps

## Using stdin (includes case insensitive search by default)
./basicapps hamlet.txt FORTINBRAS
