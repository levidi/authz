package security

default allow = false

allow {
    input.profile == data.rules[_].profile[_]
}
