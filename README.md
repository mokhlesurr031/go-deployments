### Users


1. Registration
    - **Request**
    ```
   URL: {{domain}}/api/v1/auth/
   Method: POST
   Payload:
    {
      "name": "Test User",
      "email":"test@gmail.com",
      "password":"1234",
      "password_confirm":"1234"
    }
    ```
    - **Response:**
   ```
   Status Code: 201
   {
     "id": 6,
     "name": "Test User 4",
     "email": "test4@gmail.com",
     "password": "",
     "password_confirm": "",
     "created_at": "2023-02-25T06:31:42.090517+06:00"
   }
   
   If user with existing email already exists:
     Status Code: 500
     Error: email already exists
    
    ```


2. Login
    - **Request**
    ```
   URL: {{domain}}/api/v1/auth/login/
   Method: POST
   Payload:
   {
      "email": "test@gmail.com",
      "password": "1234"
   }
    ```
    - **Response:**
   ```
   {
     "ExpiredIn": 3600000000000,
     "MaxAge": 3600,
     "Secret": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzczNTcyODYsImlkIjoiMSJ9.YcIhoNIPSX8iuBcf743qabOjrvaaxUNFtmLOHpB6U88",
     "Message": "success",
     "User": {
       "id": 1,
       "name": "Test User",
       "email": "test@gmail.com"
     }
   }
    
    ```