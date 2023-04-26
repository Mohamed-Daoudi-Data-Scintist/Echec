package main

import "math"

// Définition des pièces d'échecs
type CaseType int

const (
	Vide CaseType = iota
	Pion
	Cavalier
	Tour
	Fou
	Dame
	Roi
)

// Définition des couleurs des pièces
const (
	Blanc = iota
	Noir
)

// Définition du plateau de jeu
type Plateau struct {
	Cases [8][8]CaseType
}

// Définition de la structure des pièces
type Piece struct {
	Kind  CaseType
	Color int
	Case  Case
}

// Définition de la structure des coordonnées de la case
type Case struct {
	Abscisse int
	Ordonne  int
}

func MovePiece(plateau *Plateau, piece *Piece, dest Case) bool {
	// Vérifier si le déplacement est valide
	if !IsValidMove(plateau, piece, dest) {
		return false
	}

	// Déplacer la pièce
	plateau.Cases[dest.Abscisse][dest.Ordonne] = plateau.Cases[piece.Case.Abscisse][piece.Case.Ordonne]
	plateau.Cases[piece.Case.Abscisse][piece.Case.Ordonne] = Vide

	// Mettre à jour la position de la pièce
	piece.Case = dest

	return true
}

func IsValidMove(plateau *Plateau, piece *Piece, dest Case) bool {
	// Vérifier si la destination est hors plateau
	if dest.Abscisse < 0 || dest.Abscisse > 7 || dest.Ordonne < 0 || dest.Ordonne > 7 {
		return false
	}

	// Vérifier si la case de destination est occupée par une pièce de même couleur
	if plateau.Cases[dest.Abscisse][dest.Ordonne] != Vide && plateau.Cases[dest.Abscisse][dest.Ordonne] != piece.Kind {
		return false
	}
	// Vérifier si le déplacement est valide en fonction du type de pièce
	switch piece.Kind {
	case Pion:
		// On commence avec les Blancs
		if piece.Color == Blanc {
			// Déplacement vers l'avant
			if dest.Ordonne == piece.Case.Ordonne-1 && plateau.Cases[dest.Abscisse][dest.Ordonne] == Vide {
				return true
			}
			// Déplacement initial de deux cases vers l'avant
			if dest.Ordonne == piece.Case.Ordonne-2 && piece.Case.Ordonne == 6 && plateau.Cases[dest.Abscisse][dest.Ordonne] == Vide && plateau.Cases[dest.Abscisse][dest.Ordonne+1] == Vide {
				return true
			}
			// Capture diagonale
			if dest.Ordonne == piece.Case.Ordonne-1 && math.Abs(float64(dest.Abscisse-piece.Case.Abscisse)) == 1 && plateau.Cases[dest.Abscisse][dest.Ordonne] != Vide {
				return true
			}
		}
	}
}
