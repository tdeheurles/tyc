# TYC

## Generate or update infrastructure
```bash
export AWS_ACCESS_KEY=...
export AWS_ACCESS_SECRET_KEY=...
./scripts/terraform.sh
```

## Update website content
### prerequisites
- install hugo
- install aws cli v2
### execute
```bash
export AWS_ACCESS_KEY=...
export AWS_ACCESS_SECRET_KEY=...

# will generate all public files for the website
./scripts/create_website.sh

# push files to s3
 (cd ./hugo/public && aws s3 sync . s3://the-youtube-synopsis/ --delete)
```
