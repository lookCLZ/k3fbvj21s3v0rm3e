all: run pyrun
run:
	cd /Users/liuhongrui/heima/contacts/python/2019.12/py3-web/app && npm run start & > app.log
pyrun:
	cd /Users/liuhongrui/heima/contacts/python/2019.12/py3-web/www && python3 app.py > serve.log
	