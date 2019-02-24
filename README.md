# docker build and run

```docker build -t verisart .```

```docker run --rm -p 8000:8000 verisart```

# create a cert
```curl -d '{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "sam", "Year": 1990, "Note": ""}' -X POST http://localhost:8000/certificates/1```

# list certs
```curl localhost:8000/users/sam/certificates```

# delete a cert
```curl -X "DELETE" localhost:8000/certificates/1```

# notes

- /certificates route accepts POST, DELETE, PATCH
- /users route accepts GET
