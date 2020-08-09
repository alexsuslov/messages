{{- define "debug"}}Debug message from err_aggr.tpl: [{{.Source.PID}}] some error err={{.Source.Err.Error}} filename={{.Filename}}:{{.Line}}
{{- end}}
