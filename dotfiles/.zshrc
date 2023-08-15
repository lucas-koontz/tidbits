# AWS
alias aws_profile="--profile <profile>"
alias ecr-login="aws ecr get-login-password --region us-east-1 --profile <profile> | docker login --username AWS --password-stdin https://<account>.dkr.ecr.us-east-1.amazonaws.com"

# Kubernetes Clusters Alias
alias k="kubectl" 
alias klist="kubectl config get-contexts"
alias kenv="kubectl --context='<cluster>'"
alias kapp="kenv -n <app>"

# Postgres
alias pg_start="launchctl load ~/Library/LaunchAgents"
alias pg_stop="launchctl unload ~/Library/LaunchAgents"
