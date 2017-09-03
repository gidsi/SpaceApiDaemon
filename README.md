SpaceApiDaemon deployment
=========================

**Work in progress, do not use in production!**

First draft of the SpaceApiDaemon, a tool that provides your space with an easy to use, valid, SpaceApi.

Goal is to give more people/systems access to change values in your SpaceApi easily (example: Your new temperature sensor which is hooked to an esp gets his own access token and can update temperature readings etc. by itself).

It will also provide a history for e.g. your opening times.

How to start
------------
1. Install [docker](https://docs.docker.com/engine/installation/) and [docker-compose](https://docs.docker.com/compose/install/).
2. Change the enviroment variables in the .evn file according to your needs. (Set the SPACEAPIURL to localhost if you gonna use a reverse proxy)
3. Create data and config folder
```bash
# mkdir -p /opt/spaceapidaemon/{db,config,certificates}
```
4. Clone the repo and switch into the project folder.
5. Create the [config file](https://github.com/gidsi/SpaceApiDaemon/tree/master/spaceApiDaemon#config)
```bash
# cp config_example.yaml /opt/spaceapidaemon/config/config.yaml
```
6a. (reverse proxy / testing) If you're already running a webserver, start it with:
```bash
# docker-compose up -d
```
This will open port 2015 (default caddy port), you have to setup your webserver to reverse proxy to that port ([nginx](https://www.nginx.com/resources/admin-guide/reverse-proxy/), [apache](https://httpd.apache.org/docs/2.4/howto/reverse_proxy.html)). You have to take care of ssl etc by yourself.

6b. (standalone) If you're not running a webserver use the following command to start:
```bash
# docker-compose -f standalone.yml up -d
```
Caddy will take care of getting and update SSL certificates from Lets Encrypt.  

7. Copy your token, e.g. the part between the curly braces showing up in the log, it should looks like this. 
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MDE3OTQzNzN9.qOW5dnW-HtXm9GotAOk9EJ2lDj-qTkdt3kL9EAbc3IE
```
8. Open a browser and open https://yourdomain/edit/, and log-in (Button Authorize). You should at least provide data to basedata, location, state, contact and issue_report_channels. 

9. Your SpaceApi is now running at https://yourdomain/api/

10. Profit!
