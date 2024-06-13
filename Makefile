version = "v0.0.14"

release:
	git tag $(version)
	git push origin $(version)
