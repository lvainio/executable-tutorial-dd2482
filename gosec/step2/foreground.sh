clear
echo "Installing necessary packages for go programs..."
cd ./unsafe_sql_concat
go get github.com/mattn/go-sqlite3
cd ../
echo DONE