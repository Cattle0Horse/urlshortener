(
	gomodifytags -all
		-add-tags json,yaml,mapstructure
		-w
		-file "config/config.go"
)