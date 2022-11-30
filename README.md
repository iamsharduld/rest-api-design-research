## REST API Design Research

### Project Description

The research aims to study developer perception of different REST API error handling methodologies and aims to measure various metrics and provide evidence to improve this error handling. The repository provides starter code for getting started with 9 different API combinations. We have used Go Programming language and implemented the APIs with the help of [echo](https://echo.labstack.com/) Framework.

 Weather (/weather) | Stock Market (/stockPrice) | Heart Rate (/heartRate) |
--- | --- | --- 
400 Bad Request | 400 Bad Request | 400 Bad Request |
500 Internal Server Error | 500 Internal Server Error | 500 Internal Server Error |
200 OK | 200 OK | 200 OK |


### Getting started

1. Installing go - https://go.dev/doc/install
2. Clone this repository
3. Go to project's root and run ``` go run main.go```


### Testing APIs

The API endpoints are named /heartRate1, /heartRate2, /heartRate3 and have respective names for other endpoints. Each API request expects a key in the body (POST) which is given to uniquely identify the user requesting the resource.

Please refer to below screenshot for more details

<img width="1317" alt="image" src="https://user-images.githubusercontent.com/8417988/204929611-d22dd750-d336-4475-9c58-13daf9e847a5.png">


### Deployment and Data Collection

Deployment steps are - https://devcenter.heroku.com/articles/getting-started-with-go

We use variety of Heroku add-ons to log API invocation information. The `key` passed in every request body is used to uniquely identify the end user. This information is then processed to collect metrics like time to implement for combinations of API endpoints.


### Next Steps

* Adding complexity to APIs with different ways to get errors inside a single endpoint.
* Design and develop user interface for research study users to implement APIs on. Figma prototype is in progress, please contact repo collaborators to get access.


### Study Survey and Sign up

Please fill the below form to sign up for the study
https://forms.gle/e563JRjtsDCtBZ4D7

For edit permissions please contact - sdeshpan@ucsd.edu, mcoblenz@ucsd.edu
