#!/bin/sh

echo "Goose version:"
goose -version

# Exécuter les migrations
#goose -dir ./migrations postgres postgres://menusplanningapp_user:menusplanningapp_pass@db:5432/menusplanningapp_db up
cd ./migrations
goose up

# Exécuter la commande passée en argument (démarrage de l'application)
exec "$@"