
# Serve mkdocs
serve:
	mkdocs serve --watch ./site --config-file ./site/mkdocs.yml

# Publish mkdocs
publish:
	mkdocs gh-deploy --force --config-file ./site/mkdocs.yml  --clean
