resource "aws_s3_bucket" "webserver" {
  bucket = "the-youtube-synopsis"
}

resource "aws_s3_bucket_website_configuration" "example" {
  bucket = aws_s3_bucket.webserver.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

resource "aws_s3_bucket_acl" "webserver" {
  bucket = aws_s3_bucket.webserver.id
  acl    = "public-read"
}

resource "aws_s3_bucket_policy" "webserver" {
  bucket = aws_s3_bucket.webserver.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicReadGetObject",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::${aws_s3_bucket.webserver.bucket}/*"
    }
  ]
}
EOF
}
