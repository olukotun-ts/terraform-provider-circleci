terraform {
    required_providers {
        circleci = {
            version = "0.0.1"
            source = "github.com/olukotun-ts/circleci"
        }
    }
}

resource "circleci_project" "demo" {
    # slug = "gh/olukotun-ts/name-button"
}

# output "demo" {
#     value = circleci_project
# }