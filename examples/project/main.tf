terraform {
    required_providers {
        circleci = {
            version = "0.0.1"
            source = "github.com/olukotun-ts/circleci"
        }
    }
}

resource "circleci_project" "demo" {
    slug = "gh/olukotun-ts/name-button"
    branch = "master"
}

output "project_slug" {
    value = circleci_project.demo.slug
}