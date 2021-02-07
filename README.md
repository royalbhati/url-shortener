
URL Shortener



Distributed Unique ID generator

While there few soltions available for this problem each of them have certain benifits
and downfalls

 - UUIDs
   Unique but too long for the use case 

 - Ticket Server
    Having a ticket server is the most simple and performant way of dealing with the problem of generating unique IDs
    Simple way of doing that is having a redis service every time
    a key is requested increment the popped key

     but its has a SPOF and maintaining replicas is again a overhead

 - Twitter's snowflake like approach 
   We can generate unique keys using the combination timestamp,machine id,node id and seq ids
   nodeids and machine ids are fixed at startup while the timestamp and seq ids are genrated on the fly

Since the project is more of me learning rather than making it a production ready system I'll be using the Ticket Server approach as its easy to implement




  