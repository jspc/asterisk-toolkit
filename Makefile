README.md: *.go
	goreadme -badge-godoc -badge-goreportcard -functions -methods -types -title "Asterisk Toolkit" > $@
