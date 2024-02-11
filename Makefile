version = "v0.0.7"

release:
	git tag $(version)
	git push origin $(version)
