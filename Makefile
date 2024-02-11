version = "v0.0.6"

release:
	git tag $(version)
	git push origin $(version)
