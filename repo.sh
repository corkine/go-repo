#export GOROOT=/usr/local/go
export GO111MODULE=on

cd /go-repo

git pull

kill $(ps axu | grep "/tmp/go.*/exe/repo.*" | grep -v grep | awk '{print $2}')
nohup go run repo.go 1>>/var/log/go_repo.log 2>&1 &
echo "goRepo run on Port \
$(ps aux | grep "go run.*repo.go" | grep -v grep | awk '{print $2}')"

echo "========================================================================="
echo ""
echo "server log here, press ctrl + c to exit (the server running normally) "
echo ""
echo "========================================================================="