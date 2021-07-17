This is Nikhil Hulle, applying for the post of software dev 
i have finished the task except setting up the swagger docs, i was finding it a little challenging to set it up, but i have attempted to write documentation in the each files for the ease of understanding......

I have currently not uploaded this to heroku, and will update this readme if i do it, other than that i have completed all the requirements to the best of my ability as i had to learn everything from scratch since i am a fresher.

NOTE : FOR THE DATABASE I HAVE USED MONGODB ATLAS AND THIS REPO HAS THE CREDENTIALS TO THAT CLUSTER .

There are 2 APIs here
1) API_CovidSTATS
2) API_LocationStats

To execute the task, follow the below procedure

1) clone the repo
2) run both the api using the command run ```go run cmd/server/main.go```
3) The API_CovidSTATS api runs on port 1453.
4) The API_LocationStats api runs on port 2445.
5) For this task running the below uri in a browser tab does the work.
6) ```http://localhost:1453/location?lat=18.3817844&lng=77.11884189999999``` --->  this is an example call to that api.
7) You will get a screen with the required information



EXPLANATION:

API_CovidSTATS takes the inputs in the form of coordinates as shown in the example below and converts that to a location using the reversegeo coding api available from opencage, then using this converted locatiom, i make an internal api call to  API_LocationStats which does the work of updating the mongodb of the latest covid stats which i take from ```https://api.rootnet.in/covid19-in/stats/latest``` . After having done the storing/updation part in the db , using the location from the previous api , i query the database and return the covid statistics for india and the location pertaining to the coordinates entered by the user.





