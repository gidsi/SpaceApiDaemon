SpaceApiDaemon
==============

First draft of the SpaceApiDaemon, a tool that provides your space with an easy to use SpaceApi.

Goal is to give more people/systems access to change values in your SpaceApi easily (example: Your new temperature sensor which is hooked to an esp gets his own access token and can update temperature readings etc. by itself).

Gets all messy in golang right now, might rewrite it, or not, not sure at this moment.

Todo:
-----
* add docker / docker-compose file
* add an deployment and running example
* add routes for most used values