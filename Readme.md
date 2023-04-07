# This is a basic application to try docker in local and kubernetes cluster.
This project contains
- two applications
    - redis
        - RDB and AOF implementation using conf
        - persistance storage
        - databases (logical partitions)
    - user management
        - supports C.R.D. operations
- docker
    - persistent volume
- kubernetes
    - using local image for both applications in the deployments
    - persistent volume(PV) and persistent volume claim(PVC)
    - accessing redis host and password details from secrets
    - using env variables

### Building the application in local:
~~~
docker build -t usermanagement-redis -f redis.Dockerfile .  
docker run --name usermanagement-redis -p 6379:6379 -v /<home>/go/src/github.com/userManagement/persitent_redis_data:/data/app/redis -d usermanagement-redis

source .env
go build && ./userManagement
docker build -t usermanagement -f usermanagement.Dockerfile .
docker run --name usermanagement-redis   -p 6379:6379 -v usermanagement_data:/data/app/redis -d usermanagement-redis 
~~~

### Testing in postman:
~~~
GET: localhost/users?username=Donald

POST:localhost/users
    body:{
            "username":"Donald",
            "Password":"Trump"
        }
~~~

### Deploying in k8s cluster
~~~
# for redis app:
    kubectl apply -f redis_deployment.yaml

# for user management app:
    # create pv and pvc:
    kubectl apply -f redis_pv.yaml
    kubectl get pv

    # create pod:
    kubectl apply -f redis_deployment.yaml
    kubectl get pods -o wide 
 
    # create service:
    kubectl apply -f redis_service.yaml
    kubectl get svc -o wide 
~~~

### for user management app:
~~~
    # create secret:
    kubectl apply -f user_management_secret.yaml
    kubectl get secrets

    # create pod:
    kubectl apply -f user_management_deployment.yaml
    kubectl get pods -o wide
 
    # create service:
    kubectl apply -f user_management_service.yaml
    kubectl get svc -o wide 
~~~   


### Testing persistance:
~~~
    kubectl delete pod usermanagement-redis-<POD-ID>
    kubectl scale deploy usermanagement-redis --replicas=0
    kubectl scale deploy usermanagement-redis --replicas=1
~~~
