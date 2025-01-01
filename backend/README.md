Running go server on docker:
-
1. cd into ./backend directory in your terminal
2. Spin up the docker containers
``docker build -t backend .``
3. Run the application
``docker run -p 8000:8000 backend``
4. Visit http://localhost:8000/ 
5. To stop the server
``CTRL+C``
6. 