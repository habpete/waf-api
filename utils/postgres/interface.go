package pg_client

type IPgClient interface {
	QUERY(queryString string) (interface{}, error)
}
