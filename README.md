# REST
Golang Rest Api 

A Basic Level Restful Api with Golang that calls two 3rd Party Api's "Last FM" & "musixmatch"
https://www.last.fm/api
https://developer.musixmatch.com/

Database - Mysql [ DB STRUCTURE ADDED BELOW ]
**Base Url - http://localhost:3200**

Get Data from Database - http://localhost:3200/api/posts -- **Only to Test DB Fetch**

Get Musical Track Name & Artist Name with respect to Country (It can be extended to location and other functionality like Pagination and Limit)
**Now Only Max 2 Record is fetched**
HTTP REQUEST **POST https://www.last.fm/api/getTopTrack**
JSON Body - {"country":"india"} **full country name**

/**** TURN OFF DB USAGE  *****/
REMOVE COMMENT FROM LINE 5 & 9 FROM main.go to turn ON DB FOR TEST DATA


Get Musical Track Name & Artist Name, Album Name, Lyrics with respect to Country (It can be extended with other functionality like Pagination and Limit)
**Now Only Max 2 Record is fetched**
HTTP REQUEST **POST https://www.last.fm/api/getTopTrackLyrics**
JSON Body - {"country":"in"} **extract first two letter of a country , example - india will be turned to in**

**A Basic Middleware to log Api Request and Time to Respond for each Api Calls**, it will be extended in future for something good

**Planned Improvements**
Error Interface
Implementing Rate Limitter
Creating Microservice
Dockerizing the Application
Reqriting few Code Base
Implenting a Framework to structure it



CREATE TABLE `cook-accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `state` varchar(56) NOT NULL,
  `name` varchar(256) NOT NULL,
  `email` varchar(256) NOT NULL,
  `contactnum` int(19) NOT NULL,
  `gender` varchar(11) NOT NULL,
  `age` int(11) NOT NULL,
  `status` smallint(2) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci


