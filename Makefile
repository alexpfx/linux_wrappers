version = "v0.0.4"

release:
	git tag $(version)
	git push origin $(version)
