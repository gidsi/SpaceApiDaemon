SpaceApiDaemon
==============

This daemon gives you an api to manage and control your [SpaceApi](http://spaceapi.net/)

config
------
You can either provide a config file at /etc/spaceapidaemon/config.yml ([example](https://github.com/gidsi/SpaceApiDaemon/blob/master/config_example.yaml)) or apply corresponding flags to the executable
```
--apiPort
    Sets the port the api open (default: 8080)
--signingKey
    Signing key for access token, has to be provided
--mongoConnectionString
    Connection string to the mongodb, documentation can be found [here](http://api.mongodb.com/java/current/com/mongodb/ConnectionString.html) (default: localhost). 
--mongoCollection
    Collection to be used (default: spaceApi)
--allowedOrigins
    Comma seperated list of allowed origins (default: http://localhost)
```

example call
------------
The following call will set your spaces state to open
```bash
curl -X PUT "http://YOURDOMAIN/state?Authorization=YOURTOKEN" -H  "accept: application/json" -H  "content-type: application/json" -d "{  \"open\": true }"
```

clients
-------
You can generate clients for a bunch of languages under [editor.swagger.io](http://editor.swagger.io/), just import [this](https://raw.githubusercontent.com/gidsi/SpaceApiDaemon/master/spaceApiDaemon/main/swagger.json) (keep in mind, this is the latest version of the file, if you haven an older installation you should use the one provided in your api). 