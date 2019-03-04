# Planches de Cartes Savage Worlds

Ayant lu avec beaucoup d'intérêt le livre [Lazy Dungeon Master](http://slyflourish.com/lazydm/), j'ai décidé de mettre en pratique ses recommandations avec [Savage Worlds](https://www.black-book-editions.fr/catalogue.php?id=58).

Pour résumer, *Lazy Dungeon Master* recommande de préparer ses parties en écrivant des fiches pour les personnages (*extras* ou *jokers*), les lieux, les intrigues, etc. Ils suggèrent des fiches bristol de couleurs différentes pour les différents types (bleues pour les lieux, jaunes pour les trésors, vertes pour les intrigues, etc).

Je me suis dit que ce serait bien d'avoir aussi des fiches pour les Extras, Jokers et autres PNJs. J'ai donc conçus ces différentes fiches pour préparer mes parties. Ayant de bons yeux, j'ai préféré les imprimer au format cartes *A7* (on en imprime alors *8* sur une page *A4*), mais on peut les imprimer en *A6* (on en imprime alors *4* par page) ou *A5* (on en imprime *2*).

Voici par exemple une carte pour un *Joker* :

![](nemo.png)

On trouvera [dans l'archive](http://sweetohm.net/public/sawo-cards.zip) les images des cartes au format PNG. Pour les imprimer il faut produire des planches de *2*, *4* ou *8* cartes. On trouvera dans l'archive un script *planches.sh* qui utilise *ImageMagick* pour générer ces planches, mais cela nécessite d'installer des outils et de les appeler en ligne de commande.

Pour faciliter la réalisation de ses propres planches, j'ai conçu un service web qui permet de générer ces planches sur mon serveur. Pour l'appeler il faut d'abord obtenir une clef en m'envoyant un mail à <casa@sweetohm.net> avec votre nom et adresse mail. Vous recevrez alors votre clef à envoyer lors des requêtes.

Pour voir les images des cartes disponibles, on pourra ouvrir dans son navigateur l'URL (où l'on remplacera *CLEF* par la clef reçue par mail) :

- <http://sweetohm.net:8000/CLEF>

Sous chaque carte est indiquée son nom. Pour générer un planche *2 x 2* avec les cartes *perso*, *background*, *pnj* et *background*, on appellera l'URL :

- <http://sweetohm.net:8000/CLEF/perso/background/pnj/background>

Il est possible de générer des planches *2 x 1*, *2 x 2* et *4 x 2*. On enverra alors sur l'URL *2*, *4* ou *8* noms de cartes. Pour récupérer la planches, faire un clic droit sur l'image et sélectionner dans le menu *Enregistrer l'image sous...*.

*Enjoy!*
