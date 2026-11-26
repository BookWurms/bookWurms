# bookWurms: Product Description

bookWurms is a social media platform that revolves all around books. From sharing what you are reading, leaving bookreviews, browsing through friends' libraries, searching for specific books or reading on the built in e-reader. bookWurms has it all to keep you connected, inspired and hungry for more books!
Say you want to read a book and want to check if one of your friends has it standing in a bookshelf at home? Put the name in the search mask and see all your friends that have it. And if your lucky they left a book review for you!

## Initial Setup

When you sign up the first thing you do is set up a profile page. Add a profile picture, say a few words about yourself and put a status: share what you are reading right now. 
Then you need to fill your library. You can add books by name or ISBN. Put the books that you want your friends read (and generally every book you are willing to lend a friend). 
Then it's time to connect with your friends. You can find friends by their name or e-mail address.  
If they agreed to share it, otherwise ask them for their public userID. [optional]

## Features

### Search mask for books

A search mask that lets you search after specific books in the libraries of your friends. You can search by name or ISBN.  

Requirements from modules: 
* Implement advanced search functionality with filters, sorting, and pagination

### Private chat

A chat function to hit up your friends. Just ask them what they think about a book or if you can borrow it.  

Requirements from modules: 
* A basic chat system (send/receive messages between users)
* A friends list, that displays online/offline status
* A friends system (add/remove friends, see friends list)

### E-reader

A built in e-reader and access to a lot of books in the public domain.  

Requirements from modules:  
* Support multiple file types (.epub and .pdf)
* Client-side and server-side validation (type, size, format)
* Secure file storage with proper access control.
* File preview functionality where applicable. 
* Progress indicators for uploads. 
* Ability to delete uploaded files.

### Profile page 

In your profile page you have the option to add a profile picture and a bio.  
Others can see this page and can access your library (and posts and history) from here.  
There is also a status for sharing what book you are *currently reading*. [optional]  
You can add tags for your favorite genres or authors you like. [optional]  

Requirements from modules: 
* Users can update their profile information.
* Users can upload an avatar (with a default avatar if none provided).
* Users can add other users as friends and see their online status.
* Users have a profile page displaying their information.
* Advanced permissions system: View, edit, and delete users (CRUD). 
* Roles management (admin, user, guest, moderator, etc.) (for us admin and user)

### The feed
Core of every social media platform. Here you see a timeline of all the things your friends are doing. You will see:
* Additions to libraries
* New Bookreviews
* Changes to *currently reading* [optional]
* Freely desinged posts (sharing thoughts or quotes; asking for recommendations) [optional]  

Requirements from modules: 
* real-time features using WebSockets
* Real-time updates across clients
* filters, sorting, pagination

### Comments and likes

The items in the feed will appear as a post and have to function to like or comment on them.  

Requirements from modules: 
* Real-time updates across clients

### Currently reading [optional] 

Share what you are *currently reading*, even if it is not in your library. 
If you allow, you will be prompted to write a review, when you finished reading a book ,indicated by a change of *currently reading* status. [optional]

### Personal book history [optional]

Books that you had in your "*currently reading*" tab will automatically be added your history.
You can also add books manually. That way you can write reviews for books that neither are in your library nor were in your currenlty reading status.

### Booksharing helper [optional]

Interface for requesting to borrow a book and a status information about availibilty in the libraries. 