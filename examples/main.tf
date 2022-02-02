terraform {
    required_providers {
        circleci = {
            version = "0.0.1"
            source = "github.com/olukotun-ts/circleci"
        }
    }
}

resource "circleci_project" "demo" {
    # name = "demo"
    # slug = "gh/olukotun-ts/name-button"
    # branch = "master"
}

# output "demo" {
#     value = circleci_project
# }