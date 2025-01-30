+++
title = "Get next line"
description = "Coder une fonction qui retourne a chaque appel une ligne d'un fichier"
date = 2025-01-31

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignol"

+++

### Une derniere fonction pour enrichir notre librairie personnelle.

Voici la dernière fonction, plus complexe à créer, afin d’intégrer notre libft et partir mieux armé pour les projets qui nous attendent : coder la fonction get_next_line.

#### En quoi consiste-t-elle ?

Elle consiste à renvoyer une ligne (caractérisée par un `\n`) du contenu d’un fichier, une par une. Dit comme cela, cela paraît simple, mais la complexité réside dans le fait de pouvoir le faire avec une taille de buffer variable.

##### Explication
Il faut savoir que la fonction `read` place un curseur là où elle s’est arrêtée. 
Par exemple :
- Si nous appliquons la fonction `read` avec un buffer de taille 50 et que la première ligne de notre fichier fait seulement 25 caractères (en comptant le \n), nous renvoyons donc les 25 premiers caractères.
- Lors du prochain appel de `read`, nous repartirons à partir du 50ᵉ caractère. Les 25 caractères présents après le saut de ligne sont alors perdus !
D’où l’utilité de pouvoir stocker des valeurs entre plusieurs appels de la fonction. 
C’est ici que les variables statiques entrent en jeu. 
Elles permettent de stocker, dans cet exemple, les 25 caractères restants apres le renvoi de la premiere ligne et de repartir avec `read` à la position du curseur (donc ici 50ᵉ caractère) , concaténer les 25 caracteres precedants avec ceux que nous venons de lire et d'itérer ainsi jusqu’au prochain `\n`.

Exemple avec un buffer plus petit:
Prenons un buffer de taille 10 et une première ligne de 25 caractères (toujours en comptant le /n) :
- Il faut appeler plusieurs fois la fonction `read` et stocker son contenu dans une variable statique.
-	Après trois appels (10 + 10 + 10), nous atteignons le premier saut de ligne.
-	Nous renvoyons alors cette ligne et stockons ce qui dépasse le saut de ligne (5 caractères dans cet exemple) pour ne pas perdre le contenu existant jusqu’à la position du curseur lors du prochain appel.

Tricky part
- Gérer les leaks ! (fuites de mémoire possibles lorsque l’on alloue dynamiquement de la mémoire et qu’on ne la libère pas correctement).