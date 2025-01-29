+++ title = “Libft : Premier Projet à 42” description = “Création de notre bibliothèque personnelle en C” date = 2025-01-30
author name = “La Chignol” email = "pas d'email for you"
footer copyright = “©Lachignol” +++

## Libft

### Premier projet

Après une piscine éprouvante et un mois *splitté entre attente et satisfaction après le résultat tant attendu.
L’aventure commence, on retrouve notre équipe de la piscine (un au bouclette stylisé par schwarzkopf à son prime manque à l’appel mais nous rejoindra l’année prochaine, c’est une certitude!!) 

On est tout de suite plongé dans le bain avec la piscine reloaded.
En gros on a une semaine avant la deadline qui permet comme son nom l'indique de refaire certains des exercices de la piscine que l’on a déjà fait,
ainsi que d'autres pour lesquels nous n’avons pas eu le temps.
Pour moi, ce fut l’occasion par exemple de faire connaissance avec les Makefiles qui sont les scripts bash souvent utilisés pour compiler les différents projets C.

Une fois cette semaine passée, voici le vrai premier projet du premier cercle !!! 
La libft en réalité c'est pas trop un projet classique car le but est de constituer notre librairie afin d’avoir une multitude de fonctions essentielles qui existent déjà dans les librairies standard en C mais auxquelles nous n’avons évidemment pas le droit (42 mentality !!).
Ce pour quoi nous devons les coder nous-mêmes afin de pouvoir les utiliser plus tard,le sentiment du devoir accompli après avoir réussi le ft_split et bien d’autres;
Ce fut l’occasion dans les bonus de faire connaissance avec une structure de données bien connue : les fameuses listes chaînées.

#### Tricky part
- Toutes les conditions de contrôle à prévoir.
Exemple basique :
```c
if (str == NULL)
    return (NULL);
```

- Arriver à conceptualiser les listes chaînées.

* fuck ft_split exercice qui a eu raison de moi à l’examen final.