#!/bin/sh

# set -e

# Exécuter les migrations
#goose -dir ./migrations postgres postgres://menusplanningapp_user:menusplanningapp_pass@db:5432/menusplanningapp_db up
# cd ./migrations
goose -dir ./migrations up
# cd ..
# go build -o ./bin/api ./cmd/api 
echo "Current directory: $(pwd)"
echo " ls : $(ls)"
# cd /bin
# echo "ls 2 : $(ls)"

exec "$@"
# Run in debug or production mode based on DEBUG flag
# if [ "$DEBUG" = "true" ]; then
#   cd /menusplanningapp    
#   echo "Starting in DEBUG mode..."
#   exec dlv debug ./cmd/api --headless --listen=:40000 --api-version=2 --accept-multiclient
# else
#   echo "Starting in PRODUCTION mode..."
#   exec "$@"
# fi

# Exécuter la commande passée en argument (démarrage de l'application)
# exec "$@"