version = "v0.0.20"

release:
	git tag $(version)
	git push origin $(version)
