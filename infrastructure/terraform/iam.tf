
resource "aws_iam_group" "tys" {
  name = "the-youtube-synopsis"

}
resource "aws_iam_user" "githubaction" {
  name = "tys-githubaction"
}

resource "aws_iam_policy" "update_website" {
  name = "tys-githubaction_policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListBucket",
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::${aws_s3_bucket.webserver.bucket}",
        "arn:aws:s3:::${aws_s3_bucket.webserver.bucket}/*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_group_membership" "tys-githubaction-membership" {
  name = "tys-githubaction-membership"

  users = ["${aws_iam_user.githubaction.name}"]
  group = aws_iam_group.tys.name
}

#resource "aws_iam_user_policy_attachment" "githubaction_update_website" {
# user       = aws_iam_user.githubaction.id
#policy_arn = aws_iam_user_policy.update_website.arn
#}