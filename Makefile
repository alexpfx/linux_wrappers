version = "v0.0.10"

release:
	git tag $(version)
	git push origin $(version)
