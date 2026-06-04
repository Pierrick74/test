# Tri de données CSV en Go

Projet d'apprentissage en Go : implémentation et comparaison de plusieurs
**algorithmes de tri** appliqués à un fichier CSV (liste de communes françaises).

Le programme lit un CSV, trie les lignes selon une colonne, puis réécrit le
résultat dans un nouveau fichier. Un `defer` mesure le temps d'exécution du tri.

## Structure

```
.
├── hello.go            # point d'entrée (main) : lecture → tri → écriture
├── go.mod              # module example/hello (Go 1.26)
├── util/
│   ├── csv.go          # ReadCSV / WriteCSV (séparateur ';')
│   ├── select.go       # tri par sélection
│   ├── insert.go       # tri par insertion
│   ├── bubble.go       # tri à bulles
│   ├── shell.go        # tri de Shell
│   ├── selectIA.go     # versions « IA » : optimisées / corrigées
│   ├── insertia.go     #   (mêmes algos, plus rapides et plus propres)
│   ├── bubbleIA.go     #
│   └── shellIA.go      #
├── fr.csv              # jeu de données complet (~36 800 communes)
├── small.csv           # petit échantillon pour tester
└── frO.csv / output.csv# fichiers de sortie générés
```

## Les algorithmes

Toutes les fonctions de tri partagent la même signature et la même convention :

```go
func TriXxxInt(tab [][]string, col int) [][]string
```

- `tab` : les lignes du CSV (sans l'en-tête).
- `col` : l'index de la colonne servant de clé de tri (triée comme un entier).
- La colonne est convertie **une seule fois** en `[]int64` (`keys`) pour éviter
  de re-parser les chaînes à chaque comparaison.

Chaque algorithme existe en **deux versions** : une codée à la main, et une
codée par Claude (suffixe `IA`). Le but est de **comparer leur temps
d'exécution** et de comprendre le coût des différentes opérations.

| Algorithme | Ma version    | Version Claude (« IA ») |
|------------|---------------|-------------------------|
| Sélection  | `TriSelectInt`| `TriSelectIAInt`        |
| Insertion  | `TriInsertInt`| `TriInsertIAInt`        |
| Bulles     | `TriBubbleInt`| `TriBubbleIAInt`        |
| Shell      | `TriShellInt` | `TriShellIAInt`         |


## Lancer le programme

```bash
go run .
# ou
go build -o hello . && ./hello
```

Par défaut (`hello.go`), le programme :
1. lit `fr.csv` ;
2. retire la ligne d'en-tête ;
3. trie par la colonne d'index `9` (`codes_postaux`) via `TriSelectInt` ;
4. réinsère l'en-tête et écrit le résultat dans `frO.csv` ;
5. affiche le temps de tri.

Pour tester un **autre algorithme** ou une **autre colonne**, modifie l'appel
dans `hello.go` :

```go
orderDatas := utils.TriShellIAInt(datas, 9) // ex. tri de Shell optimisé
```

## Format CSV

Les fichiers utilisent le **point-virgule (`;`)** comme séparateur. L'en-tête :

```
EU_circo;code_region;region;chef_lieu;num_dpt;nom_ddpt;pref;num_circ;
nom_commune;codes_postaux;code_insee;latitude;longitude;dist
```

> Le tri se faisant sur des entiers (`ParseInt`), choisis une colonne numérique
> (ex. `num_dpt`, `codes_postaux`, `code_insee`).
