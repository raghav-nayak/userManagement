# Start from the latest Redis base image
FROM redis:latest

# Copy Redis configuration
COPY redis.conf /etc/redis/redis.conf

# Expose the Redis port
EXPOSE 6379

# Start Redis server
CMD ["redis-server", "/etc/redis/redis.conf"]




# docker build -t usermanagement-redis -f redis.Dockerfile .  

# docker run --name usermanagement-redis -p 6379:6379 -v usermanagement_data:/data/app/redis -d redis 

#local persistance data
# docker run --name usermanagement-redis -p 6379:6379 -v /<home>/go/src/github.com/userManagement/persitent_redis_data:/data/app/redis -d usermanagement-redis
