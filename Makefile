build: 
	go build ./...

dev: 
	dev_appserver.py ./app/app.yaml
	
prepare: 
	go mod download

deploy:
	gcloud app deploy --project $PROJECT_ID ./app/app.yaml
