
mpush:
	git add .;
	git commit -m "${MSG}";
	git push -u origin main
push:
	git tag "${TAG}";
	git push origin "${TAG}"
