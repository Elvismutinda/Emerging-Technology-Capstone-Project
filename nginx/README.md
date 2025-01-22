Test frontend

1. check if all frontend instances are running:
   docker ps | grep frontend

Test NGINX Load Balancing

1. test if the load balancer is correctly distributing traffic across the frontend instances:
   curl -I http://localhost:3000

2. You can run multiple requests to observe if traffic is distributed correctly:
   for i in {1..5}; do curl -I http://localhost:3000; done

3. check if requests are distributed:

curl http://localhost:3000/

4. sends 100 requests with a concurrency level of 10.
   ab -n 100 -c 10 http://localhost:3000/

5. Check NGINX access logs to see how requests are being routed:
   docker exec -it emerging-technology-capstone-project-nginx-1 cat /var/log/nginx/access.log

6. check if the backend is correctly distributing traffic across the backend instances:
   curl -I http://localhost:8000
