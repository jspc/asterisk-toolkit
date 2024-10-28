README.md: *.go
	goreadme -badge-godoc -badge-goreportcard -functions -methods -types -title "Asterisk Toolkit" -import-path github.com/jspc/asterisk-toolkit > $@
