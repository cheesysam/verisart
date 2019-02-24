```docker build -t verisart .```

```docker run --rm -p 8000:8000 verisart```

```curl -d '{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "sam", "Year": 1990, "Note": ""}' -X POST http://localhost:8000/certificates/1```

```curl localhost:8000/users/sam/certificates```

```curl -X "DELETE" localhost:8000/test/1```

