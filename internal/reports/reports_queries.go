package reports

const getDataQuery := "
	SELECT M.id, IPA.value, QT.value, 
	QU.value, SII.value, M.start_query_time, 
	M.finish_query_time FROM public.main AS M 
	LEFT JOIN public.ip_address AS IPA ON IPA.id = M.ip_id
	LEFT JOIN public.query_type AS QT ON QT.id = M.query_type_id
	LEFT JOIN public.query_url AS QU ON QU.id=M.query_url_id
	LEFT JOIN public.session_ids AS SII ON SII.id=M.session_ids_id"