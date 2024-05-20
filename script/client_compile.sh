BASEDIR="$(cd "$(dirname "$0")" && pwd)"
PROTODIR="${BASEDIR}/../src/client"

echo $BASEDIR
echo $PROTODIR

cd ${PROTODIR}

go build -o client.so -buildmode=c-shared client.go