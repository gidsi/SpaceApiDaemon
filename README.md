SpaceApiDaemon
==============

**Work in progress, do not use in production!**

First draft of the SpaceApiDaemon, a tool that provides your space with an easy to use, valid, SpaceApi.

Goal is to give more people/systems access to change values in your SpaceApi easily (example: Your new temperature sensor which is hooked to an esp gets his own access token and can update temperature readings etc. by itself).

How to start
------------
1. Install [docker](https://docs.docker.com/engine/installation/) and [docker-compose](https://docs.docker.com/compose/install/).
2. Create data and config folder
```bash
# mkdir -p /opt/spaceapidaemon/{db,config}
```
3. Clone the repo and switch into the project folder.
4. Create the config file
```bash
# cp config_example.yaml /opt/spaceapidaemon/config/
```
5. run:
```
# docker-compose up -d
# docker-compose logs daemon
```
6. Copy your token, e.g. the part between the curly braces showing up in the log, it should looks like this. 
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MDE3OTQzNzN9.qOW5dnW-HtXm9GotAOk9EJ2lDj-qTkdt3kL9EAbc3IE
```
7. Open a browser and open http://yourdomain/edit, and log-in (Button Authorize). You should at least provide data to basedata, location, state, contact and issue_report_channels. 

8. Your SpaceApi is now running at http://yourdomain/api/

9. Profit!

Todo:
-----
* ~~add docker / docker-compose file~~
* ~~add an deployment example and running example~~
* ~~add a better first run experience~~
* ~~add routes for most used values~~
* ~~force user to add a config file / signing key~~
* ~~add post / delete route for access tokens~~
* find a place to live for the swagger.json
* add config file documentation
* add ssl documentation