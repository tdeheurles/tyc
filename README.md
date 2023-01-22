# TYS - The Youtube Synopsis

## DEV
### prerequisites
- install [hugo](https://gohugo.io/)
### Start the webserver locally

```bash
./scripts/build_and_run.sh
```
Should result in something like
```
3 channels to build

channel Kurzegesagt - 3 videos
- aeWyp2vXxqA - Black Hole Star – The Star That Shouldn't Exist
- HpcTJW4ur54 - How To Terraform Mars - WITH LASERS
- Qsbe1pD8ocE - The Horror of the Slaver Ant
channel Andrew Huberman Lab - 3 videos
- tLS6t3FVOTI - Developing a Rational Approach to Supplementation for Health & Performance
- -wIt_WsJGfw - #105 - Dr. Sam Harris: Using Meditation to Focus, View Consciousness & Expand Your Mind
- LTGGyQS1fZE - #98 - Science-Based Tools for Increasing Happiness
channel Better Than Yesterday - 3 videos
- PXWFlea0l5w - It only takes 1 month to change your life
- TW2Kr0l2kGg - Making $10,000 per month has never been easier
- 11mVD2rPtIM - What it really takes to become successful
Start building sites …
hugo v0.109.0-47b12b83e636224e5e601813ff3e6790c191e371 linux/amd64 BuildDate=2022-12-23T10:38:11Z VendorInfo=gohugoio

                   | EN
-------------------+-----
  Pages            | 19
  Paginator pages  |  0
  Non-page files   |  0
  Static files     |  0
  Processed images |  0
  Aliases          |  0
  Sitemaps         |  1
  Cleaned          |  0

Built in 4 ms
Watching for changes in /home/asura/repositories/tdeheurles/tys/hugo/{archetypes,assets,content,data,layouts,static}
Watching for config changes in /home/asura/repositories/tdeheurles/tys/hugo/config.toml
Environment: "development"
Serving pages from memory
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
Press Ctrl+C to stop
```

Now go to http://localhost:1313/ and you should see the website


## OPS
### Define your AWS secrets
Create a .aws.secret file at the root of the project, and store your aws credentials
```bash
cat <<EOF > .aws.secret
export AWS_ACCESS_KEY=xxxxx
export AWS_ACCESS_SECRET_KEY=xxxxxxxxxxxxxx
EOF
```
### Generate or update infrastructure
#### prerequisites
- install [hugo](https://gohugo.io/)
- install [terraform](https://www.terraform.io/)
#### exec
```bash
./scripts/terraform.sh
```

### Update website content
#### prerequisites
- install [hugo](https://gohugo.io/)
- install [aws cli v2](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
#### execute
```bash
./scripts/update_website.sh
```
