[{{range $index,$val:=.}}{{ if eq $index 0}}{"Domain_id":"{{.Domain_id}}","Curve_desc":"{{.Curve_desc}}","As_of_date":"{{.As_of_date}}"{{with .Curve_yield}}{{range .}},"{{.Struct_code}}":"{{.Yield}}"{{end}}{{end}}}{{else}},{"Domain_id":"{{.Domain_id}}","Curve_desc":"{{.Curve_desc}}","As_of_date":"{{.As_of_date}}"{{with .Curve_yield}}{{range .}},"{{.Struct_code}}":"{{.Yield}}"{{end}}{{end}}}{{end}}{{end}}]