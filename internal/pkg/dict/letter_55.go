package dict

import "rewriting-history/internal/domain"

// letters of an 5*5 2D array of uint
var (
	L5Space = domain.Letter{
		{0},
		{0},
		{0},
		{0},
		{0},
	}
	L55A = domain.Letter{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
	}
	L55B = domain.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	L55C = domain.Letter{
		{0, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 1, 1, 1},
	}
	L55D = domain.Letter{
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
	}
	L55E = domain.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	L55F = domain.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	}
	L55G = domain.Letter{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	L55H = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}
	L55I = domain.Letter{
		{0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
	}
	L55J = domain.Letter{
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	L55K = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}
	L55L = domain.Letter{
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	L55M = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
	}
	L55N = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1},
	}
	L55O = domain.Letter{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	L55P = domain.Letter{
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	}
	L55Q = domain.Letter{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
	}
	L55R = domain.Letter{
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}
	L55S = domain.Letter{
		{0, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
	}
	L55T = domain.Letter{
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	}
	L55U = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	L55V = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
	}
	L55W = domain.Letter{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 1},
	}
	L55X = domain.Letter{
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	}
	L55Y = domain.Letter{
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	}
	L55Z = domain.Letter{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
)
