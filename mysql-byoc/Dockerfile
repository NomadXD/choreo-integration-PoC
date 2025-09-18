FROM mysql:8.0

# Set environment variables for MySQL
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=incidentdb
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=supersecret

# Expose MySQL port
EXPOSE 3306

# Optional: Add custom configuration or initialization scripts
# COPY ./my.cnf /etc/mysql/my.cnf
COPY ./init.sql /docker-entrypoint-initdb.d/
