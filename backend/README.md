# CVWO Backend

The backend for the CVWO Project

This project will be built with microservices (which is overkill for a project like this), mainly as a challenge for me, and to learn how and why devs build with microservices, how inter service communication works and how does one do scaling


## Architecture
1. API Gateway - The interface for the frontend to connect to the microservices. This is where the "plumbing" is done. The API gateway will provide a graphql interface for the frontend. 

2. Authentication Service - This service will handle User registration and User Login

3. Thread Service - Provides a list of threads, New Threads, Thread Deletion, Comment Creation, Comment Deletion

4. User Service - Shows the threads the user has commented on, number of comments, a general user profile

5. Admin Service - Interface for admin users to delete other threads

Between services, they'll be using gRPC to communicate. The main communication is between API Gateway and the other services. I don't expect much communication between the other services. 

The API gateway will handle authorization and putting the user and other information received from the frontend in context.

