// +build ignore

package resources

func series_get(w http.ResponseWriter, r *http.Request) {

	result, err := dbQuery(
		dbConn,
		fmt.Sprintf("select * from gunpla where series = %s", mux.Vars(r)["series"]),
	)
	logError(err)

	message, err := json.Marshal(result)
	logError(err)
	w.Write(message)
}
