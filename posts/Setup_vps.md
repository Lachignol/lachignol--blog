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

### 🛒 1. Acheter un VPS

#### Connexion en SSH

```sh
ssh root@votre_ip
```

#### Mise à jour du serveur

```sh
apt update && apt upgrade -y
```

#### Installation des paquets essentiels

```sh
apt install sudo ufw curl wget git jq openssl vim tmux fish -y
```

---

### 🌐 2. Acheter un nom de domaine

Dans la gestion DNS de votre nom de domaine, ajoutez **deux enregistrements A** :

1.  
**Type :** A  
**Nom :** `@`  
**Contenu :** IP de votre serveur

2.  
**Type :** A  
**Nom :** `*` (astérisque)  
**Contenu :** IP de votre serveur

**Explication :**
- `@` correspond au domaine racine (ex : votredomaine.com)
- `*` correspond à tous les sous-domaines non définis (ex : test.votredomaine.com)

---

### 👤 3. Création d'un utilisateur non-root

#### Création de l'utilisateur

```sh
useradd -m -s /bin/fish votre_user
usermod -aG sudo votre_user
passwd votre_user
```

#### Déconnexion de root

```sh
exit
```

---

### 🔑 4. Configuration des clés SSH

Sur **votre machine locale** (pas sur le vps) :

#### Génération de la clé

```sh
ssh-keygen -t ed25519 -C "votre_email@domaine.com"
```

#### Copie de la clé publique sur le serveur

```sh
ssh-copy-id -i ~/.ssh/id_ed25519.pub votre_user@votre_ip
```

---

#### Configuration sécurisée de SSH (sur le serveur)

```sh
sudo vim /etc/ssh/sshd_config
```

Modifiez / ajoutez les lignes suivantes :

```text
PubkeyAuthentication yes
PasswordAuthentication no
PermitRootLogin yes
```

Redémarrez SSH :

```sh
sudo systemctl restart sshd
```

---

#### ✅ Activer UFW au démarrage automatiquement

UFW est normalement activé de manière persistante, **mais pour s'assurer qu'il démarre bien au boot**, on peut forcer l’activation via systemctl :

```sh
sudo systemctl enable ufw
```

> Cette commande s’assure que le pare-feu UFW est bien lancé à chaque redémarrage du serveur.

---

### 🔒 5. Configuration du pare-feu (UFW)

#### Ouverture des ports nécessaires

```sh
sudo ufw allow 22    # SSH
sudo ufw allow 80    # HTTP
sudo ufw allow 443   # HTTPS (SSL)
sudo ufw allow 8000  # Interface Coolify
sudo ufw allow 6001  # Websockets / Laravel Echo / autres services
sudo ufw allow 53    # DNS (utile pour certains conteneurs)
```

#### Activation de UFW

```sh
sudo ufw enable
```

---

### ⚙️ 6. Installation de Coolify

1. Accédez à Coolify :  
   http://votre_ip:8000

2. Créez votre compte administrateur.

3. Configurez :
   - Le **nom de domaine personnalisé**
   - Le **SSL (Let's Encrypt)**
   - Les **sauvegardes**

---

### 🧩 7. Paramétrage dans Coolify

#### 🔧 Paramètres de l'instance

- Allez dans **Settings**  
- Renseignez :
  - **Instance domain** : `https://le-nom-que-tu-veu-pour-page-login.mondomaine.com`
  - **Instance name** : `Coolify`

---

#### 🔁 Redémarrage du proxy

- Allez dans l’onglet **Servers**
- Redémarrez le proxy
- Dans **Wildcard domain**, ajoutez votre domaine (il générera automatiquement les sous-domaines nécessaires)

---

#### 🧪 Création d’un projet

1. Créez un projet
2. Cliquez sur **New GitHub App**
3. Nommez-le (ex : `GitHub-Ton-Nom-User-coolify`)
4. Laissez-vous guider par GitHub pour terminer la configuration

---

### 🧱 8. Sécuriser l’accès à Coolify via le domaine (et pas via l'IP)

#### Problème : UFW ne s’applique pas aux conteneurs Docker par défaut

👉 Solution : utiliser **ufw-docker**

---

### 🧰 9. Installation de `ufw-docker`

#### Étapes :

1. **Téléchargez le script** `ufw-docker`
2. **Rendez-le exécutable** :

```sh
sudo chmod +x chemin/vers/ufw-docker
```

3. **Installez les règles** :

```sh
sudo ufw-docker install
```

> Cette commande adapte `/etc/ufw/after.rules` pour Docker

4. **Autorisez le conteneur `coolify-proxy`** :

```sh
ufw-docker allow coolify-proxy
```

5. **Rechargez UFW si besoin** :

```sh
sudo ufw reload
```

---

### 🧠 10. Automatiser au reboot (via cron)

**UFW-Docker** doit être relancé à chaque redémarrage du serveur.

Ajoutez cette tâche cron pour l’utilisateur root :

```sh
sudo crontab -e
```

