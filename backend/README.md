Running go server on docker:
-
1. cd into ./backend directory in your terminal
2. Spin up the docker containers
``docker compose build backend``
3. Run the application
``docker compose up backend``
4. Visit http://localhost:8000/ 
5. To stop the server
``CTRL+C``
6. To remove the containers
``docker compose down``
7. To view logs for backend
``docker compose logs -f backend``