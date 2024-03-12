version = "v0.0.11"

release:
	git tag $(version)
	git push origin $(version)
