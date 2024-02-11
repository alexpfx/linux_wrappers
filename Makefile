version = "v0.0.8"

release:
	git tag $(version)
	git push origin $(version)
