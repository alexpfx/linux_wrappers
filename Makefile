version = "v0.0.5"

release:
	git tag $(version)
	git push origin $(version)
