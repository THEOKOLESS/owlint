# Test technique backend

Owlint a besoin de toi ! L'un de nos nouveaux clients a besoin d'un système de
commentaires à différents endroits de notre plateforme.

Voici les besoins:

1. Le service devra être implémenté en Go, et le stockage des données
   devra être effectué dans MongoDB.

2. Ce système fonctionnera via une API REST qui devra respecter scrupuleusement
   la spécification OpenAPI fournie avec ce document, afin de permettre à
   l'équipe frontend et l'équipe backend de travailler de concert.

3. Les commentaires pourront être ajoutés sur diverses entités (profils,
   publications, photos, et même d'autres commentaires). La cible (_target_)
   d'un fil de commentaires sera donnée comme un identifiant externe que l'on
   assure unique (exemple : `Photo-123`, `Profil-234`, etc).

4. L'auteur d'un commentaire sera représenté par un identifiant unique
   provenant d'une autre application (exemple : `User-345`). Le système de
   gestion des commentaires ne devra pas gérer l'authentification des
   utilisateurs. 

5. Les nouveaux commentaires devront être transmis via une requête HTTP POST à
   un autre service qui se trouve à l'adresse suivante :
   `http://tech-test-back.owlint.fr:8080/on_comment`. Il attend une message
   JSON contenant les attributs `author` et `message`, tous deux de type
   `string`. Ce service peut être instable.

6. Les commentaires seront écrits en français ou en anglais. Dans un document
   au format de ton choix, explique nous comment tu aurais intégré à ton code
   la fonctionnalité suivante :

   - Si un commentaire est ajouté en français, il doit être traduit en anglais.
   - Si un commentaire est ajouté en anglais, il doit être traduit en français.

7. Les commentaires doivent être arrangés selon **une seule** de ces méthodes,
   on te laisse choisir:

   - (préféré) Soit sous forme hiérarchique: les commentaires sont restitués sous
     forme de fils de discussion, incluant toutes les réponses.

   - (si tu manques de temps) Soit sous forme de liste: une réponse à un commentaire sur `Photo-123` est
     ajouté à la suite des commentaires de `Photo-123`.

Quelques conseils:

- Nous observerons tout, aussi bien le fond que la forme de ce que tu vas
  créer, donc attention à la qualité de ton code.

- Si tu penses ne pas avoir le temps de tout faire, on comprend parfaitement,
  fais des choix et priorise (mais explique nous pourquoi).

La remise du projet est prévue par l'envoi du lien de partage d'un repo Github
privé par email ou un zip avec tout les fichiers.

You've got this <3
