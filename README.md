# Planches de Cartes Savage Worlds

Ayant lu avec beaucoup d'intérêt le livre [Lazy Dungeon Master](http://slyflourish.com/lazydm/), j'ai décidé de mettre en pratique ses recommandations avec [Savage Worlds](https://www.black-book-editions.fr/catalogue.php?id=58).

Pour résumer, *Lazy Dungeon Master* recommande de préparer ses parties en écrivant des fiches pour les personnages et adversaires, les lieux, les intrigues et les trésors. Ils suggèrent des fiches bristol de couleurs différentes pour les différents types (par exemple bleues pour les lieux, jaunes pour les trésors, vertes pour les intrigues, etc).

## Cartes Savage Worlds

Je me suis dit qu'il serait bien de faire des mini fiches de personnage pour les PNJs, Extras et Jokers. J'ai donc conçus ces différentes fiches pour préparer mes parties. Voici par exemple une carte pour un *Joker* :

![](nemo.png)

Par la suite, je me suis dit que ce serait une bonne idée de préparer d'autres fiches spécifiques pour les intrigues, lieux et trésors. Et puis j'ai fait d'autres fiches plus spécifiques à Savage Worlds, comme celle pour les pouvoirs.

Ayant de bons yeux, j'ai préféré les imprimer au format *A7*, soit à peu près la taille des cartes à jouer, ce qui permet d'en mettre *8* sur une page au format *A4*. On peut aussi les imprimer en *A6* , on en met alors *4*, ou en *A5*, on passe alors à *2* par page.

On trouvera [dans l'archive](http://sweetohm.net/public/sawo-cards.zip) les images des cartes au format PNG. Pour les imprimer il faut produire des planches de *2*, *4* ou *8* cartes. On peut le faire avec un logiciel de manipulation d'image, comme Gimp par exemple. On peut aussi le faire en ligne de commande avec un script comme celui fourni dans l'archive, *planches.sh* qui utilise *ImageMagick* pour générer ces planches, mais cela nécessite d'installer cet outil et de l'appeler en ligne de commande.

## Composition des planches

Pour faciliter la réalisation de ses propres planches, j'ai conçu un service web qui permet de générer ces planches sur mon serveur. Pour l'appeler il faut d'abord obtenir une clef en m'envoyant un message à <casa@sweetohm.net> avec votre nom et adresse mail. Vous recevrez alors votre clef à envoyer lors des requêtes.

Pour voir les images des cartes disponibles, on pourra ouvrir dans son navigateur l'URL (où l'on remplacera *CLEF* par la clef reçue par mail) :

- <http://sweetohm.net/sawo/cards/CLEF>

![](cartes.png)

Sous chaque carte est indiquée son nom. Pour générer un planche, appeler l'URL *http://sweetohm.net/sawo/cards/CLEF/* suivie des noms des cartes séparées de slashs.

Par exemple, pour générer une fiche de personnage et de PNJ accompagnées de leur background au format *A5*, on ajoutera sur l'URL *perso*, *background*, *pnj* et *background* :

- <http://sweetohm.net/sawo/cards/CLEF/perso/background/pnj/background>

Ce qui donnera la planche suivante :

![](planche-2x2.png)

Il est possible de générer des planches *2 x 1*, *2 x 2* et *4 x 2*. On enverra alors sur l'URL *2*, *4* ou *8* noms de cartes. Pour récupérer la planches, faire un clic droit sur l'image et sélectionner dans le menu *Enregistrer l'image sous...*.

Pour faire une planche *4x2* avec les principales cartes, on appellera :

- <http://sweetohm.net/sawo/cards/CLEF/perso/pnj/joker/extras/background/intrigue/lieu/tresor>

Ce qui donne la planche :

![](planche-4x2.png)

En pratique, les cartes les plus utilisées sont *intrigue*, *PNJ*, *Extras*, *Joker*, *lieu* et *tresor*. On prendra soin de les imprimer dans les bonnes proportions sur du bristol pour préparer ses parties.

Il est aussi possible de remplir les cartes avant de les imprimer. C'est ce que je fais pour préparer mes parties, je remplis les cartes avec mon Galaxy Note puis je les imprime par planches de *8* cartes que je compose avec un script qui appelle *ImageMagick* en ligne de commande. Je les numérote ensuite après impression pour ne jamais perdre le fil de mes fiches en cours de partie.

## Retour d'expérience

Je trouve que ces cartes sont le meilleur moyen de mener une partie : on n'est pas perdu dans les pages du scénario, on reste concentré sur la carte en cours. On peut étaler les fiches des personnages devant soi, ou donner aux joueurs des cartes avec des indices ou des trésors. Pour finir, ces cartes permettent de jouer debout autour de la table, pour manipuler les figurines et jeter les dés avec les joueurs.

J'imprime les feuilles des personnages en format *A5* (avec leur background) pour les joueurs et au format *A7* pour le maître de jeu.

Un court scénario tient sur une dizaine de cartes :

![](scenario.png)

Pour une campagne, telle que *Irongate* dans le cas de la photo, c'est bien plus conséquent et on arrive vite à des piles de cartes.

![](campagne.png)

Si vous avez des suggestions pour améliorer ces cartes ou en créer de nouvelles, n'hésitez pas à m'envoyer un email à <mailto:michel.casabianca@gmail.com>.

*Enjoy!*
