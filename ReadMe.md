### ChatGPT was used to generate this read me. minor editing was involved. Please note this work is not my own, but merely used to suggest that i have learned and covered many topics as it relates to golang. 

# Course: Learning Go Programming - Creating UserProfile Struct
In this course, I learned how to incorporate various fundamental concepts in Go programming to create a UserProfile struct. We covered the following key topics:

## 1. Creating Custom Types
To begin, we learned how to create custom types in Go. By using custom types, we can make our code more readable and maintainable. In this course, we defined a custom type usernameKey to use as a key for context values. This helps to prevent collisions with other packages or libraries using the same key. (but this part of the code is minor, but was cool to see alternative ways to prevent collisions)

## 2. Defining a UserProfile Struct
We then proceeded to define the UserProfile struct, which represents a user's profile data. The struct includes the following fields:

- ID: An integer representing the user's unique identifier.
- Comments: A slice of strings containing user comments.
- Likes: An integer indicating the total number of likes the user has received.
- Friends: A slice of integers representing the unique IDs of the user's friends.

## 3. Creating Functions
In this section, we implement essential functions for fetching user profile data with concurrent requests. The main function initiates the process, measures the execution time, and displays the results. The handleGetUserProfile function orchestrates concurrent data retrieval from three sources: comments, likes, and friends. Each function simulates data fetching with artificial delays. With concurrency, we optimize the overall performance of fetching the user profile data.


## 4. Working with Goroutines and Concurrency
To improve the efficiency and responsiveness of our application, we introduced goroutines and concurrency. We used goroutines to perform concurrent tasks, such as fetching user profiles from multiple users simultaneously.

We also employed the concept of contexts to manage the lifecycle of goroutines. By creating contexts with timeouts, we ensured that goroutines don't run indefinitely and can be canceled if needed.