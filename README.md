# Test Case API

## API Authorization
- Header: Bearer your_token_here

## Endpoints

- GET /api - _API_
    - POST /auth/login - _Login user_
      - identity [string]
      - password [string]
    - POST /auth/register - _Register new user_
      - email [string]
      - username [string]
      - password [string]
      - name [string]
      - address [string]
      - role [string [author, visitor]]
      - avatar_url [string]
    - All of these route are protected using Bearer token 
    - GET /user/profile - _Get current user profile_
    - GET /user/:id - _Get user profile_
    - PUT /user/:id - _Update user_
      - name [string]
      - address [string]
      - role [string [author, visitor]]
      - avatar_url [string]
    - PUT /change-password - _Change password_
      - password [string]
      - new_password [string]
    - POST /img - _Upload image [avatar, featured_image]_
      - file [file]
    - POST /article - _Create new article_
      - title [string]
      - content [string]
      - featured_image_url [string]
    - GET /article - _Get all article_
    - GET /article/{:id} - _Get article by id_
    - PUT /article/{:id} - _Update article by id_
      - title [string]
      - content [string]
      - featured_image_url [string] 
    - DELETE /article/{:id} - _Delete article by id_
    - POST /article/{:id}/comment - _Create new comment on article_
      - comment_text [string]
    - DELETE /comment/{:id} - _Delete comment_
    - PUT /comment/{:id} - _Update comment_
      - comment_text [string]
