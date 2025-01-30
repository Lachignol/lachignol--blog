+++
title = "Libft"
description = "Création de notre bibliothèque personnelle en C"
date = 2025-01-30

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignol"

+++



### Piscine reloaded 

Après une piscine éprouvante et un mois *splité entre attente et satisfaction après le résultat tant attendu, l’aventure commence. 
On retrouve notre équipe de la piscine (un aux bouclettes stylisées par Schwarzkopf à son prime manque à l’appel, mais nous rejoindra l’année prochaine, c’est une certitude !).
On est tout de suite plongés dans le bain avec la piscine reloaded.
En gros, on a une semaine avant la deadline qui permet, comme son nom l’indique, de refaire certains des exercices de la piscine que l’on a déjà faits, ainsi que d’autres pour lesquels nous n’avons pas eu le temps. 
Pour moi, ce fut l’occasion, par exemple, de faire connaissance avec les Makefiles :scripts bash souvent utilisés pour compiler les différents projets en C ou encore les fichiers ".h" qui permettent les imports de librairies ainsi que la definition de struct/macros/fonctions etc.

### Libft
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