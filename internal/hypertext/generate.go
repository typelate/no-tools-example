package hypertext

//go:generate muxt generate --use-receiver-type=Server
//go:generate counterfeiter -generate

//counterfeiter:generate -o=../fake/server.go        --fake-name=Server       . RoutesReceiver

//counterfeiter:generate -o=../fake/querier.go       --fake-name=Querier      ../database Querier
//counterfeiter:generate -o=../fake/db_connection.go --fake-name=DBConnection ../database Connection
//counterfeiter:generate -o=../fake/tx.go            --fake-name=Tx           github.com/jackc/pgx/v5.Tx
