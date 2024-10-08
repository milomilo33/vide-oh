# vide-oh
Bachelor's thesis project: Microservice-based platform for watching (streaming) and sharing video content

![image](https://github.com/milomilo33/vide-oh/assets/29868001/9e4af3ab-1eac-49ed-81dc-9bc6e7a4f7c0)

![image](https://github.com/milomilo33/vide-oh/assets/29868001/c7bfd52a-779e-4963-be78-278b2cfbde25)


# Features
The system recognizes 3 basic types of users:
- Unregistered user
- Registered user
- Administrator

Below are the functionalities that each type of user has access to.

Unregistered user:
- Registration
- Login
- Search for existing videos (with thumbnails that are automatically extracted when videos are uploaded (courtesy of _FFmpeg_))
- Watch videos (the videos will be streamed, not downloaded, and viewed within this application), control playback speed
- Download videos

Registered user:

- Search for existing videos
- Watch videos
- Download videos
- Upload videos to the system (along with details such as title, description, etc.)
- View, edit, and delete their own videos
- View the average rating and all comments on a video
- CRUD (Create, Read, Update, Delete) their own comments and ratings on each video
- Edit profile info

Administrator:

- All the functionalities that a registered user has access to (with the exception of being unable to upload videos)
- User blocking (with notification via email)
- Deletion of videos or modification of video data deemed inappropriate
- Deletion of inappropriate comments

# Architecture
The system relies on a microservice architecture.
The components of the system are as follows:

- User microservice (authentication, authorization, user addition and modification, user blocking, etc.). Technologies: _Go_ and _PostgreSQL_.
- Video microservice (CRUD operations on videos, video streaming, video search, etc.). Technologies: _Go_ and _PostgreSQL_ (for video metadata).
- Comments and ratings microservice. Technologies: _Rust_ and _PostgreSQL_.
- Central client-side frontend application that supports the functionalities of all microservices. Technologies: _Vue.js_.
- Another microservice that introduces a new type of user (technical support) with whom users can chat if they need assistance (implemented using WebSockets).
- _nginx_-based API Gateway.

<p align="center">
  <img src="https://github.com/milomilo33/vide-oh/assets/29868001/780d324c-9054-4299-95f6-e2ee83a62b79" alt="animated" />
</p>

