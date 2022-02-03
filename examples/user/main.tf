terraform {
    required_providers {
        circleci = {
            version = "0.0.1"
            source = "github.com/olukotun-ts/circleci"
        }
    }
}

resource "circleci_user" "demo" {
    organization = "olukotun-ts"
    vcs_provider = "github"
    projects = [
        "name-button",
		"confluent-kafka-go", 
		"circleci-demo", 
		"circleci-demo-ruby-rails"
    ]
    branch = "master"
}

output "followed_projects" {
    value = circleci_user.demo.projects
}
