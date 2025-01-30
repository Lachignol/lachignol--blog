+++
title = "Ft_printf"
description = "Projet qui consiste a recoder la fonction printf"
date = 2025-01-31

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignol"

+++

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
##### Par exemple :
- `(nil)` pour un pointeur nul.
- `(null)` pour une chaîne nulle.