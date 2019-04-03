export PATH=$PATH:"/usr/local/go/bin"
export GOPATH="$HOME/go_workspace"
export PATH=~/go_workspace/bin:$PATH
export MYSQL_SERVER="localhost"
export MYSQL_PORT="3306"
export MYSQL_USER="root"
export MYSQL_PASS="password"
export DB_NAME="crosssell"
#remove comment from dep init if gopkg.toml there
#dep init
dep ensure
go run $HOME/go_workspace/src/cross-sell/*.go
