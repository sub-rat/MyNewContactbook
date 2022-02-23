# MyNewContactbook

Base URL = http://localhost:8080

### Contact Book
1. Add Contact  
    fields: id, name, email, phone, address
2. List Contact
    2.1 Pagination
    2.2 Filters ( name, phone)
3. Get Contact by id
5. Bulk Delete Contact
6. Update Specific Contact
7. Delete Specific Contact


## Endpoints documents
1. GET      /contacts   Get all contacts
            /contacts?page=1&size=20
            /contacts?page=1&size=20&name="ram"&phone="909090"
2. POST     /contacts   Create new contact
                 data = { id : 1, name: "ram", email :"gopal", phone :"9898"}
3. DELETE   /contacts
4. GET      /contacts/{id}  Get detail of contact with id
5. PUT      /contacts/{id}  UPdate detail of contact with id
6. DELETE   /contacts/{id}  Delete contact with id



### Users
/// Contact with users associated
    /auth/login
    /auth/register
    /users?page=1&size=20
    /users/{id}
    /users/{id}/contacts?page=1&size=20
    /users/{id}/contacts/{id}



1. GET      /contacts   Get all contacts
            /contacts?page=1&size=20
            /contacts?page=1&size=20&name="ram"&phone="909090"
2. POST     /contacts   Create new contact
                 data = { id : 1, name: "ram", email :"gopal", phone :"9898", user_id:1}
3. DELETE   /contacts
4. GET      /contacts/{id}  Get detail of contact with id
5. PUT      /contacts/{id}  UPdate detail of contact with id
6. DELETE   /contacts/{id}  Delete contact with id

