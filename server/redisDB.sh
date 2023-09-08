start_test_db() {
  docker run --rm -it --name redis-server -p 6379:6379 -e DB_NAME=redis -e REDIS_URL=redis://localhost:6379 -d redis --loglevel warning
}

stop_test_db() {
  docker stop redis-server
}

while getopts t: flag; do
  case "${flag}" in
    t) type=${OPTARG} ;;
  esac
done

case $type in
  start)
    start_test_db
    echo "Test database started."
    ;;
  unit)
    start_test_db
    # CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out -run Unit ./...
    stop_test_db
    echo "Test database stopped."
    ;;
  integration)
    start_test_db
    # CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out -run Integration ./...
    stop_test_db
    echo "Test database stopped."
    ;;
  *)
    start_test_db
    # CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out ./...
    stop_test_db
    echo "Test database stopped."
    ;;
esac
