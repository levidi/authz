package security

import input.attributes.request.http as http_request

default allow = false

allow {
    is_token_valid
    action_allowed
}

token = encoded {
    [_, encoded] := split(http_request.headers.authorization, " ")
}

decoded_token = payload {
    [_, payload, _] := io.jwt.decode(token)
}

scopes = splitted_scopes {
    splitted_scopes := split(decoded_token.scope, " ")
}

is_token_valid {
    secret := "123"
    io.jwt.verify_hs256(token, secret)
    now := time.now_ns() / 1000000000
    now < decoded_token.exp
    decoded_token.aud[_] == "balance-service"
}

action_allowed {
    http_request.method == "GET"
    glob.match("/balances*", [], http_request.path)
    scopes[_] == "read"
}

action_allowed {
    http_request.method == "PUT"
    glob.match("/balances*", [], http_request.path)
    scopes[_] == "write"
}

