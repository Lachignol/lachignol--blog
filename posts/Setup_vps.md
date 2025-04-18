+++
title = "Installation de mon premier serveur"
description = "Tuto d'installation de mon serveur"
date = 2025-04-19

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignol"

+++


#### Je ne vais pas raconter ma vie, j’ai juste créé ce blog en Golang afin de pouvoir écrire des articles (c’est un grand mot 🤦) depuis mon téléphone au format Markdown.

### grand titre (après celui de l’article en metada

Connectez-vous en SSH :

```sh
ssh root@votre_ip

```

Mettre à jour son serveur

```sh

apt update && apt upgrade -y

```

Installation des paquets essentiels :

```sh
apt install sudo ufw curl wget git jq openssl vim tmux fish -y
```
### Acheter nom de domaine

puis aller dans la partie souvent dns:

Fait deux enregistrement
Le premier :
Type A
Nom: @
Contenu : ip du serveur
Le deuxième:
Type A
Nom:*(asterisque)
Contenu:ip du serveur 

@ : correspond au domaine racine (ex: votredomaine.com)
• * : correspond à tous les sous-domaines non spécifiquement définis (ex: test.votredomaine.com)


### Création d'un utilisateur non-root
#### Création de l'utilisateur
useradd -m -s /bin/fish votre_user
usermod -aG sudo votre_user
passwd votre_user

#### Déconnexion
exit


Configuration des clés SSH
Sur votre machine locale :
#### Génération de la clé
ssh-keygen -t ed25519 -C "votre_email@domaine.com"

#### Copie de la clé publique vers le serveur
ssh-copy-id -i ~/.ssh/id_ed25519.pub votre_user@votre_ip


Configuration SSH sécurisée
Sur le serveur :
sudo vim /etc/ssh/sshd_config

Modifiez/ajoutez les lignes suivantes :
PubkeyAuthentication yes
PasswordAuthentication no
PermitRootLogin yes
Redémarrez le service SSH :
sudo systemctl restart sshd
Configuration du pare-feu

#### Activation des ports essentiels
sudo ufw allow 22
sudo ufw allow 80
sudo ufw allow 443
sudo ufw allow 8000
sudo ufw allow 6001
sudo ufw allow 53


#### Activation du pare-feu
sudo ufw enable

### Installer coolify

#### Configuration initiale
1. Accédez à votre instance Coolify via : http://votre_ip:8000
• Créez votre compte administrateur
• Configurez les paramètres suivants :
- Nom de domaine personnalisé
- Activation SSL avec Let's Encrypt
- Configuration des sauvegardes

2. Allez dans settings

Dans le champ instance domaine:
Mettre votre nom de domaine:
ex: https://nom-que -tu-veu-pour-acceder-page-login-collify-de-ton-serveur.tonnomdedomaine.nimp

Dans instance name :
Collify


3. Allez dans serveur 
restart le proxy:
Dans wildcarddomaine tu peu aussi rentrer ton nom de domaine ( du coup il créera automatiquement les truc qui veu genre préboot etc sur un sous domaine de ton domaine

Et crée un projet et mettre ajouter GitHub
Dans New GitHub app
Dans name mettre par exemple GitHub-Lachignol-coolify et ensuite cela va ouvrir GitHub tu met aussi le nom que tu veu et tu te laisse porter loool

Dernière étape car ufw est pas appliquer au conteneur docker :

Car c’est plus secure de bloquer l’adresse :

http://votre_ip:8000

Et que soit seulement accessible la page avec certificat ssl sur ton domaine qu’à tu à renseigner dans la partie :

Allez dans settings
Et dans instance domaine :
mettre votre nom de domaine ex(https://nom-que -tu-veu-pour-acceder-page-login-collify-de-ton-serveur.tonnomdedomaine.nimp

Il faut installer

#### Ufw-docker

##### Étapes d’installation:
1. Téléchargez le script ufw-docker :

2. Rendez le script exécutable :

sudo chmod +x chemin du script

Installez les règles nécessaires dans UFW (déjà fait avant ):


sudo ufw-docker install

Cette commande sauvegarde `/etc/ufw/after.rules` et ajoute les règles nécessaires pour que UFW et Docker fonctionnent ensemble correctement


Pour autoriser un port d’un conteneur :
( ici nom du conteneur collify-proxi:
`ufw-docker allow <nom_du_conteneur> <port>`

Pensez à bien relancer UFW (`sudo ufw reload`) après modification des règles si nécessaire.

Solution temporaire
Cron tab root qui a chaque reboot fait un ufw-docker install et ufw-docker allow conteneur

( ce renseigner sur systèmed)

J’ai aussi installer fail2ban ( recommander )

