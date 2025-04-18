+++
title = "42-Projets de mon cursus"
description = "Une petite revue des projets."
date = 2025-04-18

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignol"

+++



### Piscine reloaded.

Après une piscine éprouvante et un mois *splité entre attente et satisfaction après le résultat tant attendu, l’aventure commence. 
On retrouve notre équipe de la piscine (un aux bouclettes stylisées par Schwarzkopf à son prime manque à l’appel, mais nous rejoindra l’année prochaine, c’est une certitude !).
On est tout de suite plongés dans le bain avec la piscine reloaded.
En gros, on a une semaine avant la deadline qui permet, comme son nom l’indique, de refaire certains des exercices de la piscine que l’on a déjà faits, ainsi que d’autres pour lesquels nous n’avons pas eu le temps. 
Pour moi, ce fut l’occasion, par exemple, de faire connaissance avec les Makefiles :scripts bash souvent utilisés pour compiler les différents projets en C ou encore les fichiers ".h" qui permettent les imports de librairies ainsi que la definition de structs/macros/fonctions etc.

### Libft.
Une fois cette semaine passée, voici le vrai premier projet du premier cercle !!! La libft. 
En réalité, ce n’est pas trop un projet classique, car le but est de constituer notre librairie afin d’avoir une multitude de fonctions essentielles qui existent déjà dans les librairies standards en C mais auxquelles nous n’avons évidemment pas droit (42 mentality !!). 
C’est pourquoi nous devons les coder nous-mêmes afin de pouvoir les utiliser plus tard. 
Le sentiment du devoir accompli après avoir réussi le ft_split et bien d’autres… 
Ce fut aussi l’occasion, dans les bonus, de faire connaissance avec une structure de données bien connue : les fameuses listes chaînées.

#### Tricky part
- Toutes les conditions de contrôle à prévoir.
Exemple basique :
```c
if (str == NULL)
    return (NULL);
```
- Arriver à conceptualiser les listes chaînées.

*fuck ft_split exercice qui a eu raison de moi à l’examen final de la piscine (score 66).

## Ft_printf

### Enrichir la libft.
Une fois la libft terminée, ce n’est pas fini ! 
Nous allons encore l’enrichir, mais cette fois avec une fonction particulière et bien connue de tous : le fameux printf. 
Quelle motivation de pouvoir enfin éviter de coder manuellement toutes les fonctions de conversion comme `itoa`, `atoi`, `put_nbr`, ainsi que leurs versions dans d’autres bases :binaire, octal ou hexadécimal. 
Cependant, pour afficher exactement ce que l’on veut, quand on le veut, nous devons implémenter ces conversions une dernière fois pour constituer ce fameux printf.

### Premiers défis.

#### Découvrir `va_args` / `va_list`.
- Ces outils permettent de gérer un nombre indéterminé d’arguments et de les parcourir un à un.
#### Parser le premier argument
- Il faut analyser la chaîne de format pour déterminer l’opération à effectuer en fonction du type attendu, indiqué par la lettre situé après le `%`.

Exemple simple :
```c
printf("Voici la string à afficher : %s",(char *) string);
```
Dans ce cas, nous devons boucler et afficher tous les caractères de la variable `string`.

Autre exemple :
```c
printf("Voici le nombre à afficher : %d",(int) nbr);
```
Ici, nous devons convertir en caractères la valeur de l’entier `int`.

Cas plus complexe :
```c
printf("Voici l'adresse du pointeur : %p", pointeur);
```
Dans ce cas, il faut afficher en hexadécimal l’adresse du pointeur.

#### Tricky part

- La fonction `printf` doit retourner le nombre total de caractères affichés.
- Répliquer les messages d’erreur spécifiques
Par exemple :
- `(nil)` pour un pointeur nul.
- `(null)` pour une chaîne nulle.

## Get next line

### Une derniere fonction pour enrichir notre librairie personnelle.

Voici la dernière fonction, plus complexe à créer, afin d’intégrer notre libft et partir mieux armé pour les projets qui nous attendent : coder la fonction get_next_line.

#### En quoi consiste-t-elle ?

Elle consiste à renvoyer une ligne (caractérisée par un `\n`) du contenu d’un fichier, une par une. Dit comme cela, cela paraît simple, mais la complexité réside dans le fait de pouvoir le faire avec une taille de buffer variable.

#### Explication.
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

#### Tricky part
- Gérer les leaks ! (fuites de mémoire possibles lorsque l’on alloue dynamiquement de la mémoire et qu’on ne la libère pas correctement).
