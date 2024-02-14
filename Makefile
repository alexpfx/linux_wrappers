version = "v0.0.9"

release:
	git tag $(version)
	git push origin $(version)
