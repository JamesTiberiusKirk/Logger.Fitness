# Logger.Fitness 31/12/2021
This idea originally came to me from my honours project which I had in my last
year of University. The idea was a two-parter, to make a workout logger
application and to make the app in a micro services approach. This proved to be
a bit of a time waste because I did not end up properly finishing the application.
So here I am with my second run at this idea.

## The Idea
The idea is simple, make a fuss-free workout logging application which is
customisable in the exercises that you have and the data that you have available.
Then be able to take logs and create custom historical charts based cross
referencing data that the user wants.


The user would be able to add tags to things such as exercises, or
workouts. Then, the user would be able to create a graph where it would be able
to pull in data which has that specific tag on it and show progress over time
with or without that tag.


For example, the user could tag a specific workout with tag "super tired". Then 
when compiling a graph the user could exclude that specific tag from the search.
This logic could be extended to include tags, and multiple exercises types.


However, to even think about starting the sort of data and tag magic which has 
been described above, a base logger would first have to be made, or an MVP.

### v1.0 MVP
This MVP is all about making a functional logger where you would be able to just
simply create custom exercise types, with two types of data (single values and
sets). Then be able to start workouts, log the exercises, log notes and stop the
workouts. 

### v2.0 UX-boogaloo
The main purpose of this version would be to focus a bit on how the app looks, 
the user flow between different pages of the app, and giving the app a bit of
an identity. 

Also, historical and analytical features for the app such as simple "1 rep max"
charts, and more will be introduced. This is where the system will also lay the
groundwork's for the tagging system and the flexible chart system will be
created.

## Introducing The Tech Stack
Everything is containerised with docker images, backend consisting of a Golang 
with Echo REST framework application and a MongoDB database.

The frontend is served from a separate Nginx container which also forwards all 
the API traffic to the backend container. The frontend its self is just a fairly
simple VueJS using Vuex state management, Vue router and Bootstrap for a UI
toolkit.

Everything is sitting behind an [Nginx Proxy Manager](https://nginxproxymanager.com/)
simply because of its simplicity and ease of use. This will most likely be
swapped for a manual configuration of Nginx.

## Current Progress
- Login and authentication system. (just works enough for now)
- Creation, addition and deletion of custom user exercises with either a set data type or a single value data type.
  - Example of a set for reference `Set{resistance int, repetitions int}`
- Starting, stopping, editing and deletion of workouts
- Addition of exercises to the workouts based on the custom exercise types.
- Addition and deletion of sets to exercises inside workouts.

## Next Steps
When it comes to the backend, most of my authentication system will be scrapped
in favour of some OAuth system, at the moment thinking of something like
[KeyCloak](https://www.keycloak.org/), ideally I would like to own the database
and not offload it to firebase or anything like that.
A bit or backend cleanup and test coverage would also be needed.

When it comes to the UI, there is a lot of work to be done. This would be done
with Ionic and Vuejs. However, the most important part would be some actual 
designs, will possibly contact one of my designer friends for UX and wireframes.
(or I might just swallow my pride and learn to do it myself).

### Functionality 
Up next in terms of functionality, the addition of workout templates will be
implemented. Meaning when starting a workout the user would be able to chose
from some user defined templates which would just have all the exercises needed
for that workout.

Another crucial piece of functionality is the historical and analytical part of
it. There will be one separate page to show charts for specific exercises, one 
for the custom user exercises. And information bits sprinkled thought the
application to show things suck as 1 rep maxes.

*Another Idea*: when adding a set, show the max resistance based on the
repetitions entered.

More of these features will be added.

# Conclusion
Compare to where I made it with my honours project, this is already better than
last time. This already looks to be useful for my personal use-case and some
friends already mentioned that this might be useful for them.

This project will 100% be taken further, I hope to be able to make something
useful, and potentially in the future I might even be able to work out some good
way to monetize it and turn it into a simple one man startup.
