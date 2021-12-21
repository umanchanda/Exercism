import secrets

def private_key(p):
    if secrets.randbelow(p) > p or p < 1:
        raise Exception("private key must be less than p")
    else:
        return secrets.randbelow(p)


def public_key(p, g, private):
    return (g**private) % p


def secret(p, public, private):
    return (public**private) % p
